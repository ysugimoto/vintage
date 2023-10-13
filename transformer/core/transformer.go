package core

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/snippets"
	"github.com/ysugimoto/vintage/transformer/value"
	"github.com/ysugimoto/vintage/transformer/variable"
)

type CoreTransformer struct {
	snippets            *snippets.Snippets
	acls                map[string]*value.Value
	backends            map[string]*value.Value
	tables              map[string]*value.Value
	subroutines         map[string]*value.Value
	functionSubroutines map[string]*value.Value
	loggingEndpoints    map[string]string
	Packages            value.Packages

	vars              map[string]*value.Value
	variables         variable.Variables
	regexMatchedStack *RegexMatchedGroupStack
	runtimeName       string
	outputPackageName string

	// Fastly platform support flag
	supportEdgeDictionary bool
}

func NewCoreTransfromer(opts ...TransformOption) *CoreTransformer {
	t := &CoreTransformer{
		acls:                make(map[string]*value.Value),
		backends:            make(map[string]*value.Value),
		tables:              make(map[string]*value.Value),
		vars:                make(map[string]*value.Value),
		subroutines:         make(map[string]*value.Value),
		functionSubroutines: make(map[string]*value.Value),
		loggingEndpoints:    make(map[string]string),
		Packages: value.Packages{
			"github.com/ysugimoto/vintage": {},
		},
		variables:         NewCoreVariables(),
		runtimeName:       "core",
		regexMatchedStack: &RegexMatchedGroupStack{},
		outputPackageName: "main",
	}
	for i := range opts {
		opts[i](t)
	}

	return t
}

func (tf *CoreTransformer) TemplateVariables() map[string]any {
	return map[string]any{
		"Packages":      tf.Packages.Sorted(),
		"Subroutines":   tf.subroutines,
		"Acls":          tf.acls,
		"Backends":      tf.backends,
		"Tables":        tf.tables,
		"OutputPackage": tf.outputPackageName,
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

	var buf bytes.Buffer
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
		for key := range tf.snippets.LoggingEndpoints {
			buf.Write(tf.transformLoggingEndpoint(key))
		}
	}

	vcl.Statements, err = tf.resolveIncludeStatements(rslv, vcl.Statements, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var subroutines []*ast.SubroutineDeclaration
	for _, stmt := range vcl.Statements {
		var code []byte
		switch s := stmt.(type) {
		case *ast.AclDeclaration:
			code = tf.transformAcl(s)
		case *ast.BackendDeclaration:
			code = tf.transformBackend(s)
		case *ast.DirectorDeclaration:
			code = tf.transformDirector(s)
		case *ast.TableDeclaration:
			// On table transformation, we need to care about EdgeDictionary.
			// If ast filename has "Remote.EdgeDictionary" string,
			// the table is an EdgeDictionary which is manged in Fastly remote
			// and then we generate code using compute-sdk-go/configstore package.
			isEdgeDictionary := strings.Contains(s.Meta.Token.File, "Remote.EdgeDictionary")
			if tf.supportEdgeDictionary && isEdgeDictionary {
				code = tf.transformEdgeDictionary(s)
			} else {
				// Otherwise, case of user defined table or could not support EdgeDictionary,
				// transform as Table declaration
				code = tf.transformTable(s)
			}
		case *ast.SubroutineDeclaration:
			// Store subroutine in order to hoisiting other declarations
			subroutines = append(subroutines, s)
		case *ast.ImportStatement:
			// Nothing to to for import statement
			break

		// Currently we don't support penaltybox and ratecounter
		// case *ast.PenaltyboxDeclaration:
		// case *ast.RatecounterDeclaration:
		default:
			err = fmt.Errorf("Unexpected declaration found: %v", s)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if len(code) > 0 {
			buf.Write(code)
		}
	}

	// Transform subroutines after all declaration is transformed
	for i := range subroutines {
		// Reset local variables for each subroutines
		tf.vars = make(map[string]*value.Value)
		code, err := tf.transformSubroutine(subroutines[i])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(code)
		buf.WriteString(lineFeed)
	}

	return buf.Bytes(), nil
}
