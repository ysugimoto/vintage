package fastly

const handlerTemplate = `
package main

import (
  "fmt"
  "context"
  "io"

  {{range $key, $val := .Packages}}
  {{if $val.Alias }}{{ $val.Alias }} {{end}}"{{ $key }}"
  {{- end}}
)

{{ .Declarations }}

func VclEdgeHander() fsthttp.Handler {
	return fsthttp.HandlerFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		runtime := fastly.NewRuntime(r)
		runtime.Register(
			{{- range $key, $val := .Backends}}
			vintage.BackendResource[*fastly.Runtime]("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Acls}}
			vintage.AclResource[*fastly.Runtime]("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Tables}}
			vintage.TableResource[*fastly.Runtime]("{{ $key }}", {{ $val.String }}),
			{{- end}}
			{{- range $key, $val := .Subroutines}}
			vintage.SubroutineResource("{{ $key }}", {{ $val.String }}),
			{{- end}}
		)
		resp, err := runtime.Execute(ctx)
		if err != nil {
			w.WriteHeader(fsthttp.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
		w.Header().Reset(resp.Header)
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	})
}
`
