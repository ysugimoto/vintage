package transformer

import (
	"fmt"
	"strings"

	"github.com/ysugimoto/vintage"
)

type ValueOption func(e *ExpressionValue)

func Prepare(preps ...string) ValueOption {
	return func(e *ExpressionValue) {
		var codes []string
		for i := range preps {
			if preps[i] == "" {
				continue
			}
			codes = append(codes, preps[i])
		}
		e.Prepare = strings.Join(codes, lineFeed) + lineFeed
	}
}

func Dependency(pkg, alias string) ValueOption {
	return func(e *ExpressionValue) {
		e.Dependencies = Packages{
			pkg: {alias},
		}
	}
}

func Comment(c string) ValueOption {
	return func(e *ExpressionValue) {
		e.Comment = " /* " + c + " */ "
	}
}

type ExpressionValue struct {
	Type         vintage.VCLType
	Code         string
	Prepare      string
	Dependencies Packages
	Comment      string
}

func NewExpressionValue(t vintage.VCLType, code string, preps ...ValueOption) *ExpressionValue {
	v := &ExpressionValue{
		Type: t,
		Code: code,
	}
	for i := range preps {
		preps[i](v)
	}
	return v
}

func (v *ExpressionValue) Conversion(expect vintage.VCLType) *ExpressionValue {
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

func (v *ExpressionValue) stringConversion() *ExpressionValue {
	if v.Type == vintage.STRING {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToString(%s)`, v.Code)
	v.Type = vintage.STRING
	return v
}

func (v *ExpressionValue) boolConversion() *ExpressionValue {
	if v.Type == vintage.BOOL {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToBool(%s)`, v.Code)
	v.Type = vintage.BOOL
	return v
}
