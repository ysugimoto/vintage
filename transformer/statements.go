package transformer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage"
)

func (tf *CoreTransformer) transformBlockStatement(statements []ast.Statement) ([]byte, bool, error) {
	var buf bytes.Buffer
	var returnExists bool

	for _, stmt := range statements {
		var code []byte
		var err error

		switch s := stmt.(type) {
		case *ast.DeclareStatement:
			code, err = tf.transformDeclareStatement(s)
		case *ast.ReturnStatement:
			code, err = tf.transformReturnStatement(s)
			returnExists = true
		case *ast.SetStatement:
			code, err = tf.transformSetStatement(s)
		}
		if err != nil {
			return nil, false, errors.WithStack(err)
		}
		if len(code) > 0 {
			buf.Write(code)
			buf.WriteString(lineFeed)
		}
	}

	return buf.Bytes(), returnExists, nil
}

func (tf *CoreTransformer) transformDeclareStatement(stmt *ast.DeclareStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := strings.TrimPrefix(stmt.Name.Value, "var.")
	switch vintage.VCLType(stmt.ValueType.Value) {
	case vintage.STRING:
		buf.WriteString(fmt.Sprintf("var local__%s string", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.STRING, "local__ "+name)
	case vintage.INTEGER:
		buf.WriteString(fmt.Sprintf("var local__%s int64", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.INTEGER, "local__"+name)
	case vintage.BOOL:
		buf.WriteString(fmt.Sprintf("var local__%s bool", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.BOOL, "local__"+name)
	case vintage.FLOAT:
		buf.WriteString(fmt.Sprintf("var local__%s float64", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.FLOAT, "local__"+name)
	case vintage.BACKEND:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Backend", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.BACKEND, "local__"+name)
	case vintage.IP:
		buf.WriteString(fmt.Sprintf("var local__%s net.IP", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.IP, "local__"+name)
	case vintage.RTIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Duration", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.RTIME, "local__"+name)
	case vintage.TIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Time", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.TIME, "local__"+name)
	case vintage.ACL:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Acl", name))
		tf.vars[stmt.Name.Value] = NewExpressionValue(vintage.ACL, "local__"+name)
	default:
		return nil, errors.WithStack(
			fmt.Errorf("Unexpected variable type declared: %s", stmt.ValueType.Value),
		)
	}
	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformReturnStatement(stmt *ast.ReturnStatement) ([]byte, error) {
	state := "vintage.NONE"
	if stmt.ReturnExpression != nil {
		state = "vintage." + strings.ToUpper(strings.Trim(toString(*stmt.ReturnExpression), `"`))
	}
	return []byte(fmt.Sprintf("return %s, nil", state)), nil
}

func (tf *CoreTransformer) transformSetStatement(stmt *ast.SetStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value
	switch {
	case strings.HasPrefix(name, "var."):
		v, ok := tf.vars[name]
		if !ok {
			return nil, TransformError(&stmt.Token, "variable %s undefined", name)
		}
		val, err := tf.transformExpression(v.Type, stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(val.Prepare + v.Code + " = " + val.Code + val.Comment)
		// case strings.HasPrefix(name, "req.http."):
		// 	strings.HasPrefix(name, "bereq.http."):
		// 	name := strings.TrimPrefix(
		// 		strings.TrimPrefix(name, "bereq.http."),
		// 		"req.http.",
		// 	)
		// 	val, err := tf.transformExpression(vintage.STRING, stmt.Value)
		// 	if err != nil {
		// 		return nil, errors.WithStack(err)
		// 	}
		// 	buf.WriteString(val.Prepare + fmt.Sprintf(`ctx.RequestHeader.Set("%s", %s)%s`, name, val.Code, val.Comment))
		// case strings.HasPrefix(name, "beresp.http."),
		// 	strings.HasPrefix(name, "resp.http."),
		// 	strings.HasPrefix(name, "obj.http."):
		// 	name := strings.TrimPrefix(
		// 		strings.TrimPrefix(
		// 			strings.TrimPrefix(name, "obj.http."),
		// 			"resp.http.",
		// 		),
		// 		"beresp.http.",
		// 	)
		// 	val, err := tf.transformExpression(vintage.STRING, stmt.Value)
		// 	if err != nil {
		// 		return nil, errors.WithStack(err)
		// 	}
		// 	buf.WriteString(val.Prepare + fmt.Sprintf(`ctx.ResponseHeader.Set("%s", %s)`, name, val.Code))
		// default:
		// 	break
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
