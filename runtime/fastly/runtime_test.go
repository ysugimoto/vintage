package fastly

import (
	"testing"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/ysugimoto/vintage"
)

func TestFastlyRuntime(t *testing.T) {
	r := NewRuntime(nil, &fsthttp.Request{
		Header: fsthttp.Header{},
	})
	ok(r.Context())

}

func ok(c *vintage.Context) {
}
