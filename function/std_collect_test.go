package function

import (
	"testing"
)

// Fastly built-in function testing implementation of std.collect
// Arguments may be:
// - ID
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/std-collect/
func Test_Std_collect(t *testing.T) {
	t.Skip("Skip std.collect function because headers are automatically collected")
}
