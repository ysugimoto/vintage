package core

import "github.com/ysugimoto/vintage"

type Subroutine[T EdgeRuntime] func(ctx T) (vintage.State, error)

type Resource[T EdgeRuntime] func(c *Runtime[T])

func BackendResource[T EdgeRuntime](name string, v *vintage.Backend) Resource[T] {
	return func(c *Runtime[T]) {
		c.Backends[name] = v
		if v.IsDefault {
			c.Backend = v
		}
	}
}

func AclResource[T EdgeRuntime](name string, v *vintage.Acl) Resource[T] {
	return func(c *Runtime[T]) {
		c.Acls[name] = v
	}
}

func TableResource[T EdgeRuntime](name string, v *vintage.Table) Resource[T] {
	return func(c *Runtime[T]) {
		c.Tables[name] = v
	}
}

func SubroutineResource[T EdgeRuntime](name string, v Subroutine[T]) Resource[T] {
	return func(c *Runtime[T]) {
		c.Subroutines[name] = v
	}
}

func LoggingResource[T EdgeRuntime](name string, v *vintage.LoggingEndpoint) Resource[T] {
	return func(c *Runtime[T]) {
		c.LoggingEndpoints[name] = v
	}
}

func (c *Runtime[T]) Register(resources ...Resource[T]) {
	for i := range resources {
		resources[i](c)
	}
}
