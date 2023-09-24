package transformer

import "github.com/ysugimoto/falco/context"

type Option func(t *transformer)

func WithSnippets(snip *context.FastlySnippet) Option {
	return func(t *transformer) {
		t.snippets = snip
	}
}
