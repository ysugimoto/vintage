package core

import (
	"context"

	"github.com/ysugimoto/vintage"
)

type EdgeRuntime interface {
	// Execute starts VCL request lifecycle.
	Execute(context.Context) error

	// CreateBackendRequest is hook point of making bereq from req
	// This hook will be called before calling vcl_pass or vcl_miss
	// Must return raw headers of BackendRequest
	CreateBackendRequest() vintage.RawHeader

	// CreateBackendRequest is hook point of making resp from beresp
	// This hook will be called before calling vcl_deliver
	// Must return raw headers of ClientResponse
	CreateClientResponse() (vintage.RawHeader, error)

	// CreateObjectResponse is hool poing of construct object response.
	// This hook will be called before calling vcl_error.
	// VCL may call without request proxying like "error 400" statement,
	// then object (backend response) is nil, so runtime have to construct from empty.
	// First argument is status code that is set via error [statusCode]
	// Second argument is response body that is set via error [statusCode] [response]
	CreateObjectResponse(int, string) (vintage.RawHeader, error)

	// Proxy is hook point for actual senging origin request.
	// This hook will be called before calling vcl_fetch
	// Must return raw headers of BackendResponse
	Proxy(ctx context.Context, backendName string) (vintage.RawHeader, error)

	// WriteResponse is hook point for sending client response.
	// This hook will be called before calling vcl_log
	// the response is exact 3 length of in64 slice:
	// 0. written header bytes
	// 1. written body bytes
	// 2. written response bytes (include status line)
	WriteResponse() ([3]int64, error)

	// Cache if hook point for lookup cache
	// @FIXME: Edge Runtime caching (e.g SimpleCache in Compute@Edge) is limited, so currently unsupported)
	// Cache(ctx context.Runtime, key string) error
}
