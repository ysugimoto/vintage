package variable

import "github.com/ysugimoto/vintage/transformer/value"

type Variables interface {
	Get(name string) (*value.Value, error)
	Set(name string, value *value.Value) error
	Unset(name string) error
}
