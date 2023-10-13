package main

import (
	"strings"

	"github.com/ysugimoto/twist"
)

type Config struct {
	Package string `cli:"p,package" default:"main"`
	Target  string `cli:"t,target" default:"compute"`
	Output  string `cli:"o,output" default:"./vintage.go"`

	ServiceId  string `env:"FASTLY_SERVICE_ID"`
	ApiToken   string `env:"FASTLY_API_TOKEN"`
	EntryPoint string
}

func newConfig(args []string) (*Config, error) {
	var c Config
	if err := twist.Mix(&c, twist.WithCli(args), twist.WithEnv()); err != nil {
		return nil, err
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			continue
		}
		c.EntryPoint = arg
		break
	}

	return &c, nil
}
