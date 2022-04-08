package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfig(t *testing.T) {
  a := assert.New(t)

  conf := ParseConfigs("./testdata")

  // port
  a.Equal(conf.Port, uint16(4321))

  // routines
  a.Equal(conf.Routines["ls"].Linux, "ls")
}
