package function

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ysugimoto/vintage/runtime/core"
)

// Fastly built-in function testing implementation of setcookie.get_value_by_name
// Arguments may be:
// - ID, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/setcookie-get-value-by-name/
func Test_Setcookie_get_value_by_name(t *testing.T) {
	tests := []struct {
		setCookie []string
		name      string
		where     string
		expect    string
	}{
		{
			setCookie: []string{"foo=bar"},
			name:      "baz",
			where:     "beresp",
			expect:    "",
		},
		{
			setCookie: []string{"foo=bar"},
			name:      "foo",
			where:     "beresp",
			expect:    "bar",
		},
		{
			setCookie: []string{"foo=bar", "lorem=ipsum", "lorem=ipsum2"},
			name:      "lorem",
			where:     "beresp",
			expect:    "ipsum2",
		},
		{
			setCookie: []string{"foo=bar"},
			name:      "baz",
			where:     "resp",
			expect:    "",
		},
		{
			setCookie: []string{"foo=bar"},
			name:      "foo",
			where:     "resp",
			expect:    "bar",
		},
		{
			setCookie: []string{"foo=bar", "lorem=ipsum", "lorem=ipsum2"},
			name:      "lorem",
			where:     "resp",
			expect:    "ipsum2",
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

		ret, err := Setcookie_get_value_by_name(
			ctx,
			tt.where,
			tt.name,
		)
		if err != nil {
			t.Errorf("[%d] Unexpected error: %s", i, err)
		}
		if diff := cmp.Diff(ret, tt.expect); diff != "" {
			t.Errorf("[%d] Got set-cookie value unmatch, diff=%s", i, diff)
		}
	}
}
