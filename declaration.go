package vintage

import "net"

type Backend struct {
	Name     string
	Director *Director
}

func NewBackend(name string) *Backend {
	return &Backend{
		Name: name,
	}
}

func (b *Backend) Backend() string {
	if b.Director != nil {
		return b.Director.Backend()
	}
	return b.Name
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

type Table map[string]Value
type TableOption func(t Table)

func TableItem(name string, value Value) TableOption {
	return func(t Table) {
		t[name] = value
	}
}

func NewTable(items ...TableOption) Table {

	t := make(map[string]Value)
	for i := range items {
		items[i](t)
	}
	return t
}
