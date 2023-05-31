package config

import (
	"shortsig/core/log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Routine struct {
  Linux string
  Windows string
  Darwin string
  Other string
}

type Routines map[string]Routine

type Config struct {
  Port uint16
  Routines Routines
  Whitelist []string
  ShowVersion bool
}

func ParseConfigs(configFilePaths... string) Config {
  // default values
  viper.SetDefault("port", 3090)

  // defining config file attributes
  viper.SetConfigName("config")
  viper.SetConfigType("toml")

  for _, configFilePath := range configFilePaths {
    log.PrintConsole(log.DEBUG, "added custom config path at %s", configFilePath)
    viper.AddConfigPath(configFilePath)
  }

  viper.AddConfigPath("/etc/shortsig")
  viper.AddConfigPath("$XDG_CONFIG_HOME/shortsig")
  viper.AddConfigPath("$HOME/.config/shortsig")
  viper.AddConfigPath("/home/$USER/.config/shortsig")

  err := viper.ReadInConfig()
  log.PanicErr(err)

  // parse config file
  var conf Config
  unmarshailErr := viper.Unmarshal(&conf)
  log.PanicErr(unmarshailErr)

  // flags
  pflag.Uint16VarP(&conf.Port, "port", "p", conf.Port, "set TCP port")
  pflag.BoolVarP(&conf.ShowVersion, "version", "v", false, "display version information")
  pflag.Parse()

  return conf
}
