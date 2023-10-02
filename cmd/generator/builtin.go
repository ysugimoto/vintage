package main

import (
	"bytes"
	"go/format"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
)

type FunctionSpec struct {
	Arguments [][]string `yaml:"arguments"`
	Return    string     `yaml:"return"`
	Extra     string     `yaml:"extra"`
	On        []string   `yaml:"on"`
	Ref       string     `yaml:"reference"`

	// Following fields are set inside generation
	Name                        string
	RequiredArguments           []string
	OptionalArguments           []string
	VariadicArgumentStartsIndex int
}

func (f *FunctionSpec) GenerateRuntimeSpec() {
	for i, args := range f.Arguments {
		var updates []string
		for j, a := range args {
			switch a {
			case "ID", "TABLE":
				updates = append(updates, "IDENT")
			case "STRING_LIST":
				f.VariadicArgumentStartsIndex = j
			default:
				updates = append(updates, a)
			}
		}
		f.Arguments[i] = updates
	}
	// Sort desc by argument length
	sort.Slice(f.Arguments, func(i, j int) bool {
		return len(f.Arguments[i]) < len(f.Arguments[j])
	})
	if len(f.Arguments) == 0 {
		return
	}
	f.RequiredArguments = f.Arguments[len(f.Arguments)-1]
	for _, args := range f.Arguments {
		if len(args) >= len(f.RequiredArguments) {
			continue
		}
		f.OptionalArguments = append(f.OptionalArguments, f.RequiredArguments[len(args):]...)
		f.RequiredArguments = args
	}
}

func (f *FunctionSpec) ReturnType() string {
	if f.Return == "" {
		return "value.NULL"
	}
	return "value." + strings.ToUpper(f.Return)
}

func generateBuiltinFunctionScript() error {
	fp, err := os.Open("./builtin.yml")
	if err != nil {
		return err
	}
	defer fp.Close()

	defs := map[string]*FunctionSpec{}
	if err := yaml.NewDecoder(fp).Decode(&defs); err != nil {
		return err
	}

	var functions []*FunctionSpec
	for key, def := range defs {
		def.Name = key
		def.GenerateRuntimeSpec()

		functions = append(functions, def)
	}

	sort.Slice(functions, func(i, j int) bool {
		return functions[i].Name < functions[j].Name
	})

	out := new(bytes.Buffer)
	tpl := template.New("vintage.builtin")
	tpl.Funcs(template.FuncMap{
		"toVintageFunction": func(text string) string {
			snake := []byte(strings.ReplaceAll(text, ".", "_"))
			snake[0] -= 0x20
			return "builtin." + string(snake)
		},
	})
	tpl = template.Must(tpl.Parse(tmplBuiltinFunction))
	if err := tpl.Execute(out, map[string]any{
		"Functions": functions,
	}); err != nil {
		return err
	}

	ret, err := format.Source(out.Bytes())
	if err != nil {
		return err
	}
	f, err := os.OpenFile("../../transformer/core/builtin_functions.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(ret); err != nil {
		return err
	}

	return nil
}
