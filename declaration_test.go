package vintage

import (
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestBackendDeclaration(t *testing.T) {
	b := NewBackend("F_example", BackendDefault())
	if diff := cmp.Diff(b.Backend(), "F_example"); diff != "" {
		t.Errorf("Value unmatch, diff=%s", diff)
	}
}

func TestAclDeclaration(t *testing.T) {
	tests := []struct {
		acl    *Acl
		expect string
	}{
		{acl: NewAcl("example", AclEntry("192.168.0.0/32", false)), expect: "192.168.0.0"},
		{acl: NewAcl("example", AclEntry("192.168.0.0/32", true)), expect: "192.168.0.1"},
		{acl: NewAcl("example", AclEntry("192.168.0.0/24", false)), expect: "192.168.0.255"},
		{acl: NewAcl("example", AclEntry("192.168.0.0/24", true)), expect: "192.168.1.0"},
	}

	for i, tt := range tests {
		if diff := cmp.Diff(tt.acl.Match(net.ParseIP(tt.expect)), true); diff != "" {
			t.Errorf("[%d] Value unmatch, diff=%s", i, diff)
		}
	}
}

func TestTableDeclaration(t *testing.T) {
	st := NewTable("example", "STRING", TableItem("foo", "bar"))
	if diff := cmp.Diff(st.Items["foo"], "bar"); diff != "" {
		t.Errorf("String Table: Value unmatch, diff=%s", diff)
	}
	it := NewTable("example", "INTEGER", TableItem("foo", 100))
	if diff := cmp.Diff(it.Items["foo"], 100); diff != "" {
		t.Errorf("Integer Table: Value unmatch, diff=%s", diff)
	}
	ft := NewTable("example", "FLOAT", TableItem("foo", 100.001))
	if diff := cmp.Diff(ft.Items["foo"], 100.001); diff != "" {
		t.Errorf("Float Table: Value unmatch, diff=%s", diff)
	}
	bt := NewTable("example", "BOOL", TableItem("foo", true))
	if diff := cmp.Diff(bt.Items["foo"], true); diff != "" {
		t.Errorf("Bool Table: Value unmatch, diff=%s", diff)
	}
	rt := NewTable("example", "RTIME", TableItem("foo", time.Second))
	if diff := cmp.Diff(rt.Items["foo"], time.Second); diff != "" {
		t.Errorf("RTime Table: Value unmatch, diff=%s", diff)
	}
	now := time.Now()
	tt := NewTable("example", "TIME", TableItem("foo", now))
	if diff := cmp.Diff(tt.Items["foo"], now); diff != "" {
		t.Errorf("Time Table: Value unmatch, diff=%s", diff)
	}
}
