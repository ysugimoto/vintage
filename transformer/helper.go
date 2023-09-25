package transformer

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
	case *ast.RTime:
		switch {
		case strings.HasSuffix(t.Value, "ms"):
			value := strings.TrimSuffix(t.Value, "ms")
			return fmt.Sprintf("%s * time.Millisecond", value)
		case strings.HasSuffix(t.Value, "s"):
			value := strings.TrimSuffix(t.Value, "s")
			return fmt.Sprintf("%s * time.Second", value)
		case strings.HasSuffix(t.Value, "m"):
			value := strings.TrimSuffix(t.Value, "m")
			return fmt.Sprintf("%s * time.Minute", value)
		case strings.HasSuffix(t.Value, "h"):
			value := strings.TrimSuffix(t.Value, "h")
			return fmt.Sprintf("%s * time.Hour", value)
		case strings.HasSuffix(t.Value, "d"):
			value := strings.TrimSuffix(t.Value, "d")
			return fmt.Sprintf("%s * 24 * time.Hour", value)
		case strings.HasSuffix(t.Value, "y"):
			value := strings.TrimSuffix(t.Value, "y")
			return fmt.Sprintf("%s * 24 * 365 * time.Hour", value)
		}
		return ""
	}
	return ""
}

func ucFirst(str string) string {
	b := []byte(str)
	b[0] -= 0x20
	return string(b)
}

func PrepareCodes(preps ...string) string {
	var code []string
	for i := range preps {
		if preps[i] == "" {
			continue
		}
		code = append(code, preps[i])
	}
	return strings.Join(code, lineFeed)
}

var tmpVarCounter uint

func Temporary() string {
	tmpVarCounter++
	return fmt.Sprintf("tmp_%d", tmpVarCounter)
}
