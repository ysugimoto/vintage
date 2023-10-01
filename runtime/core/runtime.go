package core

import (
	"crypto/sha256"
	"fmt"
	"net/textproto"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

type Runtime[T EdgeRuntime] struct {
	Backend               *vintage.Backend
	RequestHeader         *Header
	BackendRequestHeader  *Header
	BackendResponseHeader *Header
	ResponseHeader        *Header
	RequestStartTime      time.Time
	RequestEndTime        time.Time
	Restarts              int
	RequestHash           string
	RequestHeaderBytes    int64

	// We should implement User-Agent related matcher by ourselves
	UserAgent *UserAgent
	// Temporal WAF related variables
	Waf *Waf

	// Declaration stacks
	Backends         map[string]*vintage.Backend
	Acls             map[string]*vintage.Acl
	Tables           map[string]*vintage.Table
	Subroutines      map[string]Subroutine[T]
	LoggingEndpoints map[string]*vintage.LoggingEndpoint

	// Properties that will be assigned in the process
	OriginalHost                        string
	MaxStaleIfError                     time.Duration
	MaxStaleWhileRevalidate             time.Duration
	ClientSocketCongestionAlgorithm     string
	ClientSocketCWND                    int64
	ClientSocketPace                    int64
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
	ObjectStatus                        int64
	ObjectResponse                      string
	EnableRangeOnPass                   bool
	EnableSegmentedCaching              bool
	HashAlwaysMiss                      bool
	HashIgnoreBusy                      bool
	ResponseBytesWritten                int64
	ResponseBodyBytesWritten            int64
	ResponseHeaderBytesWritten          int64
	GeoIpOverride                       string
	GeoIpUseXForwardedFor               bool
	SegmentedCachingBlockSize           int64
	SaintMode                           bool
	ResponseStale                       bool
	ResponseStaleIsError                bool
	ResponseStaleIsRevalidating         bool
	ResponseCompleted                   bool

	// Following fields are set via builtin function
	FastlyError               string
	DisableCompressionHeaders []string
	PushResources             []string
	H3AltSvc                  bool
}

func NewRuntime[T EdgeRuntime](r map[string][]string, resources ...Resource[T]) *Runtime[T] {
	c := &Runtime[T]{
		RequestHeader:    NewHeader(textproto.MIMEHeader(r)),
		RequestStartTime: time.Now(),

		Backends:         make(map[string]*vintage.Backend),
		Acls:             make(map[string]*vintage.Acl),
		Tables:           make(map[string]*vintage.Table),
		Subroutines:      make(map[string]Subroutine[T]),
		LoggingEndpoints: make(map[string]*vintage.LoggingEndpoint),

		// Default variable values, (explicitly write default value even zero value)
		ClientSocketCongestionAlgorithm:     "cubic",
		ClientSocketCWND:                    60,
		ClientSocketPace:                    131072, // 128KiB
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
		BackendResponseStaleIfError:         time.Duration(0),
		BackendResponseStaleWhileRevalidate: time.Duration(0),
		BackendResponseTTL:                  time.Duration(0),
		ObjectStaleIfError:                  time.Duration(0),
		ObjectStaleWhileRevalidate:          time.Duration(0),
		ObjectTTL:                           time.Duration(0),
		EnableRangeOnPass:                   false,
		EnableSegmentedCaching:              false,
		HashAlwaysMiss:                      false,
		HashIgnoreBusy:                      false,
		ResponseBytesWritten:                0,
		ResponseBodyBytesWritten:            0,
		ResponseHeaderBytesWritten:          0,
		GeoIpOverride:                       "",
		GeoIpUseXForwardedFor:               false,
		SegmentedCachingBlockSize:           1,
		SaintMode:                           false,
		ResponseStale:                       false,
		ResponseStaleIsError:                false,
		ResponseStaleIsRevalidating:         false,
		ResponseCompleted:                   false,
	}
	for i := range resources {
		resources[i](c)
	}
	c.UserAgent = NewUserAgent(c.RequestHeader.Get("User-Agent"))
	c.Waf = &Waf{}
	c.RequestHeaderBytes = c.factoryInitRequestHeaderBytes(r)
	return c
}

func (c *Runtime[T]) Cleanup() {
	// Hook point to release some resources
}

func (c *Runtime[T]) RegexpMatch(src, dst string) (bool, error) {
	re, err := regexp.Compile(src)
	if err != nil {
		return false, errors.WithStack(err)
	}
	matches := re.FindStringSubmatch(dst)
	if matches == nil {
		return false, nil
	}
	return true, nil

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

// Used for req.digest
func (c *Runtime[T]) RequestDigest() string {
	if c.RequestHash == "" {
		return strings.Repeat("0", 64)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(c.RequestHash)))
}
