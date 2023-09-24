package transformer

import (
	"bytes"

	"github.com/ysugimoto/vintage"
)

func typeConversion(expect vintage.VCLType, code []byte) []byte {
	var buf bytes.Buffer

	switch expect {
	case vintage.STRING:
		buf.WriteString("vintage.ToString(")
		buf.Write(code)
		buf.WriteString(")")
	case vintage.BOOL:
		buf.WriteString("vintage.ToBool(")
		buf.Write(code)
		buf.WriteString(")")
	default:
		buf.Write(code)
	}
	return buf.Bytes()
}
