package function

import (
	"testing"
)

// Fastly built-in function testing implementation of http_status_matches
// Arguments may be:
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/http-status-matches/
func Test_Http_status_matches(t *testing.T) {

	tests := []struct {
		status int64
		format string
		expect bool
	}{
		{status: 200, format: "200,302,500", expect: true},
		{status: 200, format: "!200,302,500", expect: false},
		{status: 400, format: "200,302,500", expect: false},
		{status: 400, format: "!200,302,500", expect: true},
	}

	for _, tt := range tests {
		ret, err := Http_status_matches(
			newTestRuntime(),
			tt.status,
			tt.format,
		)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.expect {
			t.Errorf(
				"return value is unmatch, %d, %s, expect=%t, got=%t",
				tt.status, tt.format, tt.expect, ret,
			)
		}
	}
}
