package core

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/context"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/vintage/transformer/value"
	"github.com/ysugimoto/vintage/transformer/variable"
)

type CoreTransformer struct {
	snippets            *context.FastlySnippet
	acls                map[string]*value.Value
	backends            map[string]*value.Value
	tables              map[string]*value.Value
	subroutines         map[string]*value.Value
	functionSubroutines map[string]*value.Value
	Packages            value.Packages

	vars      map[string]*value.Value
	variables variable.Variables
}

func NewCoreTransfromer(opts ...TransformOption) *CoreTransformer {
	t := &CoreTransformer{
		acls:                make(map[string]*value.Value),
		backends:            make(map[string]*value.Value),
		tables:              make(map[string]*value.Value),
		vars:                make(map[string]*value.Value),
		subroutines:         make(map[string]*value.Value),
		functionSubroutines: make(map[string]*value.Value), Packages: value.Packages{
			"github.com/ysugimoto/vintage": {},
		},
	}
	for i := range opts {
		opts[i](t)
	}
	return t
}

func (tf *CoreTransformer) TemplateVariables() map[string]any {
	return map[string]any{
		"Packages":    tf.Packages.Sorted(),
		"Subroutines": tf.subroutines,
		"Acls":        tf.acls,
		"Backends":    tf.backends,
		"Tables":      tf.tables,
	}
}

func (tf *CoreTransformer) Transform(rslv resolver.Resolver) ([]byte, error) {
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

	if tf.snippets != nil {
		for _, snip := range tf.snippets.EmbedSnippets() {
			s, err := parser.New(
				lexer.NewFromString(snip.Data, lexer.WithFile(snip.Name)),
			).ParseVCL()
			if err != nil {
				return nil, errors.WithStack(err)
			}
			vcl.Statements = append(s.Statements, vcl.Statements...)
		}
	}

	vcl.Statements, err = tf.resolveIncludeStatements(rslv, vcl.Statements, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var buf bytes.Buffer
	var code []byte
	for _, stmt := range vcl.Statements {
		switch s := stmt.(type) {
		case *ast.AclDeclaration:
			code, err = tf.transformAcl(s)
		case *ast.BackendDeclaration:
			code, err = tf.transformBackend(s)
		case *ast.DirectorDeclaration:
			code, err = tf.transformDirector(s)
		case *ast.TableDeclaration:
			code, err = tf.transformTable(s)
		case *ast.SubroutineDeclaration:
			// Reset local variables for each subroutines
			tf.vars = make(map[string]*value.Value)
			code, err = tf.transformSubroutine(s)
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
		buf.WriteString(lineFeed)
	}

	return buf.Bytes(), nil
}
