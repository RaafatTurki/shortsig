package config

import (
	. "shortsig/log"

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
}

func ParseConfigs() Config {
  // default values
  viper.SetDefault("port", 3090)

  // defining config file attributes
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath("/etc/shortsig")
  viper.AddConfigPath("$XDG_CONFIG_HOME/shortsig")
  viper.AddConfigPath("$HOME/.config/shortsig")
  viper.AddConfigPath("/home/$USER/.config/shortsig")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  Assert(err)

  // parse config file
  var conf Config
  unmarshailErr := viper.Unmarshal(&conf)
  Assert(unmarshailErr)

  // flags
  pflag.Uint16VarP(&conf.Port, "port", "p", conf.Port, "help message for flagname")
  pflag.Parse()

  return conf
}
