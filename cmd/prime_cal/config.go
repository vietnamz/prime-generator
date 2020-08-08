package main

import "github.com/spf13/pflag"

type Config struct {
	logLevel  string
	logEnable bool
	port      string
	Hosts []string
	configFile string
	daemonConfig *Config
	flags *pflag.FlagSet
}

func NewConfig() *Config {
	return &Config{
		logEnable: true,
		logLevel:  "debug",
		port:      "8080",
		Hosts: []string{ "127.0.0.1:8080" },
	}
}

func (o *Config) InstallFlags( flags *pflag.FlagSet)  {
	flags.BoolVarP( &o.logEnable, "debug", "D", false, "Enable debug mode")
	flags.StringVarP( &o.logLevel, "log-level", "l", "info", `Set the logging level ("debug"|"info"|"warn"|"error"|"fatal")`)
}