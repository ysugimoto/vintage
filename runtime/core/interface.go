package core

import (
	"context"

	"github.com/ysugimoto/vintage"
)

type EdgeRuntime interface {
	// CreateBackendRequest is hook point of making bereq from req
	// This hook will be called before calling vcl_pass or vcl_miss
	CreateBackendRequest()

	// CreateBackendRequest is hook point of making resp from beresp
	// This hook will be called before calling vcl_deliver
	CreateClientResponse() error

	// Proxy is hook point for actual senging origin request.
	// This hook will be called before calling vcl_fetch
	Proxy(ctx context.Context, backend *vintage.Backend) error

	// WriteResponse is hook point for sending client response.
	// This hook will be called before calling vcl_log
	// the response values are:
	// 1. written header bytes
	// 2. written body bytes
	// 2. written response bytes (include status line)
	WriteResponse() (int64, int64, int64, error)

	// Cache if hook point for lookup cache
	// @FIXME: Edge Runtime caching (e.g SimpleCache in Compute@Edge) is limited, so currently unsupported)
	// Cache(ctx context.Runtime, key string) error
}
