package main

import (
	"fmt"
	"github.com/docker/docker/pkg/term"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vietnamz/prime-generator/cli"
	"github.com/vietnamz/prime-generator/daemon"
	"io"
	"os"
)

func initLogging(_, stderr io.Writer )  {
	logrus.SetOutput( stderr)
}
func runDaemon(opts *daemon.Config) (err error) {
	daemonCli := NewDaemonCli()
	fmt.Printf("%s", opts.Hosts)
	return daemonCli.start(opts)
}
func newDaemonCommand() (*cobra.Command, error) {
	opts := daemon.NewDaemonConfig()
	cmd := &cobra.Command{
		Use: "Prime Generation [OPTIONS]",
		Short: "A self-sufficent runtime for application",
		SilenceUsage: true,
		SilenceErrors: true,
		Args: cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Flags = cmd.Flags()
			return runDaemon(opts)
		},
		Version: fmt.Sprintf("%s, build %s", "1.0.0", "master"),
	}
	flags := cmd.Flags()
	flags.BoolP("version", "v", false, "Print version information and quit")
	opts.InstallFlags(flags)
	return cmd, nil
}
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	})
	_, stdout, stderr := term.StdStreams()
	initLogging(stdout, stderr)

	onError := func( err error ) {
		fmt.Fprintf(stderr, "%s\n", err)
		os.Exit(1)
	}
	cmd, err := newDaemonCommand()
	if err != nil {
		onError(err)
	}
	cmd.SetOut(stdout)
	if err := cmd.Execute(); err != nil {
		onError(err)
	}
}
