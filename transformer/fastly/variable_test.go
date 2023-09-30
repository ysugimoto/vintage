package fastly

import (
	"strings"
	"testing"

	"github.com/ysugimoto/vintage/transformer/value"
	"github.com/ysugimoto/vintage/transformer/variable"
)

func TestVariableImplementationGet(t *testing.T) {

	v := NewFastlyVariable()
	var notImplemented []string
	for _, name := range variable.Getable {
		if _, err := v.Get(name); err != nil {
			notImplemented = append(notImplemented, name)
		}
	}

	if len(notImplemented) > 0 {
		t.Errorf("Following variables are not implemented:\n%s", strings.Join(notImplemented, "\n"))
	}
}

func TestVariableImplementationSet(t *testing.T) {

	v := NewFastlyVariable()
	var notImplemented []string
	for _, name := range variable.Setable {
		if _, err := v.Set(name, value.NewValue(value.NULL, "")); err != nil {
			notImplemented = append(notImplemented, name)
		}
	}

	if len(notImplemented) > 0 {
		t.Errorf("Following variables are not implemented:\n%s", strings.Join(notImplemented, "\n"))
	}
}
