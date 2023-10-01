package core

import (
	"fmt"
	"log"
	"net/textproto"
	"strings"
)

type Header struct {
	MH textproto.MIMEHeader
}

func NewHeader(h map[string][]string) *Header {
	return &Header{
		MH: textproto.MIMEHeader(h),
	}
}

func (h *Header) Set(key, value string) {
	if !strings.Contains(key, ":") {
		h.MH.Set(key, value)
		return
	}
	spl := strings.SplitN(key, ":", 2)
	if strings.EqualFold(spl[0], "cookie") {
		s := fmt.Sprintf("%s=%s", sanitizeCookieName(spl[1]), sanitizeCookieValue(value))
		if c := h.MH.Get("Cookie"); c != "" {
			h.MH.Set("Cookie", c+"; "+s)
		} else {
			h.MH.Set("Cookie", s)
		}
		return
	}
	h.MH.Add(spl[0], fmt.Sprintf("%s=%s", spl[1], value))
}

func (h *Header) Add(key, value string) {
	h.MH.Add(key, value)
}

func (h *Header) Get(key string) string {
	if !strings.Contains(key, ":") {
		return strings.Join(h.MH.Values(key), ", ")
	}
	spl := strings.SplitN(key, ":", 2)
	if strings.EqualFold(spl[0], "cookie") {
		return readCookie(h.MH["Cookie"], spl[1])
	}

	for _, hv := range h.MH.Values(spl[0]) {
		kv := strings.SplitN(hv, "=", 2)
		if kv[0] == spl[1] {
			return kv[1]
		}
	}
	return ""
}

func (h *Header) Unset(key string) {
	if !strings.Contains(key, ":") {
		h.MH.Del(key)
		return
	}

	// Header name contains ":" character, then filter value by key
	spl := strings.SplitN(key, ":", 2)
	// Request header can modify cookie, then we need to unset value from Cookie pointer
	if strings.EqualFold(spl[0], "cookie") {
		h.removeCookieByName(spl[1])
		return
	}

	var filtered []string
	for _, hv := range h.MH.Values(spl[0]) {
		kv := strings.SplitN(hv, "=", 2)
		if kv[0] == spl[1] {
			continue
		}
		filtered = append(filtered, hv)
	}

	h.MH.Del(spl[0])
	if len(filtered) > 0 {
		for i := range filtered {
			h.MH.Add(spl[0], filtered[i])
		}
	}
}

func (h *Header) removeCookieByName(cookieName string) {
	lines := h.MH["Cookie"]
	if len(lines) == 0 {
		return
	}

	var filtered []string
	for _, line := range lines {
		line = textproto.TrimString(line)

		var sub []string
		var part string
		for len(line) > 0 { // continue since we have rest
			part, line, _ = strings.Cut(line, ";")
			trimmedPart := textproto.TrimString(part)
			if trimmedPart == "" {
				continue
			}
			name, _, _ := strings.Cut(trimmedPart, "=")
			name = textproto.TrimString(name)
			if name == cookieName {
				continue
			}
			sub = append(sub, part)
		}

		if len(sub) > 0 {
			filtered = append(filtered, strings.Join(sub, ";"))
		}
	}
	if len(filtered) > 0 {
		h.MH["Cookie"] = filtered
	} else {
		h.MH.Del("Cookie")
	}
}

func readCookie(lines []string, filter string) string {
	if len(lines) == 0 {
		return ""
	}

	for _, line := range lines {
		line = textproto.TrimString(line)

		var part string
		for len(line) > 0 { // continue since we have rest
			part, line, _ = strings.Cut(line, ";")
			part = textproto.TrimString(part)
			if part == "" {
				continue
			}
			name, val, _ := strings.Cut(part, "=")
			name = textproto.TrimString(name)
			if !isCookieNameValid(name) {
				continue
			}
			if filter != "" && filter != name {
				continue
			}
			val, ok := parseCookieValue(val, true)
			if !ok {
				continue
			}
			return val
		}
	}
	return ""
}

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func sanitizeCookieName(n string) string {
	return cookieNameSanitizer.Replace(n)
}

