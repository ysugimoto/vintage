package transformer

const handlerTemplate = `
package main

import (
  "fmt"
  {{range $key, $val := .Packages}}
  {{if $val.Alias }}{{ $val.Alias }} {{end}}"{{ $key }}"
  {{- end}}
)

{{ .Declarations }}

func VclEdgeHander() fsthttp.HandlerFunc {
	return fsthttp.HandlerFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		runtime := fastly.NewContext(w, r)
		runtime.Register(
			{{- range $key, $val := .Backends}}
			runtime.BackendResource("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Acls}}
			runtime.AclResource("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Tables}}
			runtime.TableResource("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Subroutines}}
			runtime.SubroutineResource("{{ $key }}", {{ $val }}),
			{{- end}}
		)
		runtime.Execute(ctx)
	})
}
`
