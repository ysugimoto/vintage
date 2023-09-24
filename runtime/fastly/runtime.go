package fastly

import (
	"context"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

type Runtime struct {
	*vintage.Context[*Runtime]
	State    vintage.State
	Request  *fsthttp.Request
	Response *fsthttp.Response
}

func NewRuntime(r *fsthttp.Request) *Runtime {
	return &Runtime{
		Context: vintage.NewContext[*Runtime](fsthttp.Header{}, r.Header),
		Request: r,
	}
}

func (r *Runtime) Execute(ctx context.Context) (*fsthttp.Response, error) {
	r.RequestHash = r.Request.URL.String()
	if err := r.Context.Lifecycle(ctx, r); err != nil {
		return nil, errors.WithStack(err)
	}
	return r.Response, nil
}

func (r *Runtime) Proxy(ctx context.Context, backend *vintage.Backend) error {
	resp, err := r.Request.Send(ctx, backend.Backend())
	if err != nil {
		return errors.WithStack(err)
	}

	r.Response = resp
	return nil
}
