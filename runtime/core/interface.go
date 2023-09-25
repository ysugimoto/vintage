package core

import (
	"context"

	"github.com/ysugimoto/vintage"
)

type EdgeRuntime interface {
	CreateBackendRequest()
	CreateClientResponse() error
	Proxy(ctx context.Context, backend *vintage.Backend) error
	// Cache(ctx context.Runtime, key string) error
}
