package core

import (
	"fmt"
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

// Used for setcookie.delete_by_name()
func (h *Header) DeleteSetCookie(name string) bool {
	return deleteSetCookie(h, name)
}

// Used for setcookie.get_value_by_name()
func (h *Header) GetSetCookie(name string) string {
	return getSetCookie(h, name)
}
