package vintage

import (
	"net"
	"time"
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
