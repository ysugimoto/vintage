package core

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/vintage/transformer/value"
)

func TestTransformDeclareStatement(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: `declare local var.Foo STRING;`, expect: "var l_Foo string"},
		{input: `declare local var.Foo INTEGER;`, expect: "var l_Foo int64"},
		{input: `declare local var.Foo FLOAT;`, expect: "var l_Foo float64"},
		{input: `declare local var.Foo BOOL;`, expect: "var l_Foo bool"},
		{input: `declare local var.Foo IP;`, expect: "var l_Foo net.IP"},
		{input: `declare local var.Foo RTIME;`, expect: "var l_Foo time.Duration"},
		{input: `declare local var.Foo TIME;`, expect: "var l_Foo time.Time"},
		{input: `declare local var.Foo BACKEND;`, expect: "var l_Foo *vintage.Backend"},
		{input: `declare local var.Foo ACL;`, expect: "var l_Foo *vintage.Acl"},
	}

	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
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

func TestTransformReturnStatement(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "return(pass);", expect: "return vintage.PASS, nil"},
		{input: "return;", expect: "return vintage.NONE, nil"},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code := tr.transformReturnStatement(stmt[0].(*ast.ReturnStatement))

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformSetStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `set req.http.Foo = "bar";`, expect: `ctx.RequestHeader.Set("Foo", "bar")`},
		{input: `set var.Foo = "bar";`, expect: `l_Foo = "bar"`},
		{input: `set client.socket.congestion_algorithm = "cubic";`, expect: `ctx.ClientSocketCongestionAlgorithm = "cubic"`},
		{input: `set var.Hoge = "huga";`, isError: true},
		{input: `set client.some.undefined_variable = 100;`, isError: true},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		tr.vars["var.Foo"] = value.NewValue(value.STRING, "l_Foo")
		code, err := tr.transformSetStatement(stmt[0].(*ast.SetStatement))
		if tt.isError {
			if err == nil {
				t.Errorf("Expects error, but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Set statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformAddStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `add resp.http.Set-Cookie = "key=value";`, expect: `ctx.ResponseHeader.Add("Set-Cookie", "key=value")`},
		{input: `add req.http.Foo = "value";`, expect: `ctx.RequestHeader.Add("Foo", "value")`},
		{input: `add bereq.http.Foo = "value";`, expect: `ctx.BackendRequestHeader.Add("Foo", "value")`},
		{input: `add beresp.http.Foo = "value";`, expect: `ctx.BackendResponseHeader.Add("Foo", "value")`},
		{input: `add var.Foo = "key=value";`, isError: true},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformAddStatement(stmt[0].(*ast.AddStatement))
		if err != nil {
			t.Errorf("Add statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformUnsetStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `unset req.http.Foo;`, expect: `ctx.RequestHeader.Unset("Foo")`},
		{input: `unset var.Foo;`, isError: true},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformUnsetStatement(stmt[0].(*ast.UnsetStatement))
		if err != nil {
			t.Errorf("Unset statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformRemoveStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `remove req.http.Foo;`, expect: `ctx.RequestHeader.Unset("Foo")`},
		{input: `remove var.Foo;`, isError: true},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformRemoveStatement(stmt[0].(*ast.RemoveStatement))
		if err != nil {
			t.Errorf("Remove statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformLogStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `log {" syslog fake_service_id fastly-logs ::"} {" foobar "};`, expect: `if err := logging__fastly_logs.Write("" + " foobar "); err != nil {` + lineFeed + "return vintage.NONE, err\n}"},
		{input: `log {" syslog fake_service_id undefined-logs ::"} {" foobar "};`, isError: true},
		{input: `log {" foobar "};`, expect: ""},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		tr.loggingEndpoints["fastly-logs"] = "logging__fastly_logs"
		code, err := tr.transformLogStatement(stmt[0].(*ast.LogStatement))
		if err != nil {
			t.Errorf("Log statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformSyntheticStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `synthetic {" foobar "};`, expect: `ctx.BackendResponse.Body = io.NopCloser(strings.NewReader(" foobar "))`},
		{input: `synthetic 0;`, expect: `ctx.BackendResponse.Body = io.NopCloser(strings.NewReader(vintage.ToString(0)))`},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformSyntheticStatement(stmt[0].(*ast.SyntheticStatement))
		if err != nil {
			t.Errorf("Synthetic statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformSyntheticBase64Statement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{input: `synthetic.base64 {" foobar "};`, expect: `ctx.BackendResponse.Body = io.NopCloser(base64.NewDecoder(base64.StdEncoding, strings.NewReader(" foobar ")))`},
		{input: `synthetic.base64 0;`, expect: `ctx.BackendResponse.Body = io.NopCloser(base64.NewDecoder(base64.StdEncoding, strings.NewReader(vintage.ToString(0))))`},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformSyntheticBase64Statement(stmt[0].(*ast.SyntheticBase64Statement))
		if err != nil {
			t.Errorf("SyntheticBase64 statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformCallStatement(t *testing.T) {
	value.UseFixedTemporalValue()
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input: `call custom_subroutine;`,
			expect: strings.Join([]string{
				"if v_fixed, err := custom_subroutine(ctx); err != nil {",
				"return vintage.NONE, err",
				"} else if v_fixed != vintage.NONE {",
				"return v_fixed, nil",
				"}",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code := tr.transformCallStatement(stmt[0].(*ast.CallStatement))

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformRestartStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input:  `restart;`,
			expect: "return vintage.RESTART, nil",
		},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformRestartStatement(stmt[0].(*ast.RestartStatement))
		if err != nil {
			t.Errorf("Restart statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformErrorStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input: `error 400;`,
			expect: strings.Join([]string{
				"ctx.ObjectStatus = 400",
				"return vintage.ERROR, nil",
			}, "\n"),
		},
		{
			input: `error 400 "bad_request";`,
			expect: strings.Join([]string{
				"ctx.ObjectStatus = 400",
				`ctx.ObjectResponse = "bad_request"`,
				"return vintage.ERROR, nil",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformErrorStatement(stmt[0].(*ast.ErrorStatement))
		if err != nil {
			t.Errorf("Restart statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformFunctionCallStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input:   `table.lookup(foo, "bar");`,
			isError: true,
		},
		{
			input: `early_hints("link: </hinted.js>; rel=preload", "link: </hinted.css>; rel=preload");`,
			expect: strings.Join([]string{
				`if _, err := function.Early_hints(ctx.Runtime, "link: </hinted.js>; rel=preload", "link: </hinted.css>; rel=preload"); err != nil {`,
				"return vintage.NONE, err",
				"}",
			}, "\n"),
		},
		{
			input: `early_hints("link: </hinted.js>; rel=preload", "link: </hinted.css>; rel=preload");`,
			expect: strings.Join([]string{
				`if _, err := function.Early_hints(ctx.Runtime, "link: </hinted.js>; rel=preload", "link: </hinted.css>; rel=preload"); err != nil {`,
				"return vintage.NONE, err",
				"}",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformFunctionCallStatement(stmt[0].(*ast.FunctionCallStatement))
		if err != nil {
			t.Errorf("FunctionCall statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}

func TestTransformIfStatement(t *testing.T) {
	tests := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input: `
if (!req.http.Foo) {
	set req.http.Foo = "bar";
} else if (!req.http.Hoge) {
	set req.http.Hoge = "huga";
} else {
	restart;
}
`,
			expect: strings.Join([]string{
				`if !vintage.ToBool(ctx.RequestHeader.Get("Foo")) {`,
				`ctx.RequestHeader.Set("Foo", "bar")`,
				`} else if !vintage.ToBool(ctx.RequestHeader.Get("Hoge")) {`,
				`ctx.RequestHeader.Set("Hoge", "huga")`,
				"} else {",
				"return vintage.RESTART, nil",
				"}",
			}, "\n"),
		},
	}
	for _, tt := range tests {
		stmt, err := parser.New(lexer.NewFromString(tt.input)).ParseSnippetVCL()
		if tt.isError {
			if err != nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected parse error: %s", err)
			continue
		}
		tr := NewCoreTransfromer()
		code, err := tr.transformIfStatement(stmt[0].(*ast.IfStatement))
		if err != nil {
			t.Errorf("If statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(string(code), tt.expect); diff != "" {
			t.Errorf("code generation mismatch, diff=%s", diff)
		}
	}
}
