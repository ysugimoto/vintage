package vintage

import (
	"math/rand"
	"time"
)

type Director struct {
	Type       DirectorType
	Properties map[string]Value
	Backends   []map[string]Value
}

type DirectorOption func(d *Director)

func DirectorProperty(key string, value Value) DirectorOption {
	return func(d *Director) {
		d.Properties[key] = value
	}
}

func DirectorBackend(opts ...DirectorOption) DirectorOption {
	return func(d *Director) {
		inner := &Director{}
		for i := range opts {
			opts[i](inner)
		}
		d.Backends = append(d.Backends, inner.Properties)
	}
}

func NewDirector(name string, dType DirectorType, opts ...DirectorOption) *Backend {
	d := &Director{
		Type: dType,
	}
	for i := range opts {
		opts[i](d)
	}
	return &Backend{
		Name:     name,
		Director: d,
	}
}

func (d *Director) Backend() string {
	switch d.Type {
	case Random:
		return d.random()
	case Fallback:
		return d.fallback()
	case Hash:
		return d.hash()
	case Client:
		return d.client()
	case CHash:
		return d.chash()
	default:
		return ""
	}
}

// Elect backend by random.
// Compute@Edge won't manage backend healthness so we treat all origins are healthy
func (d *Director) random() string {
	rand.New(rand.NewSource(time.Now().Unix()))
	backend := d.Backends[rand.Intn(len(d.Backends))]
	return backend["backend"].String()
}

// Elect backend by fallback algorithm.
// Compute@Edge won't manage backend healthness so always determine the first backend
func (d *Director) fallback() string {
	backend := d.Backends[0]
	return backend["backend"].String()
}

// Elect backend by hash algorithm.
// TODO: need basement hash string like req.hash in VCL
func (d *Director) hash() string {
	backend := d.Backends[0]
	return backend["backend"].String()
}

// Elect backend by client identity.
// TODO: need basement client ip like client.identity in VCL
func (d *Director) client() string {
	backend := d.Backends[0]
	return backend["backend"].String()
}

// Elect backend by consistent hash algorithm.
// TODO: need basement hash string like req.hash in VCL
func (d *Director) chash() string {
	backend := d.Backends[0]
	return backend["backend"].String()
}
