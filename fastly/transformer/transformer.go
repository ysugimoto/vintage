package transformer

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/context"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/falco/resolver"
)

type Transformer interface {
	Transform(resolver.Resolver) ([]byte, error)
}

type transformer struct {
	snippets    *context.FastlySnippet
	acls        map[string]*expressionValue
	backends    map[string]*expressionValue
	tables      map[string]*expressionValue
	vars        map[string]*expressionValue
	subroutines map[string]string
	packages    Packages
}

func New(opts ...Option) Transformer {
	t := &transformer{
		acls:        make(map[string]*expressionValue),
		backends:    make(map[string]*expressionValue),
		tables:      make(map[string]*expressionValue),
		vars:        make(map[string]*expressionValue),
		subroutines: make(map[string]string),
		packages: Packages{
			"github.com/ysugimoto/vintage":                {},
			"github.com/ysugimoto/vintage/runtime/fastly": {},
			"github.com/fastly/compute-sdk-go/fsthttp":    {},
		},
	}
	for i := range opts {
		opts[i](t)
	}
	return t
}

func (t *transformer) Transform(rslv resolver.Resolver) ([]byte, error) {
	main, err := rslv.MainVCL()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	vcl, err := parser.New(
		lexer.NewFromString(main.Data, lexer.WithFile(main.Name)),
	).ParseVCL()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if t.snippets != nil {
		for _, snip := range t.snippets.EmbedSnippets() {
			s, err := parser.New(
				lexer.NewFromString(snip.Data, lexer.WithFile(snip.Name)),
			).ParseVCL()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			vcl.Statements = append(s.Statements, vcl.Statements...)
		}
	}

	vcl.Statements, err = t.resolveIncludeStatements(rslv, vcl.Statements, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf, err := t.transform(vcl.Statements)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tmpl := template.Must(template.New("handler").Parse(handlerTemplate))
	var out bytes.Buffer
	if err := tmpl.Execute(&out, map[string]any{
		"Declarations": string(buf),
		"Packages":     t.packages,
		"Subroutines":  t.subroutines,
		"Acls":         t.acls,
		"Backends":     t.backends,
		"Tables":       t.tables,
	}); err != nil {
		return nil, errors.WithStack(err)
	}
	source, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Println(out.String())
		return nil, errors.WithStack(err)
	}
	return source, nil
}

func (t *transformer) transform(statements []ast.Statement) ([]byte, error) {
	var buf bytes.Buffer
	var code []byte
	var err error
	for _, stmt := range statements {
		switch s := stmt.(type) {
		case *ast.AclDeclaration:
			code, err = t.transformAcl(s)
		case *ast.BackendDeclaration:
			code, err = t.transformBackend(s)
		case *ast.DirectorDeclaration:
			code, err = t.transformDirector(s)
		case *ast.TableDeclaration:
			code, err = t.transformTable(s)
		case *ast.SubroutineDeclaration:
			code, err = t.transformSubroutine(s)
		// Currently we don't support penaltybox and ratecounter
		// case *ast.PenaltyboxDeclaration:
		// case *ast.RatecounterDeclaration:
		default:
			err = fmt.Errorf("Unexpected declaration found: %v", s)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(code)
		buf.WriteString(LF)
	}

	return buf.Bytes(), nil
}
