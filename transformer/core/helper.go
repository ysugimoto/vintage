package core

import (
	"fmt"
	"strings"

	"github.com/ysugimoto/falco/ast"
)

func toString(expr ast.Expression) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return `"` + t.Value + `"`
	case *ast.String:
		return `"` + t.Value + `"`
	case *ast.IP:
		return `"` + t.Value + `"`
	case *ast.Integer:
		return fmt.Sprint(t.Value)
	case *ast.Float:
		return fmt.Sprint(t.Value)
	case *ast.Boolean:
		return fmt.Sprintf("%t", t.Value)
	case *ast.RTime:
		switch {
		case strings.HasSuffix(t.Value, "ms"):
			value := strings.TrimSuffix(t.Value, "ms")
			return fmt.Sprintf("(%s * vintage.Millisecond)", value)
		case strings.HasSuffix(t.Value, "s"):
			value := strings.TrimSuffix(t.Value, "s")
			return fmt.Sprintf("(%s * vintage.Second)", value)
		case strings.HasSuffix(t.Value, "m"):
			value := strings.TrimSuffix(t.Value, "m")
			return fmt.Sprintf("(%s * vintage.Minute)", value)
		case strings.HasSuffix(t.Value, "h"):
			value := strings.TrimSuffix(t.Value, "h")
			return fmt.Sprintf("(%s * vintage.Hour)", value)
		case strings.HasSuffix(t.Value, "d"):
			value := strings.TrimSuffix(t.Value, "d")
			return fmt.Sprintf("(%s * vintage.Day)", value)
		case strings.HasSuffix(t.Value, "y"):
			value := strings.TrimSuffix(t.Value, "y")
			return fmt.Sprintf("(%s * vintage.Year)", value)
		}
		return ""
	}
	return ""
}

type RegexMatchedGroupStack struct {
	groups []string
}

func (r *RegexMatchedGroupStack) Push(v string) {
	r.groups = append(r.groups, v)
}

func (r *RegexMatchedGroupStack) Pop() {
	if len(r.groups) > 1 {
		r.groups = r.groups[:len(r.groups)-1]
	} else {
		r.groups = []string{}
	}
}

func (r *RegexMatchedGroupStack) Last() string {
	if len(r.groups) > 0 {
		return r.groups[len(r.groups)-1]
	}
	return ""
}
