package transformer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
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
	t.vars[name] = stmt.ValueType.Value
	switch stmt.ValueType.Value {
	case "STRING":
		buf.WriteString(fmt.Sprintf("var local__%s string", name))
	case "INTEGER":
		buf.WriteString(fmt.Sprintf("var local__%s int64", name))
	case "BOOL":
		buf.WriteString(fmt.Sprintf("var local__%s bool", name))
	case "FLOAT":
		buf.WriteString(fmt.Sprintf("var local__%s float64", name))
	case "BACKEND":
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Backend", name))
	case "IP":
		buf.WriteString(fmt.Sprintf("var local__%s net.IP", name))
	case "RTIME":
		buf.WriteString(fmt.Sprintf("var local__%s time.Duration", name))
	case "TIME":
		buf.WriteString(fmt.Sprintf("var local__%s time.Time", name))
	case "ACL":
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Acl", name))
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
		state = "vintage." + strings.ToUpper(toString(*stmt.ReturnExpression))
	}
	return []byte(fmt.Sprintf("return %s, nil", state)), nil
}

func (t *transformer) transformSetStatement(stmt *ast.SetStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value
	switch {
	case strings.HasPrefix(name, "var."):
		val, err := t.transformExpression(t.vars[name], stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(
			fmt.Sprintf(`local__%s = %s`, strings.TrimPrefix(name, "var."), val),
		)
	case strings.HasPrefix(name, "req.http."),
		strings.HasPrefix(name, "bereq.http."):
		name := strings.TrimPrefix(
			strings.TrimPrefix(name, "bereq.http."),
			"req.http.",
		)
		val, err := t.transformExpression("STRING", stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(
			fmt.Sprintf(`ctx.RequestHeader.Set("%s", "%s")`, name, val),
		)
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
		val, err := t.transformExpression("STRING", stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(
			fmt.Sprintf(`ctx.ResponseHeader.Set("%s", "%s")`, name, val),
		)
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
