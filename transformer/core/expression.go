package core

import (
	"bytes"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage/transformer/value"
)

func (tf *CoreTransformer) transformExpression(
	expect value.VCLType,
	expr ast.Expression,
) (*value.Value, error) {

	var v *value.Value
	var err error

	switch t := expr.(type) {
	case *ast.Ident:
		v, err = tf.transformIdentValue(t)
	case *ast.IP:
		v = value.NewValue(value.IP, `"`+net.ParseIP(t.Value).String()+`"`)
	case *ast.Boolean:
		v = value.NewValue(value.BOOL, fmt.Sprintf("%t", t.Value))
	case *ast.Integer:
		v = value.NewValue(value.INTEGER, fmt.Sprint(t.Value))
	case *ast.Float:
		v = value.NewValue(value.FLOAT, fmt.Sprint(t.Value))
	case *ast.String:
		v = value.NewValue(value.STRING, `"`+t.Value+`"`)
	case *ast.RTime:
		var val time.Duration
		switch {
		case strings.HasSuffix(t.Value, "d"):
			num := strings.TrimSuffix(t.Value, "d")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			v = value.NewValue(value.RTIME, fmt.Sprintf("time.Duration(%d * time.Hour)", val*24))
		case strings.HasSuffix(t.Value, "y"):
			num := strings.TrimSuffix(t.Value, "y")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			v = value.NewValue(value.RTIME, fmt.Sprintf("time.Duration(%d * time.Hour)", val*24*365))
		default:
			val, err = time.ParseDuration(t.Value)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			v = value.NewValue(value.RTIME, fmt.Sprintf("time.Duration(%d * time.Second)", val))
		}

	// Combinated expressions
	case *ast.PrefixExpression:
		v, err = tf.transformPrefixExpression(t)
	case *ast.GroupedExpression:
		v, err = tf.transformGroupedExpression(t)
	case *ast.IfExpression:
		v, err = tf.transformIfExpression(expect, t)
	case *ast.InfixExpression:
		v, err = tf.transformInfixExpression(t)
	case *ast.FunctionCallExpression:
		v, err = tf.transformFunctionCallExpression(expect, t)
	}

	if err != nil {
		return nil, TransformError(&expr.GetMeta().Token, "Undefined expression found")
	}

	// Add dependent packages for the runtime
	if v.Dependencies != nil {
		for key, pkg := range v.Dependencies {
			tf.Packages.Add(key, pkg.Alias)
		}
	}

	return v.Conversion(expect), nil
}

func (tf *CoreTransformer) transformIdentValue(ident *ast.Ident) (*value.Value, error) {
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

func (tf *CoreTransformer) transformPrefixExpression(expr *ast.PrefixExpression) (*value.Value, error) {
	switch expr.Operator {
	case "!":
		right, err := tf.transformExpression(value.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = "!" + right.Code
		return right, nil
	case "-":
		right, err := tf.transformExpression(value.INTEGER, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = "-" + right.Code
		return right, nil
	case "+":
		right, err := tf.transformExpression(value.STRING, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return right, nil
	}
	return nil, TransformError(&expr.GetMeta().Token, "Unexpected prefix operator found: %s", expr.Operator)
}

func (tf *CoreTransformer) transformGroupedExpression(expr *ast.GroupedExpression) (*value.Value, error) {
	right, err := tf.transformExpression(value.NULL, expr.Right)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	right.Code = fmt.Sprintf("(%s)", right.Code)
	return right, nil
}

func (tf *CoreTransformer) transformIfExpression(expect value.VCLType, expr *ast.IfExpression) (*value.Value, error) {
	condition, err := tf.transformExpression(value.BOOL, expr.Condition)
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

	tmp := value.Temporary()
	return value.NewValue(
		expect,
		tmp,
		value.Prepare(
			condition.Prepare,
			consequence.Prepare,
			alternative.Prepare,
			fmt.Sprintf("%s := %s", tmp, alternative.Code),
			fmt.Sprintf("if %s {", condition.Code),
			fmt.Sprintf("%s = %s", tmp, consequence.Code),
			"}",
		)), nil
}

func (tf *CoreTransformer) transformInfixExpression(expr *ast.InfixExpression) (*value.Value, error) {
	switch expr.Operator {
	case "==", "!=", ">", "<", ">=", "<=":
		left, err := tf.transformExpression(value.NULL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(value.NULL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return value.NewValue(
			value.BOOL,
			fmt.Sprintf("%s %s %s", left.String(), expr.Operator, right.String()),
			value.Prepare(left.Prepare, right.Prepare),
		), nil

	// "~" or "!~" need regular expression matching
	case "~", "!~":
		left, err := tf.transformExpression(value.NULL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(value.NULL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		// If right expression is ACL, do CIDR matching
		if right.Type == value.ACL {
			var inverse string
			if expr.Operator == "!~" {
				inverse = "!"
			}
			return value.NewValue(
				value.BOOL,
				fmt.Sprintf("%s%s.Match(%s)", inverse, right.Code, left.String()),
				value.Prepare(left.Prepare, right.Prepare),
			), nil
		}

		// Otherwise, string matching, import regexp package
		tf.Packages.Add("regexp", "")
		tmp := value.Temporary()
		return value.NewValue(
			value.BOOL,
			tmp,
			value.Prepare(
				left.Prepare,
				right.Prepare,
				fmt.Sprintf("%s, err := regexp.MatchString(%s, %s)", tmp, right.String(), left.String()),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil

	case "||", "&&":
		left, err := tf.transformExpression(value.BOOL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(value.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return value.NewValue(
			value.BOOL,
			fmt.Sprintf("%s %s %s", left.String(), expr.Operator, right.String()),
			value.Prepare(left.Prepare, right.Prepare),
		), nil

	// "+" means string concatenation
	case "+":
		left, err := tf.transformExpression(value.STRING, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(value.STRING, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return value.NewValue(
			value.BOOL,
			fmt.Sprintf("%s + %s", left.String(), right.String()),
			value.Prepare(left.Prepare, right.Prepare),
		), nil
	}
	return nil, TransformError(&expr.GetMeta().Token, "Unexpected infix operator: %s", expr.Operator)
}

func (tf *CoreTransformer) transformFunctionCallExpression(
	expect value.VCLType,
	expr *ast.FunctionCallExpression,
) (*value.Value, error) {

	tf.Packages.Add("github.com/ysugimoto/vintage/builtin", "")
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(
		"vintage.%s(",
		ucFirst(strings.ReplaceAll(expr.Function.Value, ".", "_")),
	))
	for i := range expr.Arguments {
		v, err := tf.transformExpression(value.NULL, expr.Arguments[i])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(v.String())
		if i != len(expr.Arguments)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")
	return value.NewValue(expect, buf.String()), nil
}
