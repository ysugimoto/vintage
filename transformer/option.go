package transformer

import "github.com/ysugimoto/falco/context"

type Option func(t *CoreTransformer)

func WithSnippets(snip *context.FastlySnippet) Option {
	return func(t *CoreTransformer) {
		t.snippets = snip
	}
}

func WithVariables(v Variable) Option {
	return func(t *CoreTransformer) {
		t.variables = v
	}
}
