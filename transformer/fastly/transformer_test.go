package fastly

import (
	"io/ioutil"
	"testing"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/resolver"
)

func TestTransform(t *testing.T) {
	// t.SkipNow()
	rslv, err := resolver.NewFileResolvers("../../../../works/ise-cdn/dist/default.vcl", []string{})
	if err != nil {
		t.Error(err)
		return
	}
	tf := NewFastlyTransformer()
	code, err := tf.Transform(rslv[0])
	if err != nil {
		t.Error(errors.Cause(err))
		return
	}
	ioutil.WriteFile("../../../../playground/go-compute-at-edge/ise-cdn.go", code, 0644)
	t.Error(string(code))
}
