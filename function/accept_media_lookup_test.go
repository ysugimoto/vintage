package function

import (
	"testing"
)

// Fastly built-in function testing implementation of accept.media_lookup
// Arguments may be:
// - STRING, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-media-lookup/
func Test_Accept_media_lookup(t *testing.T) {
	tests := []struct {
		Accept string
		Expect string
	}{
		{
			Accept: "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			Expect: "text/html",
		},
		{
			Accept: "application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			Expect: "text/plain",
		},
		{
			Accept: "application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,image/*,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			Expect: "image/tiff",
		},
	}

	for _, tt := range tests {
		ret, err := Accept_media_lookup(
			newTestRuntime(),
			"image/jpeg:image/png",
			"text/plain",
			"image/tiff:text/html",
			tt.Accept,
		)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if ret != tt.Expect {
			t.Errorf("Unexpected value returned, expect=%s, got=%s", tt.Expect, ret)
		}
	}
}
