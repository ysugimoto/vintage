package transformer

import "github.com/ysugimoto/falco/resolver"

type Transformer interface {
	Transform(resolver.Resolver) ([]byte, error)
}

type Variable interface {
	Get(name string) (*ExpressionValue, error)
	Set(name string, value *ExpressionValue) error
	Unset(name string) error
}
