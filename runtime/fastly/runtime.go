package fastly

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	cache "github.com/fastly/compute-sdk-go/cache/core"
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
	idx := strings.LastIndex(r.Request.RemoteAddr, ":")
	if idx == -1 {
		r.ClientIp = net.ParseIP(r.Request.RemoteAddr)
	} else {
		r.ClientIp = net.ParseIP(r.Request.RemoteAddr[:idx])
	}
	if err := r.Lifecycle(ctx, r); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r *Runtime) Proxy(ctx context.Context, backendName string) (vintage.RawHeader, error) {
	fmt.Printf("Proxy request send to %s\n", backendName)
	resp, err := r.BackendRequest.Send(ctx, backendName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	fmt.Printf("Proxy request responds status code %d\n", resp.StatusCode)

	r.BackendResponse = resp
	return vintage.RawHeader(resp.Header), nil
}

func (r *Runtime) WriteResponse() ([3]int64, error) {
	r.ClientResponse.Header().Reset(r.BackendResponse.Header)
	written, err := io.Copy(r.ClientResponse, r.BackendResponse.Body)
	if err != nil {
		return [3]int64{}, errors.WithStack(err)
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
	return [3]int64{
		headerSize,
		written,
		statusSize + headerSize + written,
	}, nil
}

func (r *Runtime) CreateBackendRequest() vintage.RawHeader {
	r.BackendRequest = r.Request.Clone()
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
	r.Response = &fsthttp.Response{
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
	r.BackendResponse = &fsthttp.Response{
		StatusCode: statusCode,
		Header:     fsthttp.Header{},
		Body:       io.NopCloser(strings.NewReader(response)),
	}

	return vintage.RawHeader(r.BackendResponse.Header), nil
}

func (r *Runtime) LookupCache() (bool, error) {
	tx, err := cache.NewTransaction([]byte(r.RequestHash), cache.LookupOptions{
		RequestHeaders: r.Request.Header,
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	defer tx.Close()

	found, err := tx.Found()
	if err != nil {
		if err == cache.ErrNotFound {
			r.BackendResponse.Header.Set("X-Cache", "MISS")
			return false, nil
		}
		return false, errors.WithStack(err)
	}
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(found.Body); err != nil {
		return false, errors.WithStack(err)
	}
	if _, err := r.CreateObjectResponse(fsthttp.StatusOK, buf.String()); err != nil {
		return false, errors.WithStack(err)
	}
	r.BackendResponse.Header.Set("X-Cache-Hits", fmt.Sprint(found.Hits))
	r.BackendResponse.Header.Set("Age", fmt.Sprint(found.Age.Seconds()))
	r.BackendResponse.Header.Set("X-Cache", "HIT")

	return true, nil
}

func (r *Runtime) SaveCache() error {
	// If backend response is not set, skip
	if r.BackendResponse == nil {
		return nil
	}
	// Set backend response is cacheable or not
	r.BackendResponseCachable = r.ObjectCacheable()
	if !r.BackendResponseCachable {
		return nil
	}
	// Skip if backend response TTL is set to zero
	if r.BackendResponseTTL.Seconds() == 0 {
		return nil
	}

	tx, err := cache.NewTransaction([]byte(r.RequestHash), cache.LookupOptions{
		RequestHeaders: r.Request.Header,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Close()

	if _, err := tx.Found(); err != nil {
		// Insert new cache if cache is not found
		if err != cache.ErrNotFound {
			return errors.WithStack(err)
		}

		// Insert new cache
		if !tx.MustInsert() {
			return errors.WithStack(err)
		}
		contents, err := r.ResponseBody(r.BackendResponse)
		if err != nil {
			return errors.WithStack(err)
		}
		w, _, err := tx.InsertAndStreamBack(cache.WriteOptions{
			TTL:            r.BackendResponseTTL,
			RequestHeaders: r.Request.Header,
			SurrogateKeys:  r.BackendResponse.Header.Values("Surrogate-Key"),
			Length:         uint64(len(contents)),
		})
		if err != nil {
			return errors.WithStack(err)
		}
		if _, err := io.WriteString(w, contents); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}

	// Overwrite cache if needed
	if !tx.MustInsertOrUpdate() {
		return nil
	}
	contents, err := r.ResponseBody(r.BackendResponse)
	if err != nil {
		return errors.WithStack(err)
	}
	w, err := tx.Insert(cache.WriteOptions{
		TTL:            r.BackendResponseTTL,
		RequestHeaders: r.Request.Header,
		SurrogateKeys:  r.BackendResponse.Header.Values("Surrogate-Key"),
		Length:         uint64(len(contents)),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if _, err := io.WriteString(w, contents); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
