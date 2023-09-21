package vintage

import "net/textproto"

type Context struct {
	Backend        *Backend
	RequestHeader  *Header
	ResponseHeader *Header

	Backends map[string]*Backend
	Acls     map[string]*Acl
	Tables   map[string]*Table
}

func NewContext(w, r map[string][]string) *Context {
	return &Context{
		RequestHeader:  NewHeader(textproto.MIMEHeader(r)),
		ResponseHeader: NewHeader(textproto.MIMEHeader(w)),
	}
}

func (c *Context) Register(name string, v any) {
	switch t := v.(type) {
	case *Backend:
		c.Backends[name] = t
	case *Acl:
		c.Acls[name] = t
	case *Table:
		c.Tables[name] = t
	}
}
