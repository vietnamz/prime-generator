package daemon

import (
	"github.com/spf13/pflag"
)

type Config struct {
	// is log level ("debug"|"info"|"warn"|"error"|"fatal")
	LogLevel  string
	// is log enabled or not
	LogEnable bool
	// port to bind to api service.
	Port      string
	// host to be attach.
	Hosts []string
	// Max thread to execute the prime seive.
	MaxThread int
	// The size of sieve in Kb. default 128Kb.
	MaxSieveSize int
	// The config file. Not used for now.
	ConfigFile string
	// cobra flags.
	Flags *pflag.FlagSet
}

func NewDaemonConfig() *Config {
	return &Config{
		LogEnable: true,
		LogLevel:  "debug",
		Port:      "8080",
		Hosts: []string{ "0.0.0.0:8080" },
		MaxThread: 0,
		MaxSieveSize: 0,
	}
}

func (o *Config) InstallFlags( flags *pflag.FlagSet)  {
	flags.BoolVarP( &o.LogEnable, "debug", "D", false, "Enable debug mode")
	flags.StringVarP( &o.LogLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
}

// Daemon is entry to keep all the backend services to serve the API.
type Daemon struct {
	 PrimeSrv *PrimeService
	 config *Config
}

// Constructor.
func NewDaemon(cfg *Config) *Daemon {
	return &Daemon{
		config: cfg,
	}
}

// initialize all the backend services.
// Support:
//			+ Prime Generator service: to return a sample prime number.
func (d *Daemon) Init() error{
	d.PrimeSrv = newPrimeService(d.config)
	return nil
}