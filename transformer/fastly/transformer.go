package fastly

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/vintage/transformer"
)

type FastlyTransformer struct {
	*transformer.CoreTransformer
}

func NewFastlyTransformer(opts ...transformer.Option) *FastlyTransformer {
	f := &FastlyTransformer{
		transformer.NewCoreTransfromer(opts...),
	}
	f.Packages.Add("github.com/ysugimoto/vintage/runtime/fastly", "")
	f.Packages.Add("github.com/fastly/compute-sdk-go/fsthttp", "")
	return f
}

func (tf *FastlyTransformer) Transform(rslv resolver.Resolver) ([]byte, error) {
	buf, err := tf.CoreTransformer.Transform(rslv)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tmpl := template.Must(template.New("fastly.handler").Parse(handlerTemplate))

	var out bytes.Buffer
	vars := tf.CoreTransformer.TemplateVariables()
	vars["Declarations"] = string(buf)
	if err := tmpl.Execute(&out, vars); err != nil {
		return nil, errors.WithStack(err)
	}
	source, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Println(out.String())
		return nil, errors.WithStack(err)
	}
	return source, nil
}
