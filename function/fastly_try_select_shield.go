package function

import (
	"strings"

	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Fastly_try_select_shield_Name = "fastly.try_select_shield"

// Fastly built-in function implementation of fastly.try_select_shield
// Arguments may be:
// - BACKEND, BACKEND
// Reference: https://www.fastly.com/documentation/reference/vcl/functions/miscellaneous/fastly-try-select-shield/
func Fastly_try_select_shield[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	shield *vintage.Backend,
	fallback *vintage.Backend,
) (*vintage.Backend, error) {

	// If first argument is not a shield director, return fallback
	if shield.Director == nil {
		return fallback, nil
	}
	if !strings.EqualFold(string(shield.Director.Type), "shield") {
		return fallback, nil
	}

	// Note that our interpreter could not consider about director/backend is healthy.
	return shield, nil
}
