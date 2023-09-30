package value

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

func DefaultValue(t VCLType) string {
	switch t {
	case INTEGER:
		return "0"
	case FLOAT:
		return "0"
	case STRING:
		return `""`
	case BOOL:
		return "false"
	case IP:
		return "nil"
	case RTIME:
		return "time.Duration(0)"
	case TIME:
		return "time.Time{}"
	case BACKEND:
		return "nil"
	case ACL:
		return "nil"
	}
	return ""
}
