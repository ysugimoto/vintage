package function

import (
	"strings"
	"testing"
)

// Fastly built-in function testing implementation of randomstr
// Arguments may be:
// - INTEGER
// - INTEGER, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randomstr/
func Test_Randomstr(t *testing.T) {
	tests := []struct {
		length     int64
		characters string
	}{
		{length: 20},
		{length: 10, characters: "1234567890abcdef"},
		{length: 5, characters: "abcdef"},
	}

	for i, tt := range tests {
		for j := 0; j < 10000; j++ {
			var ret string
			var err error
			if tt.characters == "" {
				ret, err = Randomstr(newTestRuntime(), tt.length)
			} else {
				ret, err = Randomstr(newTestRuntime(), tt.length, tt.characters)
			}

			if err != nil {
				t.Errorf("[%d] Unexpected error: %s", i, err)
			}

			for _, s := range ret {
				chars := string(Randomstr_default_characters)
				if tt.characters != "" {
					chars = tt.characters
				}
				if !strings.Contains(chars, string(s)) {
					t.Errorf(
						"[%d] Unexpected return value, character %s should be once of %s",
						i, string(s), chars,
					)
				}
			}
		}
	}
}
