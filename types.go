package vintage

import (
	"net"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

type DirectorType string

const (
	Random   DirectorType = "random"
	Fallback DirectorType = "fallback"
	Hash     DirectorType = "hash"
	Client   DirectorType = "client"
	CHash    DirectorType = "chash"
)

type Primitive interface {
	string | int64 | float64 | bool | net.IP | time.Duration | time.Time | *Backend | *Acl
}

// Alias of time.Duration for VCL RTIME value
const (
	Millisecond time.Duration = time.Millisecond
	Second      time.Duration = time.Second
	Minute      time.Duration = time.Minute
	Hour        time.Duration = time.Hour
	Day         time.Duration = 24 * time.Hour
	Year        time.Duration = 24 * 365 * time.Hour
)

// Temporal value of local IP
var LocalHost = net.IPv4(127, 0, 0, 1)

type State string

const (
	NONE          State = ""
	LOOKUP        State = "lookup"
	PASS          State = "pass"
	HASH          State = "hash"
	ERROR         State = "error"
	RESTART       State = "restart"
	DELIVER       State = "deliver"
	FETCH         State = "fetch"
	DELIVER_STALE State = "deliver_stale"
	LOG           State = "log"
)

func Error(err error) (State, error) {
	return NONE, errors.WithStack(err)
}

// RawHeader represents underlying type of http.Header.
// To abstract HTTP context and WASM runtime could not import net/http package,
// Our runtime would use as Golang underlying type
type RawHeader map[string][]string

// RegexpMatchedGroup represents regexp matched group values
// which is stored when "~" or "!~" operator is used
type RegexpMatchedGroup []string

func (re RegexpMatchedGroup) Match(pattern, subject string) (bool, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, errors.WithStack(err)
	}
	matches := r.FindStringSubmatch(subject)
	if matches == nil {
		return false, nil
	}
	re = matches
	return true, nil
}

func (re RegexpMatchedGroup) At(index int) string {
	if index >= len(re)-1 || index < 0 {
		return ""
	}
	return re[index]
}

func (re RegexpMatchedGroup) Release() {
	re = []string{}
}
