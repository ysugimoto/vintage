package lib

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// Fastly describes that the following characters are valid for header name:
// ! # $ % & ' * + - . 0-9 A-Z ^ _ ` a-z | ~
// And limited to 126 character length.
// @see https://developer.fastly.com/reference/vcl/functions/headers/header-set/#header-names
// And also accepts ":" character to treat as object-like header value
var validHeaderCharacters = regexp.MustCompile("^[!#$%&'*+-.0-9A-Z^_`a-z|~:]{1,126}$")

func IsValidHeader(name string) bool {
	return validHeaderCharacters.MatchString(name)
}

// Fastly proctects some headers.
// The proctcted headers cannot modify (set, unset) in VCL.
// We define header name as lower case to ensure easily.
// see: https://developer.fastly.com/reference/http/http-headers/
var protectedHeaders = map[string]struct{}{
	"proxy-authenticate":  {},
	"proxy-authotization": {},
	"content-length":      {},
	"content-range":       {},
	"te":                  {},
	"trailer":             {},
	"expect":              {},
	"transfer-encoding":   {},
	"upgrade":             {},
	"fastly-ff":           {},
}

func CheckProtectedHeader(name string) error {
	if _, ok := protectedHeaders[strings.ToLower(name)]; !ok {
		return nil
	}
	return errors.WithStack(
		fmt.Errorf("Could not modify %s header. This header is protected by Fastly", name),
	)
}
