package core

import "github.com/ysugimoto/falco/context"

type TransformOption func(t *CoreTransformer)

func WithSnippets(snip *context.FastlySnippet) TransformOption {
	return func(t *CoreTransformer) {
		t.snippets = snip
	}
}

func WithVariables(v Variable) TransformOption {
	return func(t *CoreTransformer) {
		t.variables = v
	}
}
