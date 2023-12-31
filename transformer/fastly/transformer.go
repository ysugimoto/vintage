package fastly

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/vintage/transformer/core"
)

type FastlyTransformer struct {
	*core.CoreTransformer
}

func NewFastlyTransformer(opts ...core.TransformOption) *FastlyTransformer {
	// Add Fastly specific variable resolver
	opts = append(
		opts,
		core.WithVariables(NewFastlyVariable()),
		core.WithFastlyPlatform(),
		core.WithRuntimeName("fastly"),
	)
	f := &FastlyTransformer{
		core.NewCoreTransfromer(opts...),
	}
	f.CoreTransformer.Packages.Add("github.com/ysugimoto/vintage/runtime/fastly", "")
	f.CoreTransformer.Packages.Add("github.com/fastly/compute-sdk-go/fsthttp", "")
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
		return nil, errors.WithStack(err)
	}
	return source, nil
}
