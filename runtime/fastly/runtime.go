package fastly

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strings"

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
	clientIdentity  string
}

func (r *Runtime) Context() *core.Runtime[*Runtime] {
	return r.Runtime
}

func NewRuntime(w fsthttp.ResponseWriter, r *fsthttp.Request) (*Runtime, error) {
	g, err := geo.Lookup(net.ParseIP(r.RemoteAddr))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	rt := &Runtime{
		Runtime:        core.NewRuntime[*Runtime](r.Header),
		ClientResponse: w,
		Request:        r,
		Geo:            g,
	}
	rt.OriginalHost = r.Host
	return rt, nil
}

func (r *Runtime) Execute(ctx context.Context) error {
	r.RequestHash = r.Request.URL.String()
	if err := r.Lifecycle(ctx, r); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Runtime) Proxy(ctx context.Context, backend *vintage.Backend) (vintage.RawHeader, error) {
	fmt.Printf("Proxy request send to %s\n", backend.Backend())
	resp, err := r.BackendRequest.Send(ctx, backend.Backend())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("Proxy request responds status code %d\n", resp.StatusCode)

	r.BackendResponse = resp
	return vintage.RawHeader(resp.Header), nil
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

func (r *Runtime) CreateBackendRequest() vintage.RawHeader {
	r.BackendRequest = r.Request.Clone()
	return vintage.RawHeader(r.BackendRequest.Header)
}

func (r *Runtime) CreateClientResponse() (vintage.RawHeader, error) {
	beresp := r.BackendResponse

	if beresp == nil {
		fmt.Println("backend response is null")
	}

	// Read and rewind backend response
	var body bytes.Buffer
	if _, err := body.ReadFrom(beresp.Body); err != nil {
		return nil, errors.WithStack(err)
	}
	beresp.Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	// Clone backend response
	r.Response = &fsthttp.Response{
		Request:    r.BackendRequest,
		Backend:    r.Backend.Backend(),
		StatusCode: beresp.StatusCode,
		Header:     beresp.Header.Clone(),
		Body:       io.NopCloser(bytes.NewReader(body.Bytes())),
	}
	return vintage.RawHeader(r.Response.Header), nil
}

func (r *Runtime) CreateObjectResponse(statusCode int, response string) (vintage.RawHeader, error) {
	// Guard process that backend response already exists
	if r.BackendResponse != nil {
		return vintage.RawHeader(r.BackendResponse.Header), nil
	}

	r.IsLocallyGenerated = true
	r.BackendResponse = &fsthttp.Response{
		StatusCode: statusCode,
		Header:     fsthttp.Header{},
		Body:       io.NopCloser(strings.NewReader(response)),
	}

	return vintage.RawHeader(r.BackendResponse.Header), nil

}
