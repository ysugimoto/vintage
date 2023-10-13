package main

import (
	"path/filepath"
	"strings"

	"github.com/ysugimoto/twist"
)

type Config struct {
	Package      string   `cli:"p,package" default:"main"`
	Target       string   `cli:"t,target" default:"compute"`
	Output       string   `cli:"o,output" default:"./vintage.go"`
	Help         bool     `cli:"h,help"`
	IncludePaths []string `cli:"I,include_path"`
	Overwrite    bool     `cli:"overwrite"`

	ServiceId  string `env:"FASTLY_SERVICE_ID"`
	ApiToken   string `env:"FASTLY_API_TOKEN"`
	EntryPoint string
}

func newConfig(args []string) (*Config, error) {
	var c Config
	err := twist.Mix(&c, twist.WithCli(args), twist.WithEnv())
	if err != nil {
		return nil, err
	}

	var commands []string
	for i := 0; i < len(args); i++ {
		if _, ok := needValueOptions[args[i]]; ok {
			i++
			continue
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}
		commands = append(commands, args[i])
	}

	if len(commands) > 0 {
		c.EntryPoint = commands[0]
		c.IncludePaths = append([]string{filepath.Dir(c.EntryPoint)}, c.IncludePaths...)
	}
	c.IncludePaths, err = toAbsolutePaths(c.IncludePaths)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

var needValueOptions = map[string]struct{}{
	"-I":             {},
	"--include_path": {},
	"-t":             {},
	"--target":       {},
	"-o":             {},
	"--output":       {},
}

func toAbsolutePaths(paths []string) ([]string, error) {
	absPaths := make([]string, len(paths))
	for i := range paths {
		a, err := filepath.Abs(paths[i])
		if err != nil {
			return nil, err
		}
		absPaths[i] = a
	}
	return absPaths, nil
}
