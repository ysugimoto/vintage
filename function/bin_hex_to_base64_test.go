package function

import (
	"testing"
)

// Fastly built-in function testing implementation of bin.hex_to_base64
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/bin-hex-to-base64/
func Test_Bin_hex_to_base64(t *testing.T) {

	table := []struct {
		input   string
		expect  string
		isError bool
	}{
		{
			input:  "010203040506070809",
			expect: "AQIDBAUGBwgJ",
		},
		{
			input:  "009F00C2003300",
			expect: "AJ8AwgAzAA==",
		},
		{
			input:   "--zz",
			expect:  "",
			isError: true,
		},
		{
			input:  "",
			expect: "",
		},
		{
			input:  "61626364",
			expect: "YWJjZA==",
		},
	}

	for _, tt := range table {
		ctx := newTestRuntime()
		ret, err := Bin_hex_to_base64(ctx, tt.input)
		if tt.isError {
			if err == nil {
				t.Errorf("Expect error but got nil")
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
