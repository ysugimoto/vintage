package native

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type CacheEntry struct {
	Buffer    []byte
	TTL       time.Time
	EntryTime time.Time
	Hits      int
}

type intermediate struct {
	Buffer    string `json:"buffer"`
	TTL       int64  `json:"ttl"`
	EntryTime int64  `json:"entry_time"`
	Hits      int    `json:"hits"`
}

// json.Unmarshaler interface implementation
func (c *CacheEntry) UnmarshalJSON(b []byte) error {
	var v intermediate
	if err := json.Unmarshal(b, &v); err != nil {
		return errors.WithStack(err)
	}

	c.Buffer = []byte(v.Buffer)
	c.TTL = time.Unix(v.TTL, 0)
	c.EntryTime = time.Unix(v.EntryTime, 0)
	c.Hits = v.Hits
	return nil
}

// json.Marshaler interface implementation
func (c *CacheEntry) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(intermediate{
		Buffer:    string(c.Buffer),
		TTL:       c.TTL.Unix(),
		EntryTime: c.EntryTime.Unix(),
		Hits:      c.Hits,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return b, nil
}
