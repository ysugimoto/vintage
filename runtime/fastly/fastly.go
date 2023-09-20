package fastly

import (
	"github.com/ysugimoto/vintage"
)

type Runtime struct {
	backends map[string]*vintage.Backend
	acls     map[string]*vintage.Acl
	tables   map[string]vintage.Table
}

func RuntimeContext() *Runtime {
	return &Runtime{
		backends: make(map[string]*vintage.Backend),
		acls:     make(map[string]*vintage.Acl),
		tables:   make(map[string]vintage.Table),
	}
}

func (r *Runtime) Register(name string, item any) {
	switch t := item.(type) {
	case *vintage.Backend:
		r.backends[name] = t
	}
}
