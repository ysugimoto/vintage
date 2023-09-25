package core

import (
	"fmt"

	"github.com/ysugimoto/falco/token"
)

type ErrTransform struct {
	token   *token.Token
	message string
}

func (e *ErrTransform) Error() string {
	var file string

	if e.token == nil {
		return fmt.Sprintf("[Transformer] %s", e.message)
	}
	t := *e.token
	if t.File != "" {
		file = " in" + t.File
	}

	return fmt.Sprintf("[Transformer] %s%s at line: %d, position: %d", e.message, file, t.Line, t.Position)
}

func TransformError(t *token.Token, format string, args ...any) *ErrTransform {
	return &ErrTransform{
		token:   t,
		message: fmt.Sprintf(format, args...),
	}
}
