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
	rslv, err := resolver.NewFileResolvers("../../example/default.vcl", []string{})
	if err != nil {
		panic(err)
	}
	fetcher := remote.NewFastlyApiFetcher(
		os.Getenv("FASTLY_SERVICE_ID"),
		os.Getenv("FASTLY_API_TOKEN"),
		10*time.Second,
	)
	s, err := snippets.Fetch(fetcher)
	if err != nil {
		panic(err)
	}
	if err := s.FetchLoggingEndpoint(fetcher); err != nil {
		panic(err)
	}
	buf, err := fastly.NewFastlyTransformer(core.WithSnippets(s)).Transform(rslv[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
