package core

import "github.com/ysugimoto/falco/resolver"

type Transformer interface {
	Transform(resolver.Resolver) ([]byte, error)
}
