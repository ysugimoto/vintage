package value

import (
	"fmt"
	"strings"
)

type ValueOption func(e *Value)

func Prepare(preps ...string) ValueOption {
	return func(e *Value) {
		var codes []string
		for i := range preps {
			if preps[i] == "" {
				continue
			}
			codes = append(codes, preps[i])
		}
		if len(codes) > 0 {
			e.Prepare = strings.Join(codes, "\n") + "\n"
		}
	}
}

func Dependency(pkg, alias string) ValueOption {
	return func(e *Value) {
		if e.Dependencies == nil {
			e.Dependencies = Packages{}
		}
		e.Dependencies[pkg] = Package{Alias: alias}
	}
}

func Comment(c string) ValueOption {
	return func(e *Value) {
		e.Comment = c
	}
}

func Deprecated() ValueOption {
	return func(e *Value) {
		e.Deprecated = true
	}
}

func FromValue(v *Value) ValueOption {
	return func(e *Value) {
		e.Prepare += v.Prepare
		e.Comment += v.Comment
		if v.Dependencies != nil {
			for key, val := range v.Dependencies {
				e.Dependencies[key] = val
			}
		}
	}
}

type Value struct {
	Type         VCLType
	Code         string
	Prepare      string
	Dependencies Packages
	Comment      string
	Deprecated   bool
}

func NewValue(t VCLType, code string, preps ...ValueOption) *Value {
	v := &Value{
		Type: t,
		Code: code,
	}
	for i := range preps {
		preps[i](v)
	}
	return v
}

func (v *Value) String() string {
	out := v.Code
	if v.Comment != "" {
		out += " /*"
		if v.Deprecated {
			out += " @deprecated "
		}
		out += v.Comment + "*/"
	}
	return out
}

func (v *Value) Conversion(expect VCLType) *Value {
	if expect == NULL {
		return v
	}

	switch expect {
	case STRING:
		return v.stringConversion()
	case BOOL:
		return v.boolConversion()
	}
	v.Type = expect
	return v
}

func (v *Value) stringConversion() *Value {
	if v.Type == STRING {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToString(%s)`, v.Code)
	v.Type = STRING
	return v
}

func (v *Value) boolConversion() *Value {
	if v.Type == BOOL {
		return v
	}

	v.Code = fmt.Sprintf(`vintage.ToBool(%s)`, v.Code)
	v.Type = BOOL
	return v
}

var temporaryVarCount int
var useFixedName bool

func UseFixedTemporalValue() {
	useFixedName = true
}

func Temporary() string {
	if useFixedName {
		return "tmp__fixed"
	}
	temporaryVarCount++
	return fmt.Sprintf("tmp__%d", temporaryVarCount)
}

var ErrorCheck = "if err != nil {\nreturn vintage.NONE, err\n}"
