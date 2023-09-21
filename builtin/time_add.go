// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"time"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Time_add_Name = "time.add"

var Time_add_ArgumentTypes = []value.Type{value.TimeType, value.RTimeType}

func Time_add_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Time_add_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Time_add_ArgumentTypes[i] {
			if i != 1 {
				return errors.TypeMismatch(Time_add_Name, i+1, Time_add_ArgumentTypes[i], args[i].Type())
			}
			// Second argument allows to pass bot of TIME and RTIME time
			// https://fiddle.fastly.dev/fiddle/f0098e7e
			if args[i].Type() != value.TimeType {
				return errors.TypeMismatch(Time_add_Name, i+1, value.TimeType, args[i].Type())
			}
		}
	}
	return nil
}

// Fastly built-in function implementation of time.add
// Arguments may be:
// - TIME, RTIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-add/
func Time_add(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Time_add_Validate(args); err != nil {
		return value.Null, err
	}

	t1 := value.Unwrap[*value.Time](args[0]).Value
	switch args[1].Type() {
	case value.TimeType:
		t2 := value.Unwrap[*value.Time](args[1]).Value
		return &value.Time{
			Value: t1.Add(time.Second * time.Duration(t2.Second())),
		}, nil
	case value.RTimeType:
		t2 := value.Unwrap[*value.RTime](args[1]).Value
		return &value.Time{
			Value: t1.Add(t2),
		}, nil
	default:
		// unreached, but need for linting
		return value.Null, errors.New(Time_add_Name, "Unexpected type of second argument: %s", args[1].Type())
	}

}
