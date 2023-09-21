package transformer

import (
	"bytes"
	"fmt"

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
	snippets *context.FastlySnippet
	acls     map[string]string
	backends map[string]string
	tables   map[string]string
	vars     map[string]string
}

func New(opts ...Option) Transformer {
	t := &transformer{
		acls:     make(map[string]string),
		backends: make(map[string]string),
		tables:   make(map[string]string),
		vars:     make(map[string]string),
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
	return buf, nil
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
	}

	return buf.Bytes(), nil
}
