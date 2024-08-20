package vintage

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

const httpTime = "Mon, 02 Jan 2006 15:04:05 GMT"

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
		return t.Name
	case *Acl:
		return t.Name
	case *Table:
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

// RegexpMatch is function that wraps regular expression matching
// Returns matches result and capture groups
func RegexpMatch(pattern, subject string) (bool, RegexpMatchedGroup, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, nil, errors.WithStack(err)
	}
	matches := r.FindStringSubmatch(subject)
	if matches == nil {
		return false, nil, nil
	}
	return true, RegexpMatchedGroup(matches), nil
}

// We heavily respect original implementation of github.com/rs/xid.
// This function picks functions only we need
// because original imports some of unsupported package on the WASM environment.
var objectCounter uint32 = func() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("xid: cannot generate random number: %v;", err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}()

var xidEncoding = "0123456789abcdefghijklmnopqrstuv"

func GenerateXid() string {
	// WASM runtime will not have any machine id so id is always generated as random bytes
	// This line comes from readPlatformMachineID on https://github.com/rs/xid/blob/master/id.go#L110
	machineId := make([]byte, 3)
	rand.Reader.Read(machineId) // nolint:errcheck

	pid := os.Getpid()

	now := time.Now()
	var id [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(id[:], uint32(now.Unix()))
	// Machine ID, 3 bytes
	id[4] = machineId[0]
	id[5] = machineId[1]
	id[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectCounter, 1)
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)

	dst := make([]byte, 20) // encodedLen is 20
	_ = dst[19]
	_ = id[11]

	dst[19] = xidEncoding[(id[11]<<4)&0x1F]
	dst[18] = xidEncoding[(id[11]>>1)&0x1F]
	dst[17] = xidEncoding[(id[11]>>6)|(id[10]<<2)&0x1F]
	dst[16] = xidEncoding[id[10]>>3]
	dst[15] = xidEncoding[id[9]&0x1F]
	dst[14] = xidEncoding[(id[9]>>5)|(id[8]<<3)&0x1F]
	dst[13] = xidEncoding[(id[8]>>2)&0x1F]
	dst[12] = xidEncoding[id[8]>>7|(id[7]<<1)&0x1F]
	dst[11] = xidEncoding[(id[7]>>4)|(id[6]<<4)&0x1F]
	dst[10] = xidEncoding[(id[6]>>1)&0x1F]
	dst[9] = xidEncoding[(id[6]>>6)|(id[5]<<2)&0x1F]
	dst[8] = xidEncoding[id[5]>>3]
	dst[7] = xidEncoding[id[4]&0x1F]
	dst[6] = xidEncoding[id[4]>>5|(id[3]<<3)&0x1F]
	dst[5] = xidEncoding[(id[3]>>2)&0x1F]
	dst[4] = xidEncoding[id[3]>>7|(id[2]<<1)&0x1F]
	dst[3] = xidEncoding[(id[2]>>4)|(id[1]<<4)&0x1F]
	dst[2] = xidEncoding[(id[1]>>1)&0x1F]
	dst[1] = xidEncoding[(id[1]>>6)|(id[0]<<2)&0x1F]
	dst[0] = xidEncoding[id[0]>>3]

	return string(dst)
}

// Fastly follows its own cache freshness rules
// see: https://developer.fastly.com/learning/concepts/cache-freshness/
var unCacheableStatusCodes = []int{200, 203, 300, 301, 302, 404, 410}

func IsCacheableStatusCode(statusCode int) bool {
	for _, v := range unCacheableStatusCodes {
		if v == statusCode {
			return true
		}
	}
	return false
}
