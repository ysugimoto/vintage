package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ysugimoto/falco/remote"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/snippets"
	"github.com/ysugimoto/vintage/transformer/core"
	"github.com/ysugimoto/vintage/transformer/fastly"
)

func main() {
	c, err := newConfig(os.Args[1:])
	if err != nil {
		panic(err)
	}

	if c.Target != "compute" {
		fmt.Fprintf(os.Stderr, "Target %s is not supported for now. Only supports 'compute' only\n", c.Target)
		os.Exit(1)
	}

	rslv, err := resolver.NewFileResolvers(c.EntryPoint, []string{})
	if err != nil {
		panic(err)
	}
	fetcher := remote.NewFastlyApiFetcher(c.ServiceId, c.ApiToken, 10*time.Second)
	s, err := snippets.Fetch(fetcher)
	if err != nil {
		panic(err)
	}
	if err := s.FetchLoggingEndpoint(fetcher); err != nil {
		panic(err)
	}
	buf, err := fastly.NewFastlyTransformer(
		core.WithSnippets(s),
		core.WithOutputPackage(c.Package),
	).Transform(rslv[0])
	if err != nil {
		panic(err)
	}

	fp, err := os.OpenFile(c.Output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fp.Write(buf) // nolint:errcheck
}
