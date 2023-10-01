package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type RuntimeError struct {
	err error
}

func (r *RuntimeError) Error() string {
	return r.err.Error()
}

func FunctionError(name, format string, args ...any) *RuntimeError {
	return &RuntimeError{
		err: errors.WithStack(
			fmt.Errorf("["+name+"] "+format, args...),
		),
	}
}

func WithStack(e error) error {
	return errors.WithStack(e)
}
