package vintage

import (
	"net"
	"strconv"
	"time"
)

func ToString[T Primitive](v T) string {
	switch t := any(v).(type) {
	case string:
		return t
	case int64:
		return strconv.FormatInt(t, 10)
	case float64:
		return strconv.FormatFloat(t, 'f', 3, 64)
	case bool:
		if t {
			return "1"
		}
		return "0"
	case net.IP:
		return t.String()
	case time.Duration:
		return strconv.FormatFloat(float64(t.Milliseconds())/1000, 'f', 3, 64)
	case time.Time:
		return t.Format(httpTime)
	case *Backend:
		return t.Backend()
	case *Acl:
		return t.Name
	default:
		return ""
	}
}

func ToBool[T Primitive](v T) bool {
	switch t := any(v).(type) {
	case string:
		return t != ""
	case bool:
		return t
	}
	return false
}
