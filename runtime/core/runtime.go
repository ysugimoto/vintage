package core

import (
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

type Subroutine[T EdgeRuntime] func(ctx T) (vintage.State, error)

type Resource[T EdgeRuntime] func(c *Runtime[T])

func BackendResource[T EdgeRuntime](name string, v *vintage.Backend) Resource[T] {
	return func(c *Runtime[T]) {
		c.Backends[name] = v
		if v.IsDefault {
			c.Backend = v
		}
	}
}
func AclResource[T EdgeRuntime](name string, v *vintage.Acl) Resource[T] {
	return func(c *Runtime[T]) {
		c.Acls[name] = v
	}
}
func TableResource[T EdgeRuntime](name string, v *vintage.Table) Resource[T] {
	return func(c *Runtime[T]) {
		c.Tables[name] = v
	}
}
func SubroutineResource[T EdgeRuntime](name string, v Subroutine[T]) Resource[T] {
	return func(c *Runtime[T]) {
		c.Subroutines[name] = v
	}
}

func (c *Runtime[T]) Register(resources ...Resource[T]) {
	for i := range resources {
		resources[i](c)
	}
}

type Runtime[T EdgeRuntime] struct {
	Backend            *vintage.Backend
	RequestHeader      *Header
	ResponseHeader     *Header
	RequestStartTime   time.Time
	RequestEndTime     time.Time
	Restarts           int
	RequestHash        string
	RequestHeaderBytes int64

	// We should implement User-Agent related matcher by ourselves
	UserAgent *UserAgent
	// Temporal WAF related variables
	Waf *Waf

	// Declaration stacks
	Backends    map[string]*vintage.Backend
	Acls        map[string]*vintage.Acl
	Tables      map[string]*vintage.Table
	Subroutines map[string]Subroutine[T]

	// Properties that will be assigned in the process
	MaxStaleIfError                     time.Duration
	MaxStaleWhileRevalidate             time.Duration
	ClientSocketCongestionAlgorithm     string
	EsiAllowInsideData                  bool
	EnableESI                           bool
	ESILevel                            int64
	IsLocallyGenerated                  bool
	TimeToFirstByte                     time.Duration
	BackendResponseBrotli               bool
	BackendResponseCachable             bool
	BackendResponseDoESI                bool
	BackendResponseDoStream             bool
	BackendResponseGrace                bool
	BackendResponseGzip                 bool
	BackendResponseHipaa                bool
	BackendResponsePCI                  bool
	BackendResponseStaleIfError         time.Duration
	BackendResponseStaleWhileRevalidate time.Duration
	BackendResponseTTL                  time.Duration
	ObjectStaleIfError                  time.Duration
	ObjectStaleWhileRevalidate          time.Duration
	ObjectTTL                           time.Duration
	EnableRangeOnPass                   bool
	EnableSegmentedCaching              bool
	HashAlwaysMiss                      bool
	HashIgnoreBusy                      bool
	ResponseBytesWritten                int64
	ResponseBodyBytesWritten            int64
	ResponseHeaderBytesWritten          int64
}

func NewRuntime[T EdgeRuntime](w, r map[string][]string, resources ...Resource[T]) *Runtime[T] {
	c := &Runtime[T]{
		RequestHeader:    NewHeader(textproto.MIMEHeader(r)),
		ResponseHeader:   NewHeader(textproto.MIMEHeader(w)),
		RequestStartTime: time.Now(),

		Backends:    make(map[string]*vintage.Backend),
		Acls:        make(map[string]*vintage.Acl),
		Tables:      make(map[string]*vintage.Table),
		Subroutines: make(map[string]Subroutine[T]),

		// Default variable values, (explicitly write default value even zero value)
		ClientSocketCongestionAlgorithm:     "cubic",
		EsiAllowInsideData:                  false,
		EnableESI:                           false,
		ESILevel:                            0,
		IsLocallyGenerated:                  false,
		BackendResponseBrotli:               false,
		BackendResponseCachable:             false,
		BackendResponseDoESI:                false,
		BackendResponseDoStream:             false,
		BackendResponseGrace:                false,
		BackendResponseGzip:                 false,
		BackendResponseHipaa:                false,
		BackendResponsePCI:                  false,
		BackendResponseStaleIfError:         0,
		BackendResponseStaleWhileRevalidate: 0,
		ObjectStaleIfError:                  0,
		ObjectStaleWhileRevalidate:          0,
		ObjectTTL:                           0,
		EnableRangeOnPass:                   false,
		EnableSegmentedCaching:              false,
		HashAlwaysMiss:                      false,
		HashIgnoreBusy:                      false,
		ResponseBytesWritten:                0,
		ResponseBodyBytesWritten:            0,
		ResponseHeaderBytesWritten:          0,
	}
	for i := range resources {
		resources[i](c)
	}
	c.UserAgent = NewUserAgent(c.RequestHeader.Get("User-Agent"))
	c.Waf = &Waf{}
	c.RequestHeaderBytes = c.factoryInitRequestHeaderBytes(r)
	return c
}

func (c *Runtime[T]) factoryInitRequestHeaderBytes(hs map[string][]string) int64 {
	var size int64
	for key, val := range hs {
		size += int64(len(key))
		for i := range val {
			size += int64(len(val[i])) + 1 // add 1 byte of comma
		}
		size -= 1 // minus last comma
	}
	return size
}

func (c *Runtime[T]) RequestRangeHeader(mode string) (int64, error) {
	spl := strings.SplitN(c.RequestHeader.Get("Range"), "-", 2)
	if len(spl) != 2 {
		return 0, nil
	}
	var val string
	switch mode {
	case "low":
		val = spl[0]
	case "high":
		val = spl[1]
	default:
		return 0, errors.New("Unexpcted range header mode, accept only low or high")
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return i, nil
}
