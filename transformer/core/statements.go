package core

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/ysugimoto/falco/ast"
	"github.com/ysugimoto/vintage/transformer/value"
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
			code = tf.transformReturnStatement(s)
			returnExists = true
		case *ast.SetStatement:
			code, err = tf.transformSetStatement(s)
		case *ast.AddStatement:
			code, err = tf.transformAddStatement(s)
		case *ast.UnsetStatement:
			code, err = tf.transformUnsetStatement(s)
		case *ast.RemoveStatement:
			code, err = tf.transformRemoveStatement(s)
		case *ast.LogStatement:
			code, err = tf.transformLogStatement(s)
		case *ast.SyntheticStatement:
			code, err = tf.transformSyntheticStatement(s)
		case *ast.SyntheticBase64Statement:
			code, err = tf.transformSyntheticBase64Statement(s)
		case *ast.CallStatement:
			code = tf.transformCallStatement(s)
		case *ast.RestartStatement:
			code, err = tf.transformRestartStatement(s)
		case *ast.ErrorStatement:
			code, err = tf.transformErrorStatement(s)
		case *ast.EsiStatement:
			code, err = tf.transformEsiStatement(s)
		case *ast.FunctionCallStatement:
			code, err = tf.transformFunctionCallStatement(s)
		case *ast.IfStatement:
			code, err = tf.transformIfStatement(s)
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
	switch value.VCLType(stmt.ValueType.Value) {
	case value.STRING:
		buf.WriteString(fmt.Sprintf("var local__%s string", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.STRING, "local__"+name)
	case value.INTEGER:
		buf.WriteString(fmt.Sprintf("var local__%s int64", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.INTEGER, "local__"+name)
	case value.BOOL:
		buf.WriteString(fmt.Sprintf("var local__%s bool", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.BOOL, "local__"+name)
	case value.FLOAT:
		buf.WriteString(fmt.Sprintf("var local__%s float64", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.FLOAT, "local__"+name)
	case value.BACKEND:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Backend", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.BACKEND, "local__"+name)
	case value.IP:
		buf.WriteString(fmt.Sprintf("var local__%s net.IP", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.IP, "local__"+name)
	case value.RTIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Duration", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.RTIME, "local__"+name)
	case value.TIME:
		buf.WriteString(fmt.Sprintf("var local__%s time.Time", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.TIME, "local__"+name)
	case value.ACL:
		buf.WriteString(fmt.Sprintf("var local__%s *vintage.Acl", name))
		tf.vars[stmt.Name.Value] = value.NewValue(value.ACL, "local__"+name)
	default:
		return nil, errors.WithStack(
			fmt.Errorf("Unexpected variable type declared: %s", stmt.ValueType.Value),
		)
	}
	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformReturnStatement(stmt *ast.ReturnStatement) []byte {
	state := "NONE"
	if stmt.ReturnExpression != nil {
		state = strings.ToUpper(strings.Trim(toString(*stmt.ReturnExpression), `"`))
	}
	return []byte(fmt.Sprintf("return vintage.%s, nil", state))
}

func (tf *CoreTransformer) transformSetStatement(stmt *ast.SetStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value

	// If statement name starts with var., it means local variable access
	if strings.HasPrefix(name, "var.") {
		v, ok := tf.vars[name]
		if !ok {
			return nil, TransformError(&stmt.Token, "local variable %s undefined", name)
		}
		val, err := tf.transformExpression(v.Type, stmt.Value)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(val.Prepare + v.Code + " = " + val.String())
		tf.Packages.Merge(val.Dependencies)
		return buf.Bytes(), nil
	}

	// Otherwise, global variable access
	val, err := tf.transformExpression(value.NULL, stmt.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	code, err := tf.variables.Set(name, val)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(code.Prepare + code.String())
	tf.Packages.Merge(code.Dependencies)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformAddStatement(stmt *ast.AddStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value

	// Add statement only enables for HTTP headers
	var field, key string
	switch {
	case strings.HasPrefix(name, "req.http."):
		field = "RequestHeader"
		key = strings.TrimPrefix(name, "req.http.")
	case strings.HasPrefix(name, "bereq.http."):
		field = "BackendRequestHeader"
		key = strings.TrimPrefix(name, "bereq.http.")
	case strings.HasPrefix(name, "beresp.http."):
		field = "BackendResponseHeader"
		key = strings.TrimPrefix(name, "beresp.http.")
	case strings.HasPrefix(name, "obj.http."):
		field = "BackendResponseHeader"
		key = strings.TrimPrefix(name, "obj.http.")
	case strings.HasPrefix(name, "resp.http."):
		field = "ResponseHeader"
		key = strings.TrimPrefix(name, "resp.http.")
	default:
		return nil, TransformError(&stmt.Token, "Cannot use add statement for %s, only for HTTP header manipulation", name)
	}
	val, err := tf.transformExpression(value.STRING, stmt.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(val.Prepare)
	buf.WriteString(
		fmt.Sprintf(`ctx.%s.Add("%s", %s)`, field, key, val.Conversion(value.STRING).String()),
	)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformUnsetStatement(stmt *ast.UnsetStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value

	// If statement name starts with var., it means local variable access
	if strings.HasPrefix(name, "var.") {
		v, ok := tf.vars[name]
		if !ok {
			return nil, TransformError(&stmt.Token, "local variable %s is undefined", name)
		}
		buf.WriteString(v.Code + " = " + value.DefaultValue(v.Type))
		return buf.Bytes(), nil
	}

	// Otherwise, global variable access
	code, err := tf.variables.Unset(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(code.Prepare + code.String())

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformRemoveStatement(stmt *ast.RemoveStatement) ([]byte, error) {
	var buf bytes.Buffer

	name := stmt.Ident.Value

	// If statememt name starts with var., it means local variable access
	if strings.HasPrefix(name, "var.") {
		v, ok := tf.vars[name]
		if !ok {
			return nil, TransformError(&stmt.Token, "local variable %s undefined", name)
		}
		buf.WriteString(v.Code + " = " + value.DefaultValue(v.Type))
		return buf.Bytes(), nil
	}

	// Otherwise, global variable access
	code, err := tf.variables.Unset(name)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(code.Prepare + code.String())

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformLogStatement(stmt *ast.LogStatement) ([]byte, error) {
	var buf bytes.Buffer

	val, err := tf.transformExpression(value.STRING, stmt.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Fastly log statement format should be:
	// log [type] service_id [name] :: [message]
	// We need to find "name" section from transformed exprssion string
	spl := strings.SplitN(val.String(), "::", 2)
	if len(spl) != 2 {
		// Not a realtime log streaming, skip this statement
		return buf.Bytes(), nil
		// return nil, TransformError(&stmt.Token, "Invalid logging format found")
	}
	// Split by whitespace for getting metadata
	meta := strings.Split(strings.TrimSpace(spl[0]), " ")
	// Last segment will be logging endpoint name
	name := meta[len(meta)-1]
	log, ok := tf.loggingEndpoints[name]
	if !ok {
		return nil, TransformError(&stmt.Token, "Logging endpoint %s not found", name)
	}
	buf.WriteString(val.Prepare)

	// Lacking double quote in Write() argument is not a typo.
	// On getting logging endpoint name, quoted string is splitted and then message only has trailing double quote.
	// So we will only add prefix double quote in the argument.
	buf.WriteString(fmt.Sprintf(`if err := %s.Write("%s); err != nil {`+lineFeed, log, strings.TrimSpace(spl[1])))
	buf.WriteString("return vintage.NONE, err\n}")
	tf.Packages.Merge(val.Dependencies)

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformSyntheticStatement(stmt *ast.SyntheticStatement) ([]byte, error) {
	var buf bytes.Buffer

	val, err := tf.transformExpression(value.STRING, stmt.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(val.Prepare)
	buf.WriteString(
		fmt.Sprintf("ctx.BackendResponse.Body = io.NopCloser(strings.NewReader(%s))", val.String()),
	)
	tf.Packages.Add("io", "")
	tf.Packages.Add("strings", "")

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformSyntheticBase64Statement(stmt *ast.SyntheticBase64Statement) ([]byte, error) {
	var buf bytes.Buffer

	val, err := tf.transformExpression(value.STRING, stmt.Value)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(val.Prepare)
	buf.WriteString(
		fmt.Sprintf("ctx.BackendResponse.Body = io.NopCloser(base64.NewDecoder(base64.StdEncoding, strings.NewReader(%s)))", val.String()),
	)
	tf.Packages.Add("io", "")
	tf.Packages.Add("encoding/base64", "")
	tf.Packages.Add("strings", "")

	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformCallStatement(stmt *ast.CallStatement) []byte {
	var buf bytes.Buffer

	tmp := value.Temporary()
	buf.WriteString(strings.Join([]string{
		fmt.Sprintf("if %s, err := %s(ctx); err != nil {", tmp, stmt.Subroutine.Value),
		"return vintage.NONE, err",
		fmt.Sprintf("} else if %s != vintage.NONE {", tmp),
		fmt.Sprintf("return %s, nil", tmp),
		"}",
	}, lineFeed))

	return buf.Bytes()
}

// nolint:unparam // stmt may be used for a foture implementation
func (tf *CoreTransformer) transformRestartStatement(stmt *ast.RestartStatement) ([]byte, error) {
	return []byte("return vintage.RESTART, nil"), nil
}

func (tf *CoreTransformer) transformErrorStatement(stmt *ast.ErrorStatement) ([]byte, error) {
	var buf bytes.Buffer

	status, err := tf.transformExpression(value.INTEGER, stmt.Code)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.WriteString(fmt.Sprintf("ctx.ObjectStatus = %s"+lineFeed, status.String()))

	if stmt.Argument != nil {
		response, err := tf.transformExpression(value.STRING, stmt.Argument)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.WriteString(fmt.Sprintf("ctx.ObjectResponse = %s"+lineFeed, response.String()))
	}

	buf.WriteString("return vintage.ERROR, nil")
	return buf.Bytes(), nil
}

// nolint:unparam // stmt may be used for a foture implementation
func (tf *CoreTransformer) transformEsiStatement(stmt *ast.EsiStatement) ([]byte, error) {
	// Nothing to do for ESI
	return []byte("// Trimmed esi statement"), nil
}

func (tf *CoreTransformer) transformFunctionCallStatement(stmt *ast.FunctionCallStatement) ([]byte, error) {
	var buf bytes.Buffer

	fn, ok := builtinFunctions[stmt.Function.Value]
	if !ok {
		return nil, errors.WithStack(
			fmt.Errorf("Undefined function %s", stmt.Function.Value),
		)
	} else if fn.ReturnType != value.NULL {
		return nil, errors.WithStack(
			fmt.Errorf("Function %s has a return value, cannot call on the statement,", stmt.Function.Value),
		)
	}

	if len(stmt.Arguments) < len(fn.Requires) {
		return nil, errors.WithStack(fmt.Errorf(
			"Not enough arguments for %s, expects=%d, actual=%d",
			stmt.Function.Value,
			len(fn.Requires),
			len(stmt.Arguments),
		))
	}

	tf.Packages.Add("github.com/ysugimoto/vintage/function", "")
	code := fn.Name + "(ctx.Runtime, "
	var prepares string
	var arguments []string
	var argIndex int
	for i := range fn.Requires {
		arg, err := tf.transformExpression(fn.Requires[i], stmt.Arguments[argIndex])
		if err != nil {
			return nil, errors.WithStack(err)
		}
		prepares += arg.Prepare
		arguments = append(arguments, arg.String())
		argIndex++
	}
	if len(stmt.Arguments) > len(fn.Requires) {
		if fn.VariadicIndex > 0 {
			for _, variadic := range stmt.Arguments[fn.VariadicIndex:] {
				arg, err := tf.transformExpression(value.STRING, variadic)
				if err != nil {
					return nil, errors.WithStack(err)
				}
				prepares += arg.Prepare
				arguments = append(arguments, arg.String())
			}
		} else {
			for _, optional := range fn.Optionals {
				arg, err := tf.transformExpression(optional, stmt.Arguments[argIndex])
				if err != nil {
					return nil, errors.WithStack(err)
				}
				prepares += arg.Prepare
				arguments = append(arguments, arg.String())
				argIndex++
			}
		}
	}

	code += strings.Join(arguments, ", ") + ")"
	buf.WriteString(prepares)
	buf.WriteString(
		fmt.Sprintf("if _, err := %s; err != nil {\nreturn vintage.NONE, err\n}", code),
	)
	return buf.Bytes(), nil
}

func (tf *CoreTransformer) transformIfStatement(stmt *ast.IfStatement) ([]byte, error) {
	var buf bytes.Buffer

	var prepares []byte
	condition, err := tf.transformExpression(value.BOOL, stmt.Condition)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	prepares = append(prepares, []byte(condition.Prepare)...)
	buf.WriteString(fmt.Sprintf("if %s {"+lineFeed, condition.String()))
	consequence, _, err := tf.transformBlockStatement(stmt.Consequence.Statements)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	buf.Write(consequence)
	buf.WriteString("}")

	for _, elseif := range stmt.Another {
		condition, err := tf.transformExpression(value.BOOL, elseif.Condition)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		prepares = append(prepares, []byte(condition.Prepare)...)
		buf.WriteString(fmt.Sprintf(" else if %s {"+lineFeed, condition.String()))
		consequence, _, err := tf.transformBlockStatement(elseif.Consequence.Statements)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(consequence)
		buf.WriteString("}")
	}

	if stmt.Alternative != nil {
		buf.WriteString(" else {" + lineFeed)
		alternative, _, err := tf.transformBlockStatement(stmt.Alternative.Statements)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		buf.Write(alternative)
		buf.WriteString("}")
	}

	return append(prepares, buf.Bytes()...), nil
}
