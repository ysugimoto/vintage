package fastly

import (
	"bytes"
	"context"
	"io"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

type FastlyRuntime struct {
	vctx            *vintage.Context
	Request         *fsthttp.Request
	Response        fsthttp.ResponseWriter
	BackendResponse *fsthttp.Response
}

func NewRuntime(w fsthttp.ResponseWriter, r *fsthttp.Request) *FastlyRuntime {
	return &FastlyRuntime{
		vctx:     vintage.NewContext(fsthttp.Header{}, r.Header),
		Request:  r,
		Response: w,
	}
}

func (r *FastlyRuntime) Context() *vintage.Context {
	return r.vctx
}

func (r *FastlyRuntime) Proxy(ctx context.Context, backend string) (*fsthttp.Response, error) {
	resp, err := r.Request.Send(ctx, backend)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// To avoid memory leak, read and rewind response body
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, errors.WithStack(err)
	}
	resp.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
	r.BackendResponse = resp
	r.vctx.ResponseHeader = vintage.NewHeader(resp.Header)
	return resp, nil
}
