package function

import (
	"context"

	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/runtime/core"
)

// TestRuntime uses for builtin function testings.
// Implements EdgeRuntime interface but returns fake values.
type TestRuntime struct {
	*core.Runtime[*TestRuntime]
}

func newTestRuntime() *core.Runtime[*TestRuntime] {
	rt := &TestRuntime{
		core.NewRuntime[*TestRuntime](map[string][]string{}),
	}
	return rt.Runtime
}

// Fake implementation of EdgeRuntime interface
func (r *TestRuntime) Execute(ctx context.Context) error {
	return nil
}

func (r *TestRuntime) CreateBackendRequest() vintage.RawHeader {
	return vintage.RawHeader(map[string][]string{})
}

func (r *TestRuntime) CreateClientResponse() (vintage.RawHeader, error) {
	return vintage.RawHeader(map[string][]string{}), nil
}

func (r *TestRuntime) CreateObjectResponse(statusCode int, body string) (vintage.RawHeader, error) {
	return vintage.RawHeader(map[string][]string{}), nil
}

func (r *TestRuntime) Proxy(ctx context.Context, backend string) (vintage.RawHeader, error) {
	return vintage.RawHeader(map[string][]string{}), nil
}

func (r *TestRuntime) WriteResponse() ([3]int64, error) {
	return [3]int64{0, 0, 0}, nil
}

func (r *TestRuntime) LookupCache() (bool, error) {
	return false, nil
}
func (r *TestRuntime) SaveCache() error {
	return nil
}
