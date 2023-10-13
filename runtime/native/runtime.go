package native

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/runtime/core"
)

type Runtime struct {
	*core.Runtime[*Runtime]
	State           vintage.State
	Request         *http.Request
	BackendRequest  *http.Request
	BackendResponse *http.Response
	Response        *http.Response
	ClientResponse  http.ResponseWriter
}

func NewRuntime(w http.ResponseWriter, r *http.Request) (*Runtime, error) {
	rt := &Runtime{
		Runtime:        core.NewRuntime[*Runtime](r.Header),
		ClientResponse: w,
		Request:        r,
	}
	rt.OriginalHost = r.Host
	return rt, nil
}

func (r *Runtime) Release() {
	// Release backend response body if exists
	if r.BackendResponse != nil {
		r.BackendResponse.Body.Close()
	}
	// Release client response body if exists
	if r.Response != nil {
		r.Response.Body.Close()
	}
}

func (r *Runtime) Execute(ctx context.Context) error {
	r.RequestHash = r.Request.URL.String()
	idx := strings.LastIndex(r.Request.RemoteAddr, ":")
	if idx == -1 {
		r.ClientIp = net.ParseIP(r.Request.RemoteAddr)
	} else {
		r.ClientIp = net.ParseIP(r.Request.RemoteAddr[:idx])
	}

	// Release some pointer resources in order to prevent memory leak
	defer r.Release()

	if err := r.Lifecycle(ctx, r); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Runtime) Proxy(ctx context.Context, backendName string) (vintage.RawHeader, error) {
	fmt.Printf("Proxy request send to %s\n", backendName)
	backend, ok := r.Backends[backendName]
	if !ok {
		return nil, errors.WithStack(
			fmt.Errorf("Backend %s is not defined", backendName),
		)
	}

	scheme := "http"
	port := 80
	if backend.SSL {
		scheme = "https"
		port = 443
	}

	url := fmt.Sprintf("%s://%s:%d%s", scheme, backend.Host, port, r.BackendRequest.URL.Path)
	if query := r.BackendRequest.URL.Query().Encode(); query != "" {
		url += "?" + query
	}

	c, timeout := context.WithTimeout(ctx, backend.FirstByteTimeout)
	defer timeout()

	req, err := http.NewRequestWithContext(c, r.BackendRequest.Method, url, r.BackendRequest.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header = r.BackendRequest.Header.Clone()
	if backend.AlwaysUseHostHeader {
		req.Header.Set("Host", backend.Host)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	r.BackendResponse = resp
	return vintage.RawHeader(resp.Header), nil
}

func (r *Runtime) WriteResponse() (int64, int64, int64, error) {
	h := r.ClientResponse.Header()
	for key, val := range r.BackendResponse.Header {
		h[key] = val
	}
	written, err := io.Copy(r.ClientResponse, r.BackendResponse.Body)
	if err != nil {
		return 0, 0, 0, errors.WithStack(err)
	}
	// Status line
	statusSize := int64(len(fmt.Sprintf(
		"HTTP/1.1 %s", // @FIXME C@E does not have response protocol so we use fixed value
		http.StatusText(r.BackendResponse.StatusCode),
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
	r.BackendRequest = r.Request.Clone(r.Request.Context())
	return vintage.RawHeader(r.BackendRequest.Header)
}

func (r *Runtime) CreateClientResponse() (vintage.RawHeader, error) {
	beresp := r.BackendResponse

	if beresp == nil {
		return nil, errors.WithStack(
			fmt.Errorf("BackendResponse have not created yet"),
		)
	}

	// Read and rewind backend response
	var body bytes.Buffer
	if _, err := body.ReadFrom(beresp.Body); err != nil {
		return nil, errors.WithStack(err)
	}
	beresp.Body = io.NopCloser(bytes.NewReader(body.Bytes()))

	// Clone backend response
	r.Response = &http.Response{
		Request:    r.BackendRequest,
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
	r.BackendResponse = &http.Response{
		StatusCode: statusCode,
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader(response)),
	}

	return vintage.RawHeader(r.BackendResponse.Header), nil
}
