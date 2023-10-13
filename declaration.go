package vintage

import (
	"io"
	"net"
	"time"

	"github.com/pkg/errors"
)

type Backend struct {
	Name                string
	IsDefault           bool
	Port                string
	Host                string
	SSL                 bool
	AlwaysUseHostHeader bool
	ConnectTimeout      time.Duration
	FirstByteTimeout    time.Duration
	BetweenBytesTimeout time.Duration
	Director            *Director
}

func NewBackend(name string, opts ...BackendOption) *Backend {
	b := &Backend{
		Name:                name,
		ConnectTimeout:      time.Second,
		FirstByteTimeout:    15 * time.Second,
		BetweenBytesTimeout: 10 * time.Second,
	}
	for i := range opts {
		opts[i](b)
	}
	return b
}

func (b *Backend) Backend(ident RequestIdentity) string {
	if b.Director != nil {
		return b.Director.Backend(ident)
	}
	return b.Name
}

type BackendOption func(b *Backend)

func BackendDefault() BackendOption {
	return func(b *Backend) {
		b.IsDefault = true
	}
}

func BackendPort(port string) BackendOption {
	return func(b *Backend) {
		b.Port = port
	}
}

func BackendHost(host string) BackendOption {
	return func(b *Backend) {
		b.Host = host
	}
}

func BackendSSL(ssl bool) BackendOption {
	return func(b *Backend) {
		b.SSL = ssl
	}
}

func BackendConnectTimeout(t time.Duration) BackendOption {
	return func(b *Backend) {
		b.ConnectTimeout = t
	}
}

func BackendFirstByteTimeout(t time.Duration) BackendOption {
	return func(b *Backend) {
		b.FirstByteTimeout = t
	}
}

func BackendBetweenBytesTimeout(t time.Duration) BackendOption {
	return func(b *Backend) {
		b.BetweenBytesTimeout = t
	}
}

func BackendAlwaysUseHostHeader(v bool) BackendOption {
	return func(b *Backend) {
		b.AlwaysUseHostHeader = v
	}
}

type Acl struct {
	Name  string
	CIDRs []struct {
		IpNet   *net.IPNet
		Inverse bool
	}
}

func NewAcl(name string, opts ...AclOption) *Acl {
	a := &Acl{
		Name: name,
	}
	for i := range opts {
		opts[i](a)
	}
	return a
}

func (a *Acl) Match(ip net.IP) bool {
	for _, entry := range a.CIDRs {
		contains := entry.IpNet.Contains(ip)
		if contains {
			return true
		}
		if entry.Inverse {
			return true
		}
	}
	return false
}

type AclOption func(a *Acl)

func AclEntry(cidr string, inverse bool) AclOption {
	return func(a *Acl) {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		a.CIDRs = append(a.CIDRs, struct {
			IpNet   *net.IPNet
			Inverse bool
		}{
			IpNet:   ipnet,
			Inverse: inverse,
		})
	}
}

type Table struct {
	Name             string
	Type             string
	Items            map[string]any
	isEdgeDictionary bool
}

type TableOption func(t *Table)

func TableItem(name string, value any) TableOption {
	return func(t *Table) {
		t.Items[name] = value
	}
}

func EdgeDictionary() TableOption {
	return func(t *Table) {
		t.isEdgeDictionary = true
	}
}

func NewTable(name, itemType string, items ...TableOption) *Table {
	t := &Table{
		Name:  name,
		Type:  itemType,
		Items: make(map[string]any),
	}

	for i := range items {
		items[i](t)
	}
	return t
}

func (t *Table) IsEdgeDictionary() bool {
	return t.isEdgeDictionary
}

type LoggingEndpoint struct {
	Name      string
	initiator LoggerInitiator
}

func NewLoggingEndpoint(name string, initiator LoggerInitiator) *LoggingEndpoint {
	return &LoggingEndpoint{
		Name:      name,
		initiator: initiator,
	}
}

func (l *LoggingEndpoint) Write(message string) error {
	w, err := l.initiator(l.Name)
	if err != nil {
		return errors.WithStack(err)
	}
	if _, err := io.WriteString(w, message); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
