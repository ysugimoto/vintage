package native

import (
	"os"
	"testing"
	"time"

	"github.com/ysugimoto/falco/remote"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/snippets"
	"github.com/ysugimoto/vintage/transformer/core"
)

func TestTransformVCL(t *testing.T) {
	rslv, err := resolver.NewFileResolvers("../../example/default.vcl", []string{})
	if err != nil {
		t.Errorf("Failed to resolve main VCL: %s", err)
		return
	}

	fetcher := remote.NewFastlyApiFetcher(
		os.Getenv("FASTLY_SERVICE_ID"),
		os.Getenv("FASTLY_API_TOKEN"),
		10*time.Second,
	)
	s, err := snippets.Fetch(fetcher)
	if err != nil {
		t.Errorf("Failed to fetch snippets: %s", err)
		return
	}
	if err := s.FetchLoggingEndpoint(fetcher); err != nil {
		t.Errorf("Failed to fetch logging endpoints: %s", err)
		return
	}
	_, err = NewNativeTransformer(core.WithSnippets(s)).Transform(rslv[0])
	if err != nil {
		t.Errorf("Failed to transform VCL: %s", err)
	}
}
