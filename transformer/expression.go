package transformer

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage"
)

func (tf *CoreTransformer) transformExpression(
	expect vintage.VCLType,
	expr ast.Expression,
) (*ExpressionValue, error) {

	var value *ExpressionValue
	var err error

	switch t := expr.(type) {
	case *ast.Ident:
		value, err = tf.transformIdentValue(t)
	case *ast.IP:
		value = NewExpressionValue(vintage.IP, `"`+net.ParseIP(t.Value).String()+`"`)
	case *ast.Boolean:
		value = NewExpressionValue(vintage.BOOL, fmt.Sprintf("%t", t.Value))
	case *ast.Integer:
		value = NewExpressionValue(vintage.INTEGER, fmt.Sprint(t.Value))
	case *ast.Float:
		value = NewExpressionValue(vintage.FLOAT, fmt.Sprint(t.Value))
	case *ast.String:
		value = NewExpressionValue(vintage.STRING, `"`+t.Value+`"`)
	case *ast.RTime:
		var val time.Duration
		switch {
		case strings.HasSuffix(t.Value, "d"):
			num := strings.TrimSuffix(t.Value, "d")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = NewExpressionValue(vintage.RTIME, fmt.Sprintf("time.Duration(%d * time.Hour)", val*24))
		case strings.HasSuffix(t.Value, "y"):
			num := strings.TrimSuffix(t.Value, "y")
			val, err = time.ParseDuration(num + "h")
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = NewExpressionValue(vintage.RTIME, fmt.Sprintf("time.Duration(%d * time.Hour)", val*24*365))
		default:
			val, err = time.ParseDuration(t.Value)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			value = NewExpressionValue(vintage.RTIME, fmt.Sprintf("time.Duration(%d * time.Second)", val))
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

	// Add dependent packages for the runtime
	if value.Dependencies != nil {
		for key, pkg := range value.Dependencies {
			tf.Packages.Add(key, pkg.Alias)
		}
	}

	return value.Conversion(expect), nil
}

func (tf *CoreTransformer) transformIdentValue(ident *ast.Ident) (*ExpressionValue, error) {
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

func (tf *CoreTransformer) transformPrefixExpression(expr *ast.PrefixExpression) (*ExpressionValue, error) {
	switch expr.Operator {
	case "!":
		right, err := tf.transformExpression(vintage.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = "!" + right.Code
		return right, nil
	case "-":
		right, err := tf.transformExpression(vintage.INTEGER, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right.Code = "-" + right.Code
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

func (tf *CoreTransformer) transformGroupedExpression(expr *ast.GroupedExpression) (*ExpressionValue, error) {
	right, err := tf.transformExpression(vintage.NULL, expr.Right)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	right.Code = fmt.Sprintf("(%s)", right.Code)
	return right, nil
}

func (tf *CoreTransformer) transformIfExpression(expect vintage.VCLType, expr *ast.IfExpression) (*ExpressionValue, error) {
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

	v := tmpVar()
	preps := prepareCodes(
		condition.Prepare,
		consequence.Prepare,
		alternative.Prepare,
		fmt.Sprintf("%s := %s", v, alternative.Code),
		fmt.Sprintf("if %s {", condition.Code),
		fmt.Sprintf(consequence.Code),
		"}",
	)
	return NewExpressionValue(expect, v, Prepare(preps)), nil
}

func (tf *CoreTransformer) transformInfixExpression(expr *ast.InfixExpression) (*ExpressionValue, error) {
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

		return NewExpressionValue(
			vintage.BOOL,
			fmt.Sprintf("%s %s %s", left.Code, expr.Operator, right.Code),
			Prepare(prepareCodes(left.Prepare, right.Prepare)),
		), nil

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
			var inverse string
			if expr.Operator == "!~" {
				inverse = "!"
			}
			return NewExpressionValue(
				vintage.BOOL,
				fmt.Sprintf("%s%s.Match(%s)", inverse, right.Code, left.Code),
				Prepare(prepareCodes(left.Prepare, right.Prepare)),
			), nil
		}

		// Otherwise, string matching, import regexp package
		tf.Packages.Add("regexp", "")
		v := tmpVar()
		preps := prepareCodes(
			left.Prepare,
			right.Prepare,
			fmt.Sprintf("%s, err := regexp.MatchString(%s, %s)", v, right.Code, left.Code),
			"if err != nil {",
			"return vintage.NONE, err",
			"}",
		)
		return NewExpressionValue(vintage.BOOL, v, Prepare(preps)), nil

	case "||", "&&":
		left, err := tf.transformExpression(vintage.BOOL, expr.Left)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		right, err := tf.transformExpression(vintage.BOOL, expr.Right)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return NewExpressionValue(
			vintage.BOOL,
			fmt.Sprintf("%s %s %s", left.Code, expr.Operator, right.Code),
			Prepare(prepareCodes(left.Prepare, right.Prepare)),
		), nil

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
		return NewExpressionValue(
			vintage.BOOL,
			fmt.Sprintf("%s + %s", left.Code, right.Code),
			Prepare(prepareCodes(left.Prepare, right.Prepare)),
		), nil
	}
	return nil, TransformError(&expr.GetMeta().Token, "Unexpected infix operator: %s", expr.Operator)
}

func (tf *CoreTransformer) transformFunctionCallExpression(
	expect vintage.VCLType,
	expr *ast.FunctionCallExpression,
) (*ExpressionValue, error) {

	tf.Packages.Add("github.com/ysugimoto/vintage/builtin", "")

	// TODO: add function arguments expression
	call := fmt.Sprintf(
		"vintage.%s()",
		ucFirst(strings.ReplaceAll(expr.Function.Value, ".", "_")),
	)
	return NewExpressionValue(expect, call), nil
}
