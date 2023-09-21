package transformer

import (
	"go/format"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
)

func TestTransformAcl(t *testing.T) {
	input := `
acl example {
  "192.168.0.1"/32;
  !"192.168.0.2"/32;
}`
	vcl, err := parser.New(lexer.NewFromString(input)).ParseVCL()
	if err != nil {
		t.Errorf("Unexpected acl parsing error: %s", err)
		return
	}
	if len(vcl.Statements) != 1 {
		t.Errorf("Unexpected statement count, expect=1, got=%d", len(vcl.Statements))
		return
	}
	acl := vcl.Statements[0]
	tr := New().(*transformer)
	code, err := tr.transformAcl(acl.(*ast.AclDeclaration))
	if err != nil {
		t.Errorf("Failed to transform ACL: %s", err)
		return
	}
	code, err = format.Source(code)
	if err != nil {
		t.Errorf("Failed to format code: %s", err)
		return
	}

	expect := `
var acl__example = vintage.NewAcl(
	vintage.AclEntry("192.168.0.1/32", false),
	vintage.AclEntry("192.168.0.2/32", true),
)
`
	if diff := cmp.Diff("\n"+string(code), expect); diff != "" {
		t.Errorf("Acl transform result mismatch, diff=%s", diff)
	}
}

func TestTransformBackend(t *testing.T) {
	input := `
backend example {
  .host = "example.com";
  .probe = {
    .request = "GET / HTTP/1.1";
  }
}`
	vcl, err := parser.New(lexer.NewFromString(input)).ParseVCL()
	if err != nil {
		t.Errorf("Unexpected acl parsing error: %s", err)
		return
	}
	if len(vcl.Statements) != 1 {
		t.Errorf("Unexpected statement count, expect=1, got=%d", len(vcl.Statements))
		return
	}
	backend := vcl.Statements[0]
	tr := New().(*transformer)
	code, err := tr.transformBackend(backend.(*ast.BackendDeclaration))
	if err != nil {
		t.Errorf("Failed to transform Backend: %s", err)
		return
	}
	code, err = format.Source(code)
	if err != nil {
		t.Errorf("Failed to format code: %s", err)
		return
	}
	expect := `
var backend__example = vintage.NewBackend("example")
`
	if diff := cmp.Diff("\n"+string(code), expect); diff != "" {
		t.Errorf("Backend transform result mismatch, diff=%s", diff)
	}
}

func TestTransformDirector(t *testing.T) {
	input := `
director example client {
  .quorum = "20%";
  { .backend = foo; .weight = 1; }
}`
	vcl, err := parser.New(lexer.NewFromString(input)).ParseVCL()
	if err != nil {
		t.Errorf("Unexpected acl parsing error: %s", err)
		return
	}
	if len(vcl.Statements) != 1 {
		t.Errorf("Unexpected statement count, expect=1, got=%d", len(vcl.Statements))
		return
	}
	d := vcl.Statements[0]
	tr := New().(*transformer)
	code, err := tr.transformDirector(d.(*ast.DirectorDeclaration))
	if err != nil {
		t.Errorf("Failed to transform Director: %s", err)
		return
	}
	code, err = format.Source(code)
	if err != nil {
		t.Errorf("Failed to format code: %s", err)
		return
	}
	expect := `
var director__example = vintage.NewDirector("example", "client",
	vintage.DirectorProperty("quorum", "20%"),
	vintage.DirectorBackend(
		vintage.DirectorProperty("backend", "foo"),
		vintage.DirectorProperty("weight", 1),
	),
)
`
	if diff := cmp.Diff("\n"+string(code), expect); diff != "" {
		t.Errorf("Director transform result mismatch, diff=%s", diff)
	}
}

func TestTransformTable(t *testing.T) {
	input := `
table example STRING {
  "foo": "bar",
}`
	vcl, err := parser.New(lexer.NewFromString(input)).ParseVCL()
	if err != nil {
		t.Errorf("Unexpected acl parsing error: %s", err)
		return
	}
	if len(vcl.Statements) != 1 {
		t.Errorf("Unexpected statement count, expect=1, got=%d", len(vcl.Statements))
		return
	}
	table := vcl.Statements[0]
	tr := New().(*transformer)
	code, err := tr.transformTable(table.(*ast.TableDeclaration))
	if err != nil {
		t.Errorf("Failed to transform Table: %s", err)
		return
	}
	code, err = format.Source(code)
	if err != nil {
		t.Errorf("Failed to format code: %s", err)
		return
	}
	expect := `
var table__example = vintage.NewTable("example", "STRING",
	vintage.TableItem("foo", "bar"),
)
`
	if diff := cmp.Diff("\n"+string(code), expect); diff != "" {
		t.Errorf("Table transform result mismatch, diff=%s", diff)
	}
}

func TestTransformSubroutine(t *testing.T) {
	input := `
sub vcl_recv {
  set req.http.Foo = "bar";
}`
	expect := "func vcl_recv(ctx *fastly.Context) (vintage.State, error) {\n}\n"
	vcl, err := parser.New(lexer.NewFromString(input)).ParseVCL()
	if err != nil {
		t.Errorf("Unexpected acl parsing error: %s", err)
		return
	}
	if len(vcl.Statements) != 1 {
		t.Errorf("Unexpected statement count, expect=1, got=%d", len(vcl.Statements))
		return
	}
	sub := vcl.Statements[0]
	tr := New().(*transformer)
	code, err := tr.transformSubroutine(sub.(*ast.SubroutineDeclaration))
	if err != nil {
		t.Errorf("Unexpected subroutine transforming error: %s", err)
		return
	}
	code, err = format.Source(code)
	if err != nil {
		t.Errorf("Failed to format code: %s", err)
		return
	}
	if diff := cmp.Diff(string(code), expect); diff != "" {
		t.Errorf("Subroutine code mismatch, diff=%s", diff)
		return
	}
}
