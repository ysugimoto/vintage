package function

import (
	"strings"
	"testing"
)

// Fastly built-in function testing implementation of querystring.filtersep
// Arguments may be:
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-filtersep/
func Test_Querystring_filtersep(t *testing.T) {
	ret, err := Querystring_filtersep(newTestRuntime())
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if !strings.EqualFold(ret, Querystring_filtersep_Sign) {
		t.Errorf("Return value unmatch, expect=%s, got=%s", Querystring_filtersep_Sign, ret)
	}
}
