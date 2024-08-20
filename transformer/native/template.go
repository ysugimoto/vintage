package native

const handlerTemplate = `
// Code generated by vintage native transformer; DO NOT EDIT.

//nolint // Generated code may have invalid syntaxes, should be ignored
package {{ .OutputPackage }}

import (
  "fmt"

  {{range .Packages}}
  {{.}}
  {{- end}}
)

{{ .Declarations }}

func errorHandler(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err.Error())
}

func VclHandler(opts ...HandlerOption) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		runtime, err := native.NewRuntime(w, r)
		if err != nil {
			errorHandler(w, err)
			return
		}
		for i := range opts {
			opts[i](r)
		}

		runtime.Register(
			{{- range $key, $val := .Backends}}
			native.BackendResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Acls}}
			native.AclResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Tables}}
			native.TableResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Subroutines}}
			native.SubroutineResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
		)
		if err := runtime.Execute(r.Context()); err != nil {
			errorHandler(w, err)
		}
	})
}

type HandlerOption func(r *native.Runtime)

func WithCacheDriver(driver vintage.CacheDriver) HandlerOption {
	return func(r *native.Runtime) {
		r.Cache = driver
	}
}
`
