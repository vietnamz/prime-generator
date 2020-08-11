package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/pkg/signal"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	apiServer "github.com/vietnamz/prime-generator/api/server"
	"github.com/vietnamz/prime-generator/api/server/middleware"
	"github.com/vietnamz/prime-generator/api/server/router"
	"github.com/vietnamz/prime-generator/api/server/router/prime"
	systemRouter "github.com/vietnamz/prime-generator/api/server/router/system"
	"github.com/vietnamz/prime-generator/cli/debug"
	"github.com/vietnamz/prime-generator/daemon"
	"os"
	"runtime"
	"strings"
)

type DaemonCli struct {
	config *daemon.Config
	flags *pflag.FlagSet
	api *apiServer.Server
	Hosts []string
	daemonConfig *daemon.Config
}

func NewDaemonCli() *DaemonCli  {
	return &DaemonCli{

	}
}
func (cli *DaemonCli) stop() {
	logrus.Info("Stop daemon")
	cli.api.Close()
}

func newAPIServerConfig( cli *DaemonCli) (*apiServer.Config, error) {
	serverConfig := &apiServer.Config{
		Logging: true,
		Version: "1.0.0",
		CorsHeader: "*",
	}
	return serverConfig, nil
}

func (cli *DaemonCli) initMiddleware( s *apiServer.Server, cfg *apiServer.Config ) error {
	logrus.Printf("init the cors middleware %s" , cfg.CorsHeader)
	if cfg.CorsHeader != "" {
		c := middleware.NewCORSMiddleware(cfg.CorsHeader)
		s.UseMiddleware(c)
	}
	return nil
}
func initRouter( opts routerOptions) {
	routers := []router.Router {
		systemRouter.NewRouter(),
		prime.NewRouter(opts.daemon),
	}
	opts.api.InitRouter(routers...)
}

func loadListeners(cli *DaemonCli, serverConfig *apiServer.Config) ([]string, error) {
	var hosts []string
	seen := make(map[string]struct{}, len(cli.config.Hosts))

	for i := 0; i < len(cli.config.Hosts); i++ {
		var err error
		if cli.config.Hosts[i], err = ParseHost(false, false, cli.config.Hosts[i]); err != nil {
			return nil, errors.Wrapf(err, "error parsing -H %s", cli.config.Hosts[i])
		}
		if _, ok := seen[cli.config.Hosts[i]]; ok {
			continue
		}
		seen[cli.config.Hosts[i]] = struct{}{}

		protoAddr := cli.config.Hosts[i]
		protoAddrParts := strings.SplitN(protoAddr, "://", 2)
		if len(protoAddrParts) != 2 {
			return nil, fmt.Errorf("bad format %s, expected PROTO://ADDR", protoAddr)
		}

		proto := protoAddrParts[0]
		addr := protoAddrParts[1]

		ls, err := InitListeners(proto, addr, "", nil)
		if err != nil {
			return nil, err
		}
		logrus.Debugf("Listener created for HTTP on %s (%s)", proto, addr)
		hosts = append(hosts, protoAddrParts[1])
		cli.api.Accept(addr, ls...)
	}

	return hosts, nil
}

func (cli *DaemonCli) start(opts *daemon.Config) (err error )  {
	logrus.Info("Start a daemon")
	stopc := make(chan bool)
	defer close(stopc)
	logrus.Info("Starting up")
	cli.flags = opts.Flags
	cli.config = opts
	if cli.config.LogEnable {
		debug.Enable()
	}
	if runtime.GOOS == "linux" && os.Getegid() != 0 {
		return fmt.Errorf("App needs to be started without root")
	}

	//TODO Start a web server
	serverConfig, err := newAPIServerConfig(cli)

	if err != nil {
		return errors.Wrap(err, "Failed to create api server")
	}

	cli.api = apiServer.New(serverConfig)
	_, err = loadListeners(cli, serverConfig)
	if err != nil {
		return errors.Wrap(err, "failed to load lister")
	}

	// create a context so that we can control how to terminate a app.
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	signal.Trap(func() {
		logrus.Print("graceful to close")
		cli.stop()
		<-stopc // wait for daemonCli.start() to return
	}, logrus.StandardLogger())


	cli.initMiddleware(cli.api, serverConfig)
	logrus.Info("Daemon has completed initialization")

	routerOptions, err := newRouterOptions(daemon.NewDaemon(cli.config))
	if err != nil {
		return err
	}

	routerOptions.api = cli.api
	initRouter(routerOptions)

	serverAPIWait := make(chan error)
	go cli.api.Wait(serverAPIWait)

	errAPI := <-serverAPIWait
	if errAPI != nil {
		return errors.Wrap(errAPI, "shutting down due to ServeAPI error")
	}
	logrus.Info("Daemon shutdown complete")
	return nil
}

type routerOptions struct {
	api *apiServer.Server
	daemon *daemon.Daemon
}

func newRouterOptions( d *daemon.Daemon ) (routerOptions, error)  {
	return routerOptions{
		daemon: d,
	}, nil
}