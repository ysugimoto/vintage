package core

import (
	"github.com/ysugimoto/falco/context"
	"github.com/ysugimoto/vintage/transformer/variable"
)

type TransformOption func(t *CoreTransformer)

func WithSnippets(snip *context.FastlySnippet) TransformOption {
	return func(t *CoreTransformer) {
		t.snippets = snip
	}
}

func WithVariables(v variable.Variables) TransformOption {
	return func(t *CoreTransformer) {
		t.variables = v
	}
}

func WithRuntimeName(v string) TransformOption {
	return func(t *CoreTransformer) {
		t.runtimeName = v
	}
}
