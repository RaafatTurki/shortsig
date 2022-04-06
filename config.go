package main

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)


type Config struct {
  Port uint16
  Cmds map[string]string
  Whitelist []string
}

func ParseConfigFile(file_path string) Config {
  var conf Config

  content, readErr := ioutil.ReadFile(file_path)
  if readErr != nil { panic(readErr) }

  _, parseErr := toml.Decode(string(content), &conf)
  if parseErr != nil { panic(readErr) }

  return conf
}

