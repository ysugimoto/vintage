package vintage

import (
	"context"
	"net/textproto"
	"time"

	"github.com/avct/uasurfer"
)

type Runtime interface {
	Proxy(ctx context.Context, backend *Backend) error
	// Cache(ctx context.Context, key string) error
}

type Subroutine[T Runtime] func(ctx T) (State, error)

type Resource[T Runtime] func(c *Context[T])

func BackendResource[T Runtime](name string, v *Backend) Resource[T] {
	return func(c *Context[T]) {
		c.Backends[name] = v
		if v.IsDefault {
			c.Backend = v
		}
	}
}
func AclResource[T Runtime](name string, v *Acl) Resource[T] {
	return func(c *Context[T]) {
		c.Acls[name] = v
	}
}
func TableResource[T Runtime](name string, v *Table) Resource[T] {
	return func(c *Context[T]) {
		c.Tables[name] = v
	}
}
func SubroutineResource[T Runtime](name string, v Subroutine[T]) Resource[T] {
	return func(c *Context[T]) {
		c.Subroutines[name] = v
	}
}

func (c *Context[T]) Register(resources ...Resource[T]) {
	for i := range resources {
		resources[i](c)
	}
}

type Context[T Runtime] struct {
	Backend          *Backend
	RequestHeader    *Header
	ResponseHeader   *Header
	RequestStartTime time.Time
	RequestEndTime   time.Time
	Restarts         int
	RequestHash      string

	// We should implement User-Agent related matcher by ourselves
	// but we trust github.com/avct/uasurfer package for now
	UserAgent *uasurfer.UserAgent

	// Declaration stacks
	Backends    map[string]*Backend
	Acls        map[string]*Acl
	Tables      map[string]*Table
	Subroutines map[string]Subroutine[T]

	// Properties that will be assigned in the process
}

func NewContext[T Runtime](w, r map[string][]string, resources ...Resource[T]) *Context[T] {
	c := &Context[T]{
		RequestHeader:    NewHeader(textproto.MIMEHeader(r)),
		ResponseHeader:   NewHeader(textproto.MIMEHeader(w)),
		RequestStartTime: time.Now(),

		Backends:    make(map[string]*Backend),
		Acls:        make(map[string]*Acl),
		Tables:      make(map[string]*Table),
		Subroutines: make(map[string]Subroutine[T]),
	}
	for i := range resources {
		resources[i](c)
	}
	c.UserAgent = uasurfer.Parse(c.RequestHeader.Get("User-Agent"))
	return c
}