// sanitizeCookieValue produces a suitable cookie-value from v.
// https://tools.ietf.org/html/rfc6265#section-4.1.1
//
//	cookie-value      = *cookie-octet / ( DQUOTE *cookie-octet DQUOTE )
//	cookie-octet      = %x21 / %x23-2B / %x2D-3A / %x3C-5B / %x5D-7E
//	          ; US-ASCII characters excluding CTLs,
//	          ; whitespace DQUOTE, comma, semicolon,
//	          ; and backslash
//
// We loosen this as spaces and commas are common in cookie values
// but we produce a quoted cookie-value if and only if v contains
// commas or spaces.
// See https://golang.org/issue/7243 for the discussion.
func sanitizeCookieValue(v string) string {
	v = sanitizeOrWarn("Cookie.Value", validCookieValueByte, v)
	if len(v) == 0 {
		return v
	}
	if strings.ContainsAny(v, " ,") {
		return `"` + v + `"`
	}
	return v
}

// path-av           = "Path=" path-value
// path-value        = <any CHAR except CTLs or ";">
func sanitizeCookiePath(v string) string {
	return sanitizeOrWarn("Cookie.Path", validCookiePathByte, v)
}

func validCookiePathByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != ';'
}

func sanitizeOrWarn(fieldName string, valid func(byte) bool, v string) string {
	ok := true
	for i := 0; i < len(v); i++ {
		if valid(v[i]) {
			continue
		}
		log.Printf("net/http: invalid byte %q in %s; dropping invalid bytes", v[i], fieldName)
		ok = false
		break
	}
	if ok {
		return v
	}
	buf := make([]byte, 0, len(v))
	for i := 0; i < len(v); i++ {
		if b := v[i]; valid(b) {
			buf = append(buf, b)
		}
	}
	return string(buf)
}
func parseCookieValue(raw string, allowDoubleQuote bool) (string, bool) {
	// Strip the quotes, if present.
	if allowDoubleQuote && len(raw) > 1 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		raw = raw[1 : len(raw)-1]
	}
	for i := 0; i < len(raw); i++ {
		if !validCookieValueByte(raw[i]) {
			return "", false
		}
	}
	return raw, true
}

func validCookieValueByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != '"' && b != ';' && b != '\\'
}

func isCookieNameValid(raw string) bool {
	if raw == "" {
		return false
	}
	return strings.IndexFunc(raw, isNotToken) < 0
}
func isNotToken(r rune) bool {
	return !isTokenRune(r)
}

var isTokenTable = [127]bool{
	'!':  true,
	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,
	'*':  true,
	'+':  true,
	'-':  true,
	'.':  true,
	'0':  true,
	'1':  true,
	'2':  true,
	'3':  true,
	'4':  true,
	'5':  true,
	'6':  true,
	'7':  true,
	'8':  true,
	'9':  true,
	'A':  true,
	'B':  true,
	'C':  true,
	'D':  true,
	'E':  true,
	'F':  true,
	'G':  true,
	'H':  true,
	'I':  true,
	'J':  true,
	'K':  true,
	'L':  true,
	'M':  true,
	'N':  true,
	'O':  true,
	'P':  true,
	'Q':  true,
	'R':  true,
	'S':  true,
	'T':  true,
	'U':  true,
	'W':  true,
	'V':  true,
	'X':  true,
	'Y':  true,
	'Z':  true,
	'^':  true,
	'_':  true,
	'`':  true,
	'a':  true,
	'b':  true,
	'c':  true,
	'd':  true,
	'e':  true,
	'f':  true,
	'g':  true,
	'h':  true,
	'i':  true,
	'j':  true,
	'k':  true,
	'l':  true,
	'm':  true,
	'n':  true,
	'o':  true,
	'p':  true,
	'q':  true,
	'r':  true,
	's':  true,
	't':  true,
	'u':  true,
	'v':  true,
	'w':  true,
	'x':  true,
	'y':  true,
	'z':  true,
	'|':  true,
	'~':  true,
}

func isTokenRune(r rune) bool {
	i := int(r)
	return i < len(isTokenTable) && isTokenTable[i]
}
