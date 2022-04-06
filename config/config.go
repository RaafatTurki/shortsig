package config

import (
	"io/ioutil"
	. "shortsig/log"

	"github.com/BurntSushi/toml"
)

type Routine struct {
  Linux string
  Windows string
  Darwin string
  Other string
}

type Config struct {
  Port uint16
  Routines map[string]Routine
  Whitelist []string
}

func ParseConfigFile(file_path string) Config {
  var conf Config

  content, readErr := ioutil.ReadFile(file_path)
  Assert(readErr)

  _, parseErr := toml.Decode(string(content), &conf)
  Assert(parseErr)

  return conf
}

