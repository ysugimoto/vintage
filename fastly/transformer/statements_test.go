package transformer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
)

func TestTransformDeclareStatement(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: `declare local var.Foo STRING;`, expect: "var local__Foo string"},
		{input: `declare local var.Foo INTEGER;`, expect: "var local__Foo int64"},
		{input: `declare local var.Foo FLOAT;`, expect: "var local__Foo float64"},
		{input: `declare local var.Foo BOOL;`, expect: "var local__Foo bool"},
		{input: `declare local var.Foo IP;`, expect: "var local__Foo net.IP"},
		{input: `declare local var.Foo RTIME;`, expect: "var local__Foo time.Duration"},
		{input: `declare local var.Foo TIME;`, expect: "var local__Foo time.Time"},
		{input: `declare local var.Foo BACKEND;`, expect: "var local__Foo *vintage.Backend"},
		{input: `declare local var.Foo ACL;`, expect: "var local__Foo *vintage.Acl"},
	}

	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := New().(*transformer)
		code, err := tr.transformDeclareStatement(stmt[0].(*ast.DeclareStatement))
		if err != nil {
			t.Errorf("Declaration transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("declaration mismatch, diff=%s", diff)
		}
	}
}
