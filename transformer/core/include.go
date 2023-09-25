package core

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/falco/resolver"
)

func (tf *CoreTransformer) resolveIncludeStatements(
	rslv resolver.Resolver, statements []ast.Statement,
	isRoot bool,
) ([]ast.Statement, error) {

	var resolved []ast.Statement
	for _, stmt := range statements {
		include, ok := stmt.(*ast.IncludeStatement)
		if !ok {
			resolved = append(resolved, stmt)
			continue
		}

		if strings.HasPrefix(include.Module.Value, "snippet::") {
			if included, err := tf.includeSnippet(include, isRoot); err != nil {
				return nil, errors.WithStack(err)
			} else {
				resolved = append(resolved, included...)
			}
			continue
		}
		included, err := tf.includeFile(rslv, include, isRoot)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		recursive, err := tf.resolveIncludeStatements(rslv, included, isRoot)
		if err != nil {
			return nil, err
		}
		resolved = append(resolved, recursive...)
	}

	return resolved, nil
}

func (tf *CoreTransformer) includeSnippet(include *ast.IncludeStatement, isRoot bool) ([]ast.Statement, error) {
	if tf.snippets == nil {
		return nil, errors.WithStack(
			fmt.Errorf("Remote snippet is not found. Did you run with '-r' option?"),
		)
	}
	snip, ok := tf.snippets.IncludeSnippets[strings.TrimPrefix(include.Module.Value, "snippet::")]
	if !ok {
		return nil, errors.WithStack(
			fmt.Errorf("Failed to include VCL snippets '%s'", include.Module.Value),
		)
	}
	if isRoot {
		return loadRootVCL(include.Module.Value, snip.Data)
	}
	return loadStatementVCL(include.Module.Value, snip.Data)
}

func (tf *CoreTransformer) includeFile(
	rslv resolver.Resolver,
	include *ast.IncludeStatement,
	isRoot bool,
) ([]ast.Statement, error) {

	module, err := rslv.Resolve(include)
	if err != nil {
		return nil, errors.WithStack(
			fmt.Errorf("Failed to include VCL module '%s'", include.Module.Value),
		)
	}

	if isRoot {
		return loadRootVCL(module.Name, module.Data)
	}
	return loadStatementVCL(module.Name, module.Data)
}

func loadRootVCL(name, content string) ([]ast.Statement, error) {
	lx := lexer.NewFromString(content, lexer.WithFile(name))
	vcl, err := parser.New(lx).ParseVCL()
	if err != nil {
		return nil, err
	}
	return vcl.Statements, nil
}

func loadStatementVCL(name, content string) ([]ast.Statement, error) {
	lx := lexer.NewFromString(content, lexer.WithFile(name))
	vcl, err := parser.New(lx).ParseSnippetVCL()
	if err != nil {
		return nil, err
	}
	return vcl, nil
}
