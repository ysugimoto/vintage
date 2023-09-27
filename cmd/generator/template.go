package main

const tmplVariables = `
package variable

import "github.com/ysugimoto/vintage/transformer/value"

// VariablesImpl is underlying variable implementation.
// All get/set/unset variables will raise an error of NotImplementedError,
// intend to check all variables must be implemented.
// So it means that all variables implementation must be extended this struct
type VariablesImpl struct {
}

func (v *VariablesImpl) Get(name string) (*value.Value, error) {
	switch (name) {
	{{- range .GetVariables}}
	case {{ toUpper . }}: return nil, ErrNotImplemented(name)
	{{- end}}
	}

	return nil, ErrNotFound(name)
}

func (v *VariablesImpl) Set(name string, value *value.Value) error {
	switch (name) {
	{{- range .SetVariables}}
	case {{ toUpper . }}: return ErrNotImplemented(name)
	{{- end}}
	}

	return ErrCannotSet(name)
}

func (v *VariablesImpl) Unset(name string) error {
	switch (name) {
	{{- range .UnsetVariables}}
	case {{ toUpper . }}: return ErrNotImplemented(name)
	{{- end}}
	}

	return ErrCannotUnset(name)
}
`

const tmplPredefinedMap = `
package variable

const (
	{{- range .Variables}}
	{{ toUpper . }} = "{{ . }}"
	{{- end }}
)

var Getable = []string{
	{{- range .GetVariables}}
	"{{ . }}",
	{{- end}}
}

var Setable = []string{
	{{- range .SetVariables}}
	"{{ . }}",
	{{- end}}
}

var Unsetable = []string{
	{{- range .UnsetVariables}}
	"{{ . }}",
	{{- end}}
}
`
