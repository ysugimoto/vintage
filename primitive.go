package vintage

import (
	"net"
	"strconv"
	"time"
)

type Value interface {
	// Need String() method for explicit type conversion in expression
	String() string
	// Need Bool() method for explicit type conversion in conditional expression
	Bool() bool
}

type String string

func (v String) String() string { return string(v) }
func (v String) Bool() bool     { return string(v) != "" }

type Integer int64

func (v Integer) String() string { return strconv.FormatInt(int64(v), 10) }
func (v Integer) Bool() bool     { return false }

type Float float64

func (v Float) String() string { return strconv.FormatFloat(float64(v), 'f', 3, 64) }
func (v Float) Bool() bool     { return false }

type Bool bool

func (v Bool) String() string {
	if bool(v) {
		return "1"
	}
	return "0"
}
func (v Bool) Bool() bool { return bool(v) }

type ip net.IP

func (v ip) String() string { return v.String() }
func (v ip) Bool() bool     { return false }

type RTime time.Duration

func (v RTime) String() string {
	return strconv.FormatFloat(float64(time.Duration(v).Milliseconds())/1000, 'f', 3, 64)
}
func (v RTime) Bool() bool { return false }

type Time time.Time

const httpTime = "Mon, 02 Jan 2006 15:04:05 GMT"

func (v Time) String() string {
	return time.Time(v).Format(httpTime)
}
func (v Time) Bool() bool { return false }
