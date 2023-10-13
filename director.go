package vintage

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Director struct {
	Type              DirectorType
	Properties        map[string]any
	Backends          []map[string]any
	determinedBackend string
}

type DirectorOption func(d *Director)

func DirectorProperty(key string, value any) DirectorOption {
	return func(d *Director) {
		d.Properties[key] = value
	}
}

func DirectorBackend(opts ...DirectorOption) DirectorOption {
	return func(d *Director) {
		inner := &Director{
			Properties: make(map[string]any),
		}
		for i := range opts {
			opts[i](inner)
		}
		d.Backends = append(d.Backends, inner.Properties)
	}
}

func NewDirector(name string, dType DirectorType, opts ...DirectorOption) *Backend {
	d := &Director{
		Type:       dType,
		Properties: make(map[string]any),
		Backends:   []map[string]any{},
	}
	for i := range opts {
		opts[i](d)
	}
	return &Backend{
		Name:     name,
		Director: d,
	}
}

func (d *Director) Backend(ident RequestIdentity) string {
	if d.determinedBackend == "" {
		var backend string
		switch d.Type {
		case Random:
			backend = d.random()
		case Fallback:
			backend = d.fallback()
		case Hash:
			backend = d.hash(ident)
		case Client:
			backend = d.client(ident)
		case CHash:
			backend = d.chash(ident)
		}

		d.determinedBackend = backend
	}
	return d.determinedBackend
}

// Elect backend by random.
// NOTE: Fastly Compute does not know that backend is healthy or not,
// so runtime does not consider about quorum weight
func (d *Director) random() string {
	lottery := make([]int, 1000)
	var last int
	for index, v := range d.Backends {
		var weight int
		if w, ok := v["weight"].(int); ok {
			weight = w
		}
		for i := 0; i < weight; i++ {
			lottery[last] = index
			last++
			if last > len(lottery) {
				extend := make([]int, 2000)
				for j := 0; j < len(lottery); j++ {
					extend[j] = lottery[i]
				}
				lottery = extend
			}
		}
	}
	rand.New(rand.NewSource(time.Now().Unix()))
	lottery = lottery[0:last]
	backend := d.Backends[lottery[rand.Intn(last)]]
	return backend["backend"].(string)
}

// Elect backend by fallback algorithm.
// NOTE: Fastly Compute does not know that backend is healthy or not,
// so runtime always chooses first backend
func (d *Director) fallback() string {
	backend := d.Backends[0]
	return backend["backend"].(string)
}

// Elect backend by hash algorithm.
// TODO: need basement hash string like req.hash in VCL
func (d *Director) hash(ident RequestIdentity) string {
	hash := sha256.Sum256([]byte(ident.Hash))
	return d.getBackendByHash(hash[:])
}

// Elect backend by client identity.
// TODO: need basement client ip like client.identity in VCL
func (d *Director) client(ident RequestIdentity) string {
	hash := sha256.Sum256([]byte(ident.Client))
	return d.getBackendByHash(hash[:])
}

// Elect backend by consistent hash algorithm.
// TODO: need basement hash string like req.hash in VCL
func (d *Director) chash(ident RequestIdentity) string {
	var circles []uint32
	var max uint32 = 10000
	hashTable := make(map[uint32]string)

	var seed uint32
	if s, ok := d.Properties["seed"].(int); ok {
		seed = uint32(s)
	}

	for _, v := range d.Backends {
		backend := v["backend"].(string) // nolint:errcheck
		// typically loop three times in order to find suitable ring position
		for i := 0; i < 3; i++ {
			buf := make([]byte, 4)
			binary.BigEndian.PutUint32(buf, seed)
			hash := sha256.New()
			hash.Write(buf)
			hash.Write([]byte(backend))
			hash.Write([]byte(fmt.Sprint(i)))
			h := hash.Sum(nil)
			num := binary.BigEndian.Uint32(h[:8]) % max
			hashTable[num] = backend
			circles = append(circles, num)
		}
	}

	// sort slice for binary search
	sort.Slice(circles, func(i, j int) bool {
		return circles[i] < circles[j]
	})

	var hashKey [32]byte
	key := "client"
	if v, ok := d.Properties["key"].(string); ok {
		key = v
	}
	switch key {
	case "object":
		hashKey = sha256.Sum256([]byte(ident.Hash))
	default: // client
		hashKey = sha256.Sum256([]byte(ident.Client))
	}

	k := binary.BigEndian.Uint32(hashKey[:8]) % max
	index := sort.Search(len(circles), func(i int) bool {
		return circles[i] >= k
	})
	if index == len(circles) {
		index = 0
	}

	return hashTable[circles[index]]
}

func (d *Director) getBackendByHash(hash []byte) string {
	var target string // determined backend name

	for m := 4; m <= 16; m += 2 {
		max := uint64(math.Pow(10, float64(m)))
		num := binary.BigEndian.Uint64(hash[:8]) % max

		for _, v := range d.Backends {
			backend := v["backend"].(string) // nolint:errcheck
			bh := sha256.Sum256([]byte(backend))
			b := binary.BigEndian.Uint64(bh[:8])
			if b%(max*10) >= num && b%(max*10) < num+max {
				target = backend
				goto DETERMINED
			}
		}
	}
DETERMINED:

	// When target is not determined, use first healthy backend
	if target == "" {
		target = d.Backends[0]["backend"].(string) // nolint:errcheck
	}
	return target
}
