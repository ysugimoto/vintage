package fastly

const handlerTemplate = `
package main

import (
  "fmt"
  "context"

  {{range .Packages}}
  {{.}}
  {{- end}}
)

{{ .Declarations }}

func errorHandler(w fsthttp.ResponseWriter, err error) {
	w.WriteHeader(fsthttp.StatusInternalServerError)
	fmt.Fprint(w, err.Error())
}

func VclEdgeHander() fsthttp.Handler {
	return fsthttp.HandlerFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		runtime, err := fastly.NewRuntime(w, r)
		if err != nil {
			errorHandler(w, err)
			return
		}

		runtime.Register(
			{{- range $key, $val := .Backends}}
			fastly.BackendResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Acls}}
			fastly.AclResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Tables}}
			fastly.TableResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
			{{- range $key, $val := .Subroutines}}
			fastly.SubroutineResource("{{ $key }}", {{ $val.Code }}),
			{{- end}}
		)
		if err := runtime.Execute(ctx); err != nil {
			errorHandler(w, err)
		}
	})
}
`
