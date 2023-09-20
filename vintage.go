package vintage

import (
	"github.com/ysugimoto/vintage/context"
)

type Runtime struct {
	ctx *context.Context
}

func NewVCLRuntime() *Runtime {
	return &Runtime{
		ctx: context.New(),
	}
}
