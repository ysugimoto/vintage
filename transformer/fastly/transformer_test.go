package fastly

import (
	"io/ioutil"
	"testing"

	"github.com/ysugimoto/falco/resolver"
)

func TestTransform(t *testing.T) {
	t.SkipNow()
	rslv, err := resolver.NewFileResolvers("../../example/default.vcl", []string{})
	if err != nil {
		t.Error(err)
		return
	}
	tf := NewFastlyTransformer()
	code, err := tf.Transform(rslv[0])
	if err != nil {
		t.Error(err)
		return
	}
	ioutil.WriteFile("../../../../playground/go-compute-at-edge/vintage.go", code, 0644)
	t.Error(string(code))
}
