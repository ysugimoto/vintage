package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ysugimoto/falco/remote"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/snippets"
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/transformer/core"
	"github.com/ysugimoto/vintage/transformer/fastly"
)

func main() {
	if err := _main(); err != nil {
		panic(err)
	}
}

func _main() error {
	c, err := newConfig(os.Args[1:])
	if err != nil {
		return errors.WithStack(err)
	}

	if c.Help {
		printHelp()
		os.Exit(1)
	}

	if c.Target != "compute" {
		return fmt.Errorf("Target %s is not supported for now. Only supports 'compute' only\n", c.Target)
	}

	rslv, err := resolver.NewFileResolvers(c.EntryPoint, c.IncludePaths)
	if err != nil {
		return errors.WithStack(err)
	}
	fetcher := remote.NewFastlyApiFetcher(c.ServiceId, c.ApiToken, 10*time.Second)
	s, err := snippets.Fetch(fetcher)
	if err != nil {
		return errors.WithStack(err)
	}
	if err := s.FetchLoggingEndpoint(fetcher); err != nil {
		return errors.WithStack(err)
	}
	options := []core.TransformOption{
		core.WithSnippets(s),
		core.WithOutputPackage(c.Package),
	}
	buf, err := fastly.NewFastlyTransformer(options...).Transform(rslv[0])
	if err != nil {
		return errors.WithStack(err)
	}

	fp, err := os.OpenFile(c.Output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer fp.Close()
	if _, err := fp.Write(buf); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
