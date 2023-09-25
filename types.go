package vintage

import (
	"net"
	"time"
)

type DirectorType string

const (
	Random   DirectorType = "random"
	Fallback DirectorType = "fallback"
	Hash     DirectorType = "hash"
	Client   DirectorType = "client"
	CHash    DirectorType = "chash"
)

type Primitive interface {
	string | int64 | float64 | bool | net.IP | time.Duration | time.Time | *Backend | *Acl
}

type VCLType string

const (
	IDENT   VCLType = "IDENT"
	INTEGER VCLType = "INTEGER"
	FLOAT   VCLType = "FLOAT"
	STRING  VCLType = "STRING"
	BOOL    VCLType = "BOOL"
	IP      VCLType = "IP"
	RTIME   VCLType = "RTIME"
	TIME    VCLType = "TIME"
	BACKEND VCLType = "BACKEND"
	ACL     VCLType = "ACL"
	NULL    VCLType = "NULL"
)

func GoTypeString(t VCLType) string {
	switch t {
	case INTEGER:
		return "int64"
	case FLOAT:
		return "float64"
	case STRING:
		return "string"
	case BOOL:
		return "bool"
	case IP:
		return "net.IP"
	case RTIME:
		return "time.Duration"
	case TIME:
		return "time.Time"
	case BACKEND:
		return "*vintage.Backend"
	case ACL:
		return "*vintage.Acl"
	}
	return ""
}