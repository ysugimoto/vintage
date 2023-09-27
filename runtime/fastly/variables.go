package fastly

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/netip"
	"strconv"
	"strings"
	"time"

	"github.com/fastly/compute-sdk-go/fsthttp"
	"github.com/pkg/errors"
)

// Used for client.identity
func (r *Runtime) ClientIdentity() string {
	if r.clientIdentity == "" {
		// default as client.ip
		r.clientIdentity = r.ClientIP().String()
	}
	return r.clientIdentity
}

// Used for client.ip
func (r *Runtime) ClientIP() net.IP {
	idx := strings.LastIndex(r.Request.RemoteAddr, ":")
	if idx == -1 {
		return net.ParseIP(r.Request.RemoteAddr)
	}
	return net.ParseIP(r.Request.RemoteAddr[:idx])
}

// Used for req.body
func (r *Runtime) RequestBody() (string, error) {
	switch r.Request.Method {
	case fsthttp.MethodPatch, fsthttp.MethodPost, fsthttp.MethodPut:
		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Request.Body); err != nil {
			return "", errors.WithStack(
				fmt.Errorf("Failed to read request body: %s", err),
			)
		}
		r.Request.Body = io.NopCloser(bytes.NewReader(b.Bytes()))
		// size is limited to 8KB
		if len(b.Bytes()) > 1024*8 {
			return "", errors.WithStack(
				fmt.Errorf("Request body is limited by 8KB, overflow"),
			)
		}
		return b.String(), nil
	default:
		return "", nil
	}
}

// Used for req.body.base64
func (r *Runtime) RequestBodyBase64() (string, error) {
	body, err := r.RequestBody()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return base64.StdEncoding.EncodeToString([]byte(body)), nil
}

// Used for req.digest
func (r *Runtime) RequestDigest() string {
	if r.RequestHash == "" {
		return strings.Repeat("0", 64)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(r.RequestHash)))
}

// Used for req.url and bereq.url
func (r *Runtime) RequestURL(req *fsthttp.Request) string {
	u := req.URL.Path
	if v := req.URL.RawQuery; v != "" {
		u += "?" + v
	}
	if v := req.URL.RawFragment; v != "" {
		u += "#" + v
	}
	return u
}

// Used for bereq.body_bytes_written
func (r *Runtime) BackendBodyBytesWritten() (int64, error) {
	if r.BackendRequest == nil {
		return 0, nil
	}

	switch r.BackendRequest.Method {
	case fsthttp.MethodPatch, fsthttp.MethodPost, fsthttp.MethodPut:
		var b bytes.Buffer
		if _, err := b.ReadFrom(r.BackendRequest.Body); err != nil {
			return 0, errors.WithStack(
				fmt.Errorf("Failed to read backend request body: %s", err),
			)
		}
		r.BackendRequest.Body = io.NopCloser(bytes.NewReader(b.Bytes()))
		return int64(len(b.Bytes())), nil
	default:
		return 0, nil
	}
}

// Used for bereq.header_bytes_written
func (r *Runtime) BackendHeaderBytesWritten() int64 {
	if r.BackendRequest == nil {
		return 0
	}

	var size int
	for key, val := range r.BackendRequest.Header {
		size += len(key)
		for i := range val {
			size += len(val[i]) + 1
		}
		size--
	}
	return int64(size)
}

// Used for bereq.bytes_written
func (r *Runtime) BackendBytesWritten() (int64, error) {
	if r.BackendRequest == nil {
		return 0, nil
	}

	var size int64
	// Request line
	size += int64(len(fmt.Sprintf(
		"%s %s %s",
		r.BackendRequest.Method,
		r.BackendRequest.URL.String(),
		r.BackendRequest.Proto,
	)))
	// Headers
	size += r.BackendHeaderBytesWritten() + 1
	// Body
	body, err := r.BackendBodyBytesWritten()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	size += body
	return size, nil
}

// Used for obj.age
func (r *Runtime) ObjectAge() time.Duration {
	if r.BackendResponse == nil {
		return 0
	}
	dur := "0"
	if v := r.BackendResponse.Header.Get("Age"); v != "" {
		dur = v
	}
	d, err := time.ParseDuration(dur + "s")
	if err != nil {
		return 0
	}
	return d
}

// Fastly follows its own cache freshness rules
// see: https://developer.fastly.com/learning/concepts/cache-freshness/
var cacheableStatusCodes = []int{200, 203, 300, 301, 302, 404, 410}

// Used for obj.cachable
func (r *Runtime) ObjectCacheable() bool {
	if r.BackendResponse == nil {
		return false
	}
	for i := range cacheableStatusCodes {
		if r.BackendResponse.StatusCode == cacheableStatusCodes[i] {
			return true
		}
	}
	return false
}

// Used for obj.hits
func (r *Runtime) ObjectHits() int64 {
	if r.BackendResponse == nil {
		return 0
	}
	hits := "0"
	if v := r.BackendResponse.Header.Get("X-Cache-Hits"); v != "" {
		hits = v
	}
	h, err := strconv.ParseInt(hits, 10, 64)
	if err != nil {
		return 0
	}
	return h
}

// Used for req.is_ipv6
func (r *Runtime) IsIpv6() bool {
	parsed, err := netip.ParseAddr(r.Request.RemoteAddr)
	if err != nil {
		return false
	}
	return parsed.Is6()
}

// Used for req.body_bytes_read
func (r *Runtime) RequestBodyBytesRead() (int64, error) {
	switch r.Request.Method {
	case fsthttp.MethodPatch, fsthttp.MethodPost, fsthttp.MethodPut:
		var b bytes.Buffer
		if _, err := b.ReadFrom(r.Request.Body); err != nil {
			return 0, errors.WithStack(
				fmt.Errorf("Failed to read backend request body: %s", err),
			)
		}
		r.Request.Body = io.NopCloser(bytes.NewReader(b.Bytes()))
		return int64(len(b.Bytes())), nil
	default:
		return 0, nil
	}
}

// Used for req.bytes_read
func (r *Runtime) RequestBytesRead() (int64, error) {
	var size int64
	// Request line
	size += int64(len(fmt.Sprintf(
		"%s %s %s",
		r.Request.Method,
		r.Request.URL.String(),
		r.Request.Proto,
	)))
	// Headers
	for key, val := range r.Request.Header {
		size += int64(len(key))
		for i := range val {
			size += int64(len(val[i])) + 1
		}
		size--
	}
	// Body
	body, err := r.RequestBodyBytesRead()
	if err != nil {
		return 0, errors.WithStack(err)
	}
	size += body
	return size, nil
}

// Used for resp.response and beresp.response
func (r *Runtime) ResponseBody(resp *fsthttp.Response) (string, error) {
	if resp == nil {
		return "", nil
	}
	var b bytes.Buffer
	if _, err := b.ReadFrom(resp.Body); err != nil {
		return "", errors.WithStack(
			fmt.Errorf("Failed to read response body: %s", err),
		)
	}
	// rewind
	resp.Body = io.NopCloser(bytes.NewReader(b.Bytes()))
	return b.String(), nil
}
