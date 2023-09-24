package transformer

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage"
)

func (tf *transformer) transformExpression(
	expect vintage.VCLType,
	expr ast.Expression,
) (*expressionValue, error) {

	var value *expressionValue
	var err error

	switch t := expr.(type) {
	case *ast.Ident:
		value, err = tf.transformIdentValue(t)
	case *ast.IP:
		value = newExpressionValue(vintage.IP, []byte(`"`+net.ParseIP(t.Value).String()+`"`))
	case *ast.Boolean:
		value = newExpressionValue(vintage.BOOL, []byte(fmt.Sprintf("%t", t.Value)))
	case *ast.Integer:
		value = newExpressionValue(vintage.INTEGER, []byte(fmt.Sprint(t.Value)))
	case *ast.Float:
		value = newExpressionValue(vintage.FLOAT, []byte(fmt.Sprint(t.Value)))
	case *ast.String:
		value = newExpressionValue(vintage.STRING, []byte(`"`+t.Value+`"`))
	case *ast.RTime:
		var val time.Duration
		switch {
		case strings.HasSuffix(t.Value, "d"):
			num := strings.TrimSuffix(t.Value, "d")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = newExpressionValue(vintage.RTIME, []byte(fmt.Sprintf("(%d * time.Hour)", val*24)))
		case strings.HasSuffix(t.Value, "y"):
			num := strings.TrimSuffix(t.Value, "y")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = newExpressionValue(vintage.RTIME, []byte(fmt.Sprintf("(%d * time.Hour)", val*24*365)))
		default:
			val, err = time.ParseDuration(t.Value)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = newExpressionValue(vintage.RTIME, []byte(fmt.Sprintf("(%d * time.Second)", val)))
		}

	// Combinated expressions
	case *ast.PrefixExpression:
		value, err = tf.transformPrefixExpression(t)
	case *ast.GroupedExpression:
		value, err = tf.transformGroupedExpression(t)
	case *ast.IfExpression:
		value, err = tf.transformIfExpression(expect, t)
	case *ast.InfixExpression:
		value, err = tf.transformInfixExpression(t)
	case *ast.FunctionCallExpression:
		value, err = tf.transformFunctionCallExpression(expect, t)
	}

	if err != nil {
		return nil, TransformError(&expr.GetMeta().Token, "Undefined expression found")
	}
	return value.Conversion(expect), nil
}

func (tf *transformer) transformIdentValue(ident *ast.Ident) (*expressionValue, error) {
	name := ident.Value
	if v, ok := tf.backends[name]; ok {
		return v, nil
	} else if v, ok := tf.acls[name]; ok {
		return v, nil
	} else if v, ok := tf.tables[name]; ok {
		return v, nil
	} else if strings.HasPrefix(name, "var.") {
		if v, ok := tf.vars[name]; !ok {
			return nil, TransformError(&ident.GetMeta().Token, "local variable %s is undefined", name)
			// } else if v, ok := predefinedVariables(name); ok {
			// 	identValue = v()
		} else {
			return v, nil
		}
	}
	return nil, TransformError(&ident.GetMeta().Token, "Undefined indent: %s", name)
}

func (tf *transformer) transformPrefixExpression(expr *ast.PrefixExpression) (*expressionValue, error) {
	switch expr.Operator {
	case "!":
		right, err := tf.transformExpression(vintage.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = append([]byte("!"), right.Code...)
		return right, nil
	case "-":
		right, err := tf.transformExpression(vintage.INTEGER, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = append([]byte("-"), right.Code...)
		return right, nil
	case "+":
		right, err := tf.transformExpression(vintage.STRING, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return right, nil
	}
	return nil, TransformError(&expr.GetMeta().Token, "Unexpected prefix operator found: %s", expr.Operator)
}

func (tf *transformer) transformGroupedExpression(expr *ast.GroupedExpression) (*expressionValue, error) {
	right, err := tf.transformExpression(vintage.NULL, expr.Right)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	right.Code = append([]byte("("), right.Code...)
	right.Code = append(right.Code, ')')
	return right, nil
}

func (tf *transformer) transformIfExpression(expect vintage.VCLType, expr *ast.IfExpression) (*expressionValue, error) {
	condition, err := tf.transformExpression(vintage.BOOL, expr.Condition)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	consequence, err := tf.transformExpression(expect, expr.Consequence)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	alternative, err := tf.transformExpression(expect, expr.Alternative)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("vintage.TernaryOperator[%s](", vintage.GoTypeString(expect)))
	buf.Write(condition.Code)
	buf.WriteString(", ")
	buf.Write(consequence.Code)
	buf.WriteString(", ")
	buf.Write(alternative.Code)
	buf.WriteString(")")

	return newExpressionValue(expect, buf.Bytes()), nil
}
func (tf *transformer) transformInfixExpression(expr *ast.InfixExpression) (*expressionValue, error) {
	var buf bytes.Buffer

	switch expr.Operator {
	case "==", "!=", ">", "<", ">=", "<=":
		left, err := tf.transformExpression(vintage.NULL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(vintage.NULL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(left.Code)
		buf.WriteString(" " + expr.Operator + " ")
		buf.Write(right.Code)
		return newExpressionValue(vintage.BOOL, buf.Bytes()), nil
	// "~" or "!~" need regular expression matching
	case "~", "!~":
		left, err := tf.transformExpression(vintage.NULL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(vintage.NULL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// If right expression is ACL, do CIDR matching
		if right.Type == vintage.ACL {
			if expr.Operator == "!~" {
				buf.WriteString("!")
			}
			buf.Write(right.Code)
			buf.WriteString(".Match(")
			buf.Write(left.Code)
			buf.WriteString(")")
			return newExpressionValue(vintage.BOOL, buf.Bytes()), nil
		}

		// Otherwise, string matching
		buf.WriteString("vintage.RegexpMatchOperator(")
		buf.Write(right.Code)
		buf.WriteString(", ")
		buf.Write(left.Code)
		return newExpressionValue(vintage.BOOL, buf.Bytes()), nil

	case "||", "&&":
		left, err := tf.transformExpression(vintage.BOOL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(vintage.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(left.Code)
		buf.WriteString(" " + expr.Operator + " ")
		buf.Write(right.Code)
		return newExpressionValue(vintage.BOOL, buf.Bytes()), nil
	// "+" means string concatenation
	case "+":
		left, err := tf.transformExpression(vintage.STRING, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(vintage.STRING, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(left.Code)
		buf.WriteString(" + ")
		buf.Write(right.Code)
		return newExpressionValue(vintage.STRING, buf.Bytes()), nil
	}
	return nil, TransformError(&expr.GetMeta().Token, "Unexpected infix operator: %s", expr.Operator)
}

func (tf *transformer) transformFunctionCallExpression(expect vintage.VCLType, expr *ast.FunctionCallExpression) (*expressionValue, error) {
	tf.packages.Add("github.com/ysugimoto/vintage/builtin", "")

	// TODO: add function arguments expression
	call := fmt.Sprintf(
		"vintage.%s()",
		ucFirst(strings.ReplaceAll(expr.Function.Value, ".", "_")),
	)
	return newExpressionValue(expect, []byte(call)), nil
}
