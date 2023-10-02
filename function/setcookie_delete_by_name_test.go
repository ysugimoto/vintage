package function

import (
	"testing"

	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of setcookie.delete_by_name
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/setcookie-delete-by-name/
func Test_Setcookie_delete_by_name(t *testing.T) {
	tests := []struct {
		setCookie     []string
		where         string
		ignore        string
		expect        bool
		expectCookies []string
	}{
		{
			setCookie:     []string{"foo=bar"},
			ignore:        "baz",
			expect:        false,
			where:         "beresp",
			expectCookies: []string{"foo"},
		},
		{
			setCookie:     []string{"foo=bar"},
			ignore:        "foo",
			expect:        true,
			where:         "beresp",
			expectCookies: []string{},
		},
		{
			setCookie:     []string{"foo=bar", "lorem=ipsum"},
			ignore:        "foo",
			expect:        true,
			where:         "beresp",
			expectCookies: []string{"lorem"},
		},
		{
			setCookie:     []string{"foo=bar"},
			ignore:        "baz",
			expect:        false,
			where:         "resp",
			expectCookies: []string{"foo"},
		},
		{
			setCookie:     []string{"foo=bar"},
			ignore:        "foo",
			expect:        true,
			where:         "resp",
			expectCookies: []string{},
		},
		{
			setCookie:     []string{"foo=bar", "lorem=ipsum"},
			ignore:        "foo",
			expect:        true,
			where:         "resp",
			expectCookies: []string{"lorem"},
		},
	}

	for i, tt := range tests {
		ctx := newTestRuntime()
		ctx.BackendResponseHeader = core.NewHeader(map[string][]string{
			"Set-Cookie": tt.setCookie,
		})
		ctx.ResponseHeader = core.NewHeader(map[string][]string{
			"Set-Cookie": tt.setCookie,
		})

		ret, err := Setcookie_delete_by_name(
			ctx,
			tt.where,
			tt.ignore,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if ret != tt.expect {
			t.Errorf("[%d] Return value unmatch, expect=%t, got=%t", i, tt.expect, ret)
		}
	}

}
