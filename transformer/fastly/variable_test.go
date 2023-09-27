package fastly

import (
	"strings"
	"testing"

	"github.com/ysugimoto/vintage/transformer/variable"
)

func TestVariableImplementation(t *testing.T) {

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
