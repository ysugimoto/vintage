package function

import (
	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Resp_tarpit_Name = "resp.tarpit"

// Fastly built-in function implementation of resp.tarpit
// Arguments may be:
// - INTEGER, INTEGER
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/resp-tarpit/
func Resp_tarpit[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	internalS int64,
	optional ...any,
) error {
	// @TODO: currently we do not support tarpitting, so this function has no effect
	// see: https://en.wikipedia.org/wiki/Tarpit_(networking)
	return errors.FunctionError(Resp_tarpit_Name, "Not Implemented")
}
