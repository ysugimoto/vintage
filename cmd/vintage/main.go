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
	"github.com/ysugimoto/vintage/transformer/native"
)

const (
	targetCompute = "compute"
	targetNative  = "native"
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

	var transformer core.Transformer
	switch c.Target {
	case targetCompute:
		transformer = fastly.NewFastlyTransformer(options...)
	case targetNative:
		transformer = native.NewNativeTransformer(options...)
	default:
		return fmt.Errorf(`Target %s is not supported. Only supports "compute" or "native"`+"\n", c.Target)
	}

	buf, err := transformer.Transform(rslv[0])
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
