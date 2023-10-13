package variable

import "regexp"

var (
	RequestHttpHeaderRegex         = regexp.MustCompile(`^req\.http\.(.+)`)
	BackendRequestHttpHeaderRegex  = regexp.MustCompile(`^bereq\.http\.(.+)`)
	BackendResponseHttpHeaderRegex = regexp.MustCompile(`^beresp\.http\.(.+)`)
	ResponseHttpHeaderRegex        = regexp.MustCompile(`^resp\.http\.(.+)`)
	ObjectHttpHeaderRegex          = regexp.MustCompile(`^obj\.http\.(.+)`)

	// Currently RateCounter is unsupported
	// RateCounterRegex            = regexp.MustCompile(`ratecounter\.([^\.]+)\.(.+)`)
)
