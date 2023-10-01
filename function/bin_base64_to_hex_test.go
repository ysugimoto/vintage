package function

import (
	"testing"
)

// Fastly built-in function testing implementation of bin.base64_to_hex
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/bin-base64-to-hex/
func Test_Bin_base64_to_hex(t *testing.T) {

	table := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input:  "AQIDBAUGBwgJ",
			expect: "010203040506070809",
		},
		{
			input:  "AJ8AwgAzAA==",
			expect: "009F00C2003300",
		},
		{
			input:   "-YWJjZA==",
			expect:  "",
			isError: true,
		},
		{
			input:  "",
			expect: "",
		},
		{
			input:  "YWJjZA==",
			expect: "61626364",
		},
	}

	for _, tt := range table {
		ctx := newTestRuntime()
		ret, err := Bin_base64_to_hex(ctx, tt.input)
		if tt.isError {
			if err == nil {
				t.Errorf("Expect error but go nil")
			}
			if ctx.FastlyError != "EINVAL" {
				t.Errorf("EINVAL should be set on fastly.error")
			}
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.expect {
			t.Errorf("Unexpected value returned, expect=%s, got=%s", tt.expect, ret)
		}
	}
}
