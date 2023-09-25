package fastly

import (
	"bytes"
	"context"
	"io"
	"net"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/fastly/compute-sdk-go/geo"
	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/runtime/core"
)

type Runtime struct {
	*core.Runtime[*Runtime]
	State           core.State
	Request         *fsthttp.Request
	BackendRequest  *fsthttp.Request
	BackendResponse *fsthttp.Response
	Response        *fsthttp.Response
	Geo             *geo.Geo

	// Properties that will be assigned in the process
	clientIdentity string
}

func NewRuntime(r *fsthttp.Request) (*Runtime, error) {
	g, err := geo.Lookup(net.ParseIP(r.RemoteAddr))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Runtime{
		Runtime: core.NewRuntime[*Runtime](fsthttp.Header{}, r.Header),
		Request: r,
		Geo:     g,
	}, nil
}

func (r *Runtime) Execute(ctx context.Context) (*fsthttp.Response, error) {
	r.RequestHash = r.Request.URL.String()
	if err := r.Lifecycle(ctx, r); err != nil {
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

func (r *Runtime) CreateBackendRequest() {
	r.BackendRequest = r.Request.Clone()
}

func (r *Runtime) CreateClientResponse() error {
	beresp := r.BackendResponse

	// Read and rewind backend response
	var body bytes.Buffer
	if _, err := body.ReadFrom(beresp.Body); err != nil {
		return errors.WithStack(err)
	}
	beresp.Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	// Clone backend response
	r.Response = &fsthttp.Response{
		Request:    r.BackendRequest,
		Backend:    r.Backend.Backend(),
		StatusCode: beresp.StatusCode,
		Header:     beresp.Header,
		Body:       io.NopCloser(bytes.NewReader(body.Bytes())),
	}
	return nil
}
