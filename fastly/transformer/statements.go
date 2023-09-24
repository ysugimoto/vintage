package transformer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage"
)

func (t *transformer) transformBlockStatement(statements []ast.Statement) ([]byte, error) {
	var buf bytes.Buffer

	for _, stmt := range statements {
		var code []byte
		var err error

		switch s := stmt.(type) {
		case *ast.DeclareStatement:
			code, err = t.transformDeclareStatement(s)
		case *ast.ReturnStatement:
			code, err = t.transformReturnStatement(s)
		case *ast.SetStatement:
			code, err = t.transformSetStatement(s)
		}
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(code)
		buf.WriteString(LF)
	}

	return buf.Bytes(), nil
}

func (t *transformer) transformDeclareStatement(stmt *ast.DeclareStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := strings.TrimPrefix(stmt.Name.Value, "var.")
	switch vintage.VCLType(stmt.ValueType.Value) {
	case vintage.STRING:
		buf.WriteString(fmt.Sprintf("var local__%s string", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.STRING, []byte("local__ "+name))
	case vintage.INTEGER:
		buf.WriteString(fmt.Sprintf("var local__%s int64", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.INTEGER, []byte("local__"+name))
	case vintage.BOOL:
		buf.WriteString(fmt.Sprintf("var local__%s bool", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.BOOL, []byte("local__"+name))
	case vintage.FLOAT:
		buf.WriteString(fmt.Sprintf("var local__%s float64", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.FLOAT, []byte("local__"+name))
	case vintage.BACKEND:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Backend", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.BACKEND, []byte("local__"+name))
	case vintage.IP:
		buf.WriteString(fmt.Sprintf("var local__%s net.IP", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.IP, []byte("local__"+name))
	case vintage.RTIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Duration", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.RTIME, []byte("local__"+name))
	case vintage.TIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Time", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.TIME, []byte("local__"+name))
	case vintage.ACL:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Acl", name))
		t.vars[stmt.Name.Value] = newExpressionValue(vintage.ACL, []byte("local__"+name))
	default:
		return nil, errors.WithStack(
			fmt.Errorf("Unexpected variable type declared: %s", stmt.ValueType.Value),
		)
	}
	return buf.Bytes(), nil
}

func (t *transformer) transformReturnStatement(stmt *ast.ReturnStatement) ([]byte, error) {
	state := "vintage.NONE"
	if stmt.ReturnExpression != nil {
		state = "vintage." + strings.ToUpper(strings.Trim(toString(*stmt.ReturnExpression), `"`))
	}
	return []byte(fmt.Sprintf("return %s, nil", state)), nil
}

func (t *transformer) transformSetStatement(stmt *ast.SetStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value
	switch {
	case strings.HasPrefix(name, "var."):
		v, ok := t.vars[name]
		if !ok {
			return nil, TransformError(&stmt.Token, "variable %s undefined", name)
		}
		val, err := t.transformExpression(v.Type, stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(v.Code)
		buf.WriteString(" = ")
		buf.Write(val.Code)
	case strings.HasPrefix(name, "req.http."),
		strings.HasPrefix(name, "bereq.http."):
		name := strings.TrimPrefix(
			strings.TrimPrefix(name, "bereq.http."),
			"req.http.",
		)
		val, err := t.transformExpression(vintage.STRING, stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(`ctx.RequestHeader.Set("` + name + `", `)
		buf.Write(val.Code)
		buf.WriteString(")")
	case strings.HasPrefix(name, "beresp.http."),
		strings.HasPrefix(name, "resp.http."),
		strings.HasPrefix(name, "obj.http."):
		name := strings.TrimPrefix(
			strings.TrimPrefix(
				strings.TrimPrefix(name, "obj.http."),
				"resp.http.",
			),
			"beresp.http.",
		)
		val, err := t.transformExpression(vintage.STRING, stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(`ctx.ResponseHeader.Set("` + name + `", `)
		buf.Write(val.Code)
		buf.WriteString(")")
	default:
		break
		// format, expectType, err := variables.GetType(name)
		// if err != nil {
		// 	return nil, errors.WithStack(err)
		// }
		// val, err := t.transformExpression(expectType, stmt.Value)
		// if err != nil {
		// 	return nil, errors.WithStack(err)
		// }
		// buf.WriteString(fmt.Sprintf(format+LF, val))
	}

	return buf.Bytes(), nil
}
