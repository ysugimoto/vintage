package lib

import (
	"net/url"
	"sort"
	"strings"
)

type QueryString struct {
	Key   string
	Value []string // nil indicates not set in VCL
}

// We implement original querytring struct in order to manage URL queries that keep its present order.
// url.Values are useful *url.Values.Encode() does not care the order because it is managed in map and is sorted by query name.
// This struct keeps query present order and we will use this strut for query string manipulation.
type QueryStrings struct {
	Prefix string // protocol, host, port, path
	Items  []*QueryString
}

// ParseQuery parses raw query string and return QueryStrings pointer
func ParseQuery(qs string) (*QueryStrings, error) {
	// Find querystring sign
	idx := strings.Index(qs, "?")
	if idx == -1 {
		return &QueryStrings{Prefix: qs}, nil
	}

	ret := &QueryStrings{
		Prefix: qs[0:idx],
	}
	qs = qs[idx+1:]
	for _, q := range strings.Split(qs, "&") {
		sp := strings.SplitN(q, "=", 2)
		if len(sp) == 0 {
			continue
		}
		key, err := url.QueryUnescape(sp[0])
		if err != nil {
			return nil, err
		}
		if len(sp) == 1 {
			// e.g ?foo -- equal sign is not present
			ret.Items = append(ret.Items, &QueryString{Key: key, Value: nil})
			continue
		}
		val, err := url.QueryUnescape(sp[1])
		if err != nil {
			return nil, err
		}
		ret.Add(key, val)
	}
	return ret, nil
}

// Set sets a new query string. If key exists, overwrite it
func (q *QueryStrings) Set(name, val string) {
	for i := range q.Items {
		if q.Items[i].Key != name {
			continue
		}
		q.Items[i].Value = []string{val}
		return
	}

	// set new
	q.Items = append(q.Items, &QueryString{Key: name, Value: []string{val}})
}

// Add adds a new query string. If key exists, append it
func (q *QueryStrings) Add(name, val string) {
	for i := range q.Items {
		if q.Items[i].Key != name {
			continue
		}
		if q.Items[i].Value == nil {
			q.Items[i].Value = []string{}
		}
		q.Items[i].Value = append(q.Items[i].Value, val)
		return
	}

	// append new
	q.Items = append(q.Items, &QueryString{Key: name, Value: []string{val}})
}

// Get gets query string value for key
func (q *QueryStrings) Get(name string) *string {
	for i := range q.Items {
		if q.Items[i].Key != name {
			continue
		}
		if q.Items[i].Value == nil {
			return nil // nil returns not set string in VCL
		}
		return &q.Items[i].Value[0]
	}
	return nil
}

// Clean filters empty key query string
func (q *QueryStrings) Clean() {
	var cleaned []*QueryString
	for _, v := range q.Items {
		if v.Key == "" {
			continue
		}
		cleaned = append(cleaned, v)
	}
	q.Items = cleaned
}

// Filter filters query string which filter function returned true
func (q *QueryStrings) Filter(filter func(name string) bool) {
	var filtered []*QueryString
	for _, v := range q.Items {
		if !filter(v.Key) {
			continue
		}
		filtered = append(filtered, v)
	}
	q.Items = filtered
}

type SortMode string

const (
	SortAsc  SortMode = "asc"
	SortDesc SortMode = "desc"
)

// Sort sorts key with sort mode (asc or desc)
func (q *QueryStrings) Sort(mode SortMode) {
	sort.Slice(q.Items, func(i, j int) bool {
		v := q.Items[i].Key > q.Items[j].Key
		if mode == SortAsc {
			return !v
		}
		return v
	})
}

// String() implements fmt.Stringer interface, return formatted raw query string
func (q *QueryStrings) String() string {
	var buf strings.Builder
	for i, v := range q.Items {
		key := q.Items[i].Key
		if v.Value == nil {
			buf.WriteString(key)
		} else {
			for j := range v.Value {
				buf.WriteString(key)
				buf.WriteString("=")
				buf.WriteString(url.QueryEscape(v.Value[j]))
				if j != len(v.Value)-1 {
					buf.WriteString("&")
				}
			}
		}
		if i != len(q.Items)-1 {
			buf.WriteString("&")
		}
	}
	var sign string
	if buf.Len() > 0 {
		sign = "?"
	}
	return q.Prefix + sign + buf.String()
}
