package function

import (
	"testing"
)

// Fastly built-in function testing implementation of boltsort.sort
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/boltsort-sort/
func Test_Boltsort_sort(t *testing.T) {

	ret, err := Boltsort_sort(
		newTestRuntime(),
		"/foo?b=1&a=2&c=3",
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret != "/foo?a=2&b=1&c=3" {
		t.Errorf("Unexpected value returned, expect=/foo?a=2&b=1&c=3, got=%s", ret)
	}
}
