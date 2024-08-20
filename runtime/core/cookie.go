package core

import (
	"log"
	"net/textproto"
	"strings"
)

// SameSite allows a server to define a cookie attribute making it impossible for
// the browser to send this cookie along with cross-site requests. The main
// goal is to mitigate the risk of cross-origin information leakage, and provide
// some protection against cross-site request forgery attacks.
//
// See https://tools.ietf.org/html/draft-ietf-httpbis-cookie-same-site-00 for details.
type SameSite int

const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)

// deleteSetCookie deletes matched Set-Cookie header.
// This function high respects net/http package's readSetCookies function.
func deleteSetCookie(h *Header, cookieName string) bool {
	var isDeleted bool
	cookieCount := len(h.MH["Set-Cookie"])
	if cookieCount == 0 {
		return isDeleted
	}

	var filtered []string
	for _, line := range h.MH["Set-Cookie"] {
		parts := strings.Split(textproto.TrimString(line), ";")
		if len(parts) == 1 && parts[0] == "" {
			continue
		}
		parts[0] = textproto.TrimString(parts[0])
		name, _, ok := strings.Cut(parts[0], "=")
		if !ok {
			continue
		}
		name = textproto.TrimString(name)
		if !isCookieNameValid(name) {
			continue
		}
		if name == cookieName {
			isDeleted = true
			continue
		}
		filtered = append(filtered, line)
	}

	if len(filtered) > 0 {
		h.MH["Set-Cookie"] = filtered
	} else {
		h.MH.Del("Set-Cookie")
	}
	return isDeleted
}

// getSetCookie retrieves matched Set-Cookie header value.
// This function high respects net/http package's readSetCookies function.
func getSetCookie(h *Header, cookieName string) string {
	cookieCount := len(h.MH["Set-Cookie"])
	if cookieCount == 0 {
		return ""
	}
	cookies := map[string]string{}
	for _, line := range h.MH["Set-Cookie"] {
		parts := strings.Split(textproto.TrimString(line), ";")
		if len(parts) == 1 && parts[0] == "" {
			continue
		}
		parts[0] = textproto.TrimString(parts[0])
		name, value, ok := strings.Cut(parts[0], "=")
		if !ok {
			continue
		}
		name = textproto.TrimString(name)
		if !isCookieNameValid(name) {
			continue
		}
		value, ok = parseCookieValue(value, true)
		if !ok {
			continue
		}
		cookies[name] = value
	}

	if v, ok := cookies[cookieName]; ok {
		return v
	}
	return ""
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
		for line != "" { // continue since we have rest
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
		for line != "" { // continue since we have rest
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
	if v == "" {
		return v
	}
	if strings.ContainsAny(v, " ,") {
		return `"` + v + `"`
	}
	return v
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
