package core

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/falco/lexer"
	"github.com/ysugimoto/falco/parser"
	"github.com/ysugimoto/vintage/transformer/value"
)

func TestTransformExpression(t *testing.T) {
	value.UseFixedTemporalValue()
	tests := []struct {
		input   string
		T       value.VCLType
		expect  *value.Value
		isError bool
	}{
		{
			input:  "set req.http.Foo = true;",
			expect: value.NewValue(value.STRING, "vintage.ToString(true)"),
		},
		{
			input:  "set req.http.Foo = false;",
			expect: value.NewValue(value.STRING, "vintage.ToString(false)"),
		},
		{
			input:  "set req.http.Foo = 10;",
			expect: value.NewValue(value.STRING, "vintage.ToString(10)"),
		},
		{
			input:  "set req.http.Foo = 10.01;",
			expect: value.NewValue(value.STRING, "vintage.ToString(10.01)"),
		},
		{
			input:  `set req.http.Foo = "bar";`,
			expect: value.NewValue(value.STRING, `"bar"`),
		},
		{
			input:  `set req.http.Foo = 10d;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(time.Duration(240 * time.Hour))`),
		},
		{
			input:  `set req.http.Foo = 1y;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(time.Duration(8760 * time.Hour))`),
		},
		{
			input:  `set req.http.Foo = 10m;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(time.Duration(600 * time.Second))`),
		},
		{
			input:  `set req.http.Foo = !req.http.Bar;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(!vintage.ToBool(ctx.RequestHeader.Get("Bar")))`),
		},
		{
			input:  `set req.http.Foo = -10;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(-10)`),
		},
		{
			input:  `set req.http.Foo = -10.01;`,
			expect: value.NewValue(value.STRING, `vintage.ToString(-10.01)`),
		},
		{
			input:  `set req.http.Foo = (10);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10))`),
		},
		{
			input: `set req.http.Foo = if(req.http.Bar, "hoge", "huga");`,
			expect: value.NewValue(value.STRING, "v__fixed", value.Prepare(
				`v__fixed := "huga"`,
				`if vintage.ToBool(ctx.RequestHeader.Get("Bar")) {`,
				`v__fixed = "hoge"`,
				`}`,
			)),
		},
		{
			input:  `set req.http.Foo = (10 > 0);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10 > 0))`),
		},
		{
			input:  `set req.http.Foo = (10 == 10);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10 == 10))`),
		},
		{
			input:  `set req.http.Foo = (10 != 10);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10 != 10))`),
		},
		{
			input:  `set req.http.Foo = (0 < 10);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((0 < 10))`),
		},
		{
			input:  `set req.http.Foo = (0 <= 10);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((0 <= 10))`),
		},
		{
			input:  `set req.http.Foo = (10 >= 0);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10 >= 0))`),
		},
		{
			input:  `set req.http.Foo = (10 >= 0);`,
			expect: value.NewValue(value.STRING, `vintage.ToString((10 >= 0))`),
		},
		{
			input: `set req.http.Foo = ("foobar" ~ "bar");`,
			expect: value.NewValue(value.STRING, "vintage.ToString((v__fixed))", value.Prepare(
				`v__fixed, v__fixed_group, err := vintage.RegexpMatch(`+"`bar`"+`, "foobar")`,
				value.ErrorCheck,
				`_ = v__fixed_group // implicitly avoid compilation error`,
			)),
		},
		{
			input: `set req.http.Foo = (req.http.Bar && "bar");`,
			expect: value.NewValue(
				value.STRING,
				`vintage.ToString((vintage.ToBool(ctx.RequestHeader.Get("Bar")) && vintage.ToBool("bar")))`,
			),
		},
		{
			input:  `set req.http.Foo = "foo" "bar" "baz";`,
			expect: value.NewValue(value.STRING, `"foo" + "bar" + "baz"`),
		},
		{
			input:  `set req.http.Foo = {" foo bar baz "};`,
			expect: value.NewValue(value.STRING, `" foo bar baz "`),
		},
		{
			input: `set req.http.Foo = substr(req.http.Bar, 1);`,
			expect: value.NewValue(value.STRING, "v__fixed", value.Prepare(
				`v__fixed, err := function.Substr(ctx.Runtime, ctx.RequestHeader.Get("Bar"), 1)`,
				value.ErrorCheck,
			)),
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
		tr.acls["example"] = value.NewValue(value.ACL, "acl__example")
		s := stmt[0].(*ast.SetStatement)
		v, err := tr.transformExpression(value.STRING, s.Value)
		if err != nil {
			t.Errorf("expression transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(v, tt.expect); diff != "" {
			t.Errorf("value mismatch, diff=%s", diff)
		}
	}
}

func TestTransformIdentValue(t *testing.T) {
	tests := []struct {
		input   string
		expect  *value.Value
		isError bool
	}{
		{
			input:  `test_backend`,
			expect: value.NewValue(value.BACKEND, "B__test_backend"),
		},
		{
			input:  `test_acl`,
			expect: value.NewValue(value.ACL, "A__test_acl"),
		},
		{
			input:  `test_table`,
			expect: value.NewValue(value.IDENT, "T__test_table"),
		},
		{
			input:  `var.FOO`,
			expect: value.NewValue(value.STRING, "local__FOO"),
		},
		{
			input:   `re.group.1`,
			isError: true,
		},
		{
			input:  `aes256`,
			expect: value.NewValue(value.IDENT, "aes256"),
		},
		{
			input:  `req.http.Foo`,
			expect: value.NewValue(value.STRING, `ctx.RequestHeader.Get("Foo")`),
		},
		{
			input:   `req.undefined`,
			isError: true,
		},
	}

	for _, tt := range tests {
		tr := NewCoreTransfromer()
		tr.backends["test_backend"] = value.NewValue(value.BACKEND, "B__test_backend")
		tr.acls["test_acl"] = value.NewValue(value.ACL, "A__test_acl")
		tr.tables["test_table"] = value.NewValue(value.IDENT, "T__test_table")
		tr.vars["var.FOO"] = value.NewValue(value.STRING, "local__FOO")
		ident := &ast.Ident{
			Value: tt.input,
			Meta:  &ast.Meta{},
		}
		v, err := tr.transformExpression(value.NULL, ident)
		if tt.isError {
			if err == nil {
				t.Errorf("Expects error but got nil")
				continue
			}
			continue
		}
		if err != nil {
			t.Errorf("If statement transforming error: %s", err)
			continue
		}

		if diff := cmp.Diff(v, tt.expect); diff != "" {
			t.Errorf("value mismatch, diff=%s", diff)
		}
	}
}
