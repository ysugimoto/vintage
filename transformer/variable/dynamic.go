package variable

import "regexp"

var (
	RequestHttpHeaderRegex         = regexp.MustCompile(`^req\.http\.(.+)`)
	BackendRequestHttpHeaderRegex  = regexp.MustCompile(`^bereq\.http\.(.+)`)
	BackendResponseHttpHeaderRegex = regexp.MustCompile(`^beresp\.http\.(.+)`)
	ResponseHttpHeaderRegex        = regexp.MustCompile(`^resp\.http\.(.+)`)
	ObjectHttpHeaderRegex          = regexp.MustCompile(`^obj\.http\.(.+)`)
	// RegexMatchedRegex              = regexp.MustCompile(`re\.group\.([0-9]+)`)
	// RateCounterRegex               = regexp.MustCompile(`ratecounter\.([^\.]+)\.(.+)`)
)
