package transformer

import (
	"fmt"

	"github.com/ysugimoto/vintage"
)

type ValueOption func(e *expressionValue)

func Prepare(code string) ValueOption {
	return func(e *expressionValue) {
		e.Prepare = code + lineFeed
	}
}

func Dependency(pkg, alias string) ValueOption {
	return func(e *expressionValue) {
		e.Dependencies = Packages{
			pkg: {alias},
		}
	}
}

type expressionValue struct {
	Type         vintage.VCLType
	Code         string
	Prepare      string
	Dependencies Packages
}

func ExpressionValue(t vintage.VCLType, code string, preps ...ValueOption) *expressionValue {
	v := &expressionValue{
		Type: t,
		Code: code,
	}
	for i := range preps {
		preps[i](v)
	}
	return v
}

func (v *expressionValue) Conversion(expect vintage.VCLType) *expressionValue {
	if expect == vintage.NULL {
		return v
	}

	switch expect {
	case vintage.STRING:
		return v.stringConversion()
	case vintage.BOOL:
		return v.boolConversion()
	}
	v.Type = expect
	return v
}

func (v *expressionValue) stringConversion() *expressionValue {
	if v.Type == vintage.STRING {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToString(%s)`, v.Code)
	v.Type = vintage.STRING
	return v
}

func (v *expressionValue) boolConversion() *expressionValue {
	if v.Type == vintage.BOOL {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToBool(%s)`, v.Code)
	v.Type = vintage.BOOL
	return v
}
