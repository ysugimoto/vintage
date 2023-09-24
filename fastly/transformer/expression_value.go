package transformer

import (
	"bytes"

	"github.com/ysugimoto/vintage"
)

type expressionValue struct {
	Type vintage.VCLType
	Code []byte
}

func (e *expressionValue) String() string {
	return string(e.Code)
}

func newExpressionValue(t vintage.VCLType, code []byte) *expressionValue {
	return &expressionValue{
		Type: t,
		Code: code,
	}
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

	var buf bytes.Buffer
	buf.WriteString("vintage.ToString(")
	buf.Write(v.Code)
	buf.WriteString(")")
	v.Code = buf.Bytes()
	v.Type = vintage.STRING
	return v
}

func (v *expressionValue) boolConversion() *expressionValue {
	if v.Type == vintage.BOOL {
		return v
	}

	var buf bytes.Buffer
	buf.WriteString("vintage.ToBool(")
	buf.Write(v.Code)
	buf.WriteString(")")
	v.Code = buf.Bytes()
	v.Type = vintage.BOOL
	return v
}
