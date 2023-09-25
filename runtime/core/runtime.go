package core

import (
	"fmt"
	"net/textproto"
	"time"

	"github.com/ysugimoto/vintage"
)

type Subroutine[T EdgeRuntime] func(ctx T) (State, error)

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

func (c *Runtime[T]) Register(resources ...Resource[T]) {
	for i := range resources {
		resources[i](c)
	}
}

type Runtime[T EdgeRuntime] struct {
	Backend            *vintage.Backend
	RequestHeader      *Header
	ResponseHeader     *Header
	RequestStartTime   time.Time
	RequestEndTime     time.Time
	Restarts           int
	RequestHash        string
	RequestHeaderBytes int64

	// We should implement User-Agent related matcher by ourselves
	// but we trust github.com/avct/uasurfer package for now
	UserAgent *UserAgent

	// Declaration stacks
	Backends    map[string]*vintage.Backend
	Acls        map[string]*vintage.Acl
	Tables      map[string]*vintage.Table
	Subroutines map[string]Subroutine[T]

	// Properties that will be assigned in the process
	MaxStaleIfError                 time.Duration
	MaxStaleWhileRevalidate         time.Duration
	ClientSocketCongestionAlgorithm string
	EsiAllowInsideData              bool
	EnableESI                       bool
	ESILevel                        int64
	IsLocallyGenerated              bool
	TimeToFirstByte                 time.Duration
}

func NewRuntime[T EdgeRuntime](w, r map[string][]string, resources ...Resource[T]) *Runtime[T] {
	c := &Runtime[T]{
		RequestHeader:    NewHeader(textproto.MIMEHeader(r)),
		ResponseHeader:   NewHeader(textproto.MIMEHeader(w)),
		RequestStartTime: time.Now(),

		Backends:    make(map[string]*vintage.Backend),
		Acls:        make(map[string]*vintage.Acl),
		Tables:      make(map[string]*vintage.Table),
		Subroutines: make(map[string]Subroutine[T]),

		// Default variable values
		ClientSocketCongestionAlgorithm: "cubic",
		EsiAllowInsideData:              false,
		EnableESI:                       false,
		ESILevel:                        0,
		IsLocallyGenerated:              false,
	}
	for i := range resources {
		resources[i](c)
	}
	c.UserAgent = NewUserAgent(c.RequestHeader.Get("User-Agent"))
	c.RequestHeaderBytes = c.factoryInitRequestHeaderBytes(r)
	return c
}

func (c *Runtime[T]) factoryInitRequestHeaderBytes(hs map[string][]string) int64 {
	var size int64
	for key, val := range hs {
		size += int64(len(key))
		for i := range val {
			size += int64(len(val[i])) + 1 // add 1 byte of comma
		}
		size -= 1 // minus last comma
	}
	return size
}

// used for now.sec
func (c *Runtime[T]) NowSec() string {
	return fmt.Sprint(time.Now().Unix())
}
