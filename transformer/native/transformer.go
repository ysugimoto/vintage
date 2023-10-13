package native

import (
	"bytes"
	"go/format"
	"text/template"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/vintage/transformer/core"
)

type NativeTransformer struct {
	*core.CoreTransformer
}

func NewNativeTransformer(opts ...core.TransformOption) *NativeTransformer {
	// Add Native specific variable resolver
	opts = append(
		opts,
		core.WithVariables(NewNativeVariable()),
		core.WithRuntimeName("native"),
	)
	f := &NativeTransformer{
		core.NewCoreTransfromer(opts...),
	}
	f.CoreTransformer.Packages.Add("github.com/ysugimoto/vintage/runtime/native", "")
	f.CoreTransformer.Packages.Add("net/http", "")
	return f
}

func (tf *NativeTransformer) Transform(rslv resolver.Resolver) ([]byte, error) {
	buf, err := tf.CoreTransformer.Transform(rslv)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	tmpl := template.Must(template.New("native.handler").Parse(handlerTemplate))

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
