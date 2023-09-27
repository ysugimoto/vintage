package fastly

import (
	"bytes"
	"context"
	"fmt"
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
	State           vintage.State
	Request         *fsthttp.Request
	BackendRequest  *fsthttp.Request
	BackendResponse *fsthttp.Response
	Response        *fsthttp.Response
	ClientResponse  fsthttp.ResponseWriter
	Geo             *geo.Geo

	// Properties that will be assigned in the process
	clientIdentity string
	OriginalHost   string
}

func NewRuntime(w fsthttp.ResponseWriter, r *fsthttp.Request) (*Runtime, error) {
	g, err := geo.Lookup(net.ParseIP(r.RemoteAddr))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Runtime{
		Runtime:        core.NewRuntime[*Runtime](fsthttp.Header{}, r.Header),
		ClientResponse: w,
		Request:        r,
		Geo:            g,
		OriginalHost:   r.Host,
	}, nil
}

func (r *Runtime) Execute(ctx context.Context) error {
	r.RequestHash = r.Request.URL.String()
	if err := r.Lifecycle(ctx, r); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *Runtime) Proxy(ctx context.Context, backend *vintage.Backend) error {
	resp, err := r.Request.Send(ctx, backend.Backend())
	if err != nil {
		return errors.WithStack(err)
	}

	r.Response = resp
	return nil
}

func (r *Runtime) WriteResponse() (int64, int64, int64, error) {
	r.ClientResponse.Header().Reset(r.BackendResponse.Header)
	written, err := io.Copy(r.ClientResponse, r.BackendResponse.Body)
	if err != nil {
		return 0, 0, 0, errors.WithStack(err)
	}
	// Status line
	statusSize := int64(len(fmt.Sprintf(
		"HTTP/1.1 %s", // @FIXME C@E does not have response protocol so we use fixed value
		fsthttp.StatusText(r.BackendResponse.StatusCode),
	)))
	// Headers
	var headerSize int64
	for key, val := range r.ClientResponse.Header() {
		headerSize += int64(len(key))
		for i := range val {
			headerSize += int64(len(val[i])) + 1
		}
		headerSize--
	}
	// return with header, body, all bytes
	return headerSize, written, statusSize + headerSize + written, nil
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
