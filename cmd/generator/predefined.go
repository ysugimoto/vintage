package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/go-yaml/yaml"
)

type Definition struct {
	Get   string   `yaml:"get"`
	Set   string   `yaml:"set"`
	Unset bool     `yaml:"unset"`
	On    []string `yaml:"on"`
	Ref   string   `yaml:"reference"`
}

func generatePredefinedScript() error {
	fp, err := os.Open("./predefined.yml")
	if err != nil {
		return err
	}
	defer fp.Close()

	defs := map[string]*Definition{}
	if err := yaml.NewDecoder(fp).Decode(&defs); err != nil {
		return err
	}

	var (
		gets      []string
		sets      []string
		unsets    []string
		variables []string
	)

	for key, val := range defs {
		if strings.Contains(key, "%") {
			continue
		}
		if val.Get != "" {
			gets = append(gets, key)
		}
		if val.Set != "" {
			sets = append(sets, key)
		}
		if val.Unset {
			unsets = append(unsets, key)
		}
		variables = append(variables, key)
	}

	sort.Strings(gets)
	sort.Strings(sets)
	sort.Strings(unsets)
	sort.Strings(variables)

	vars := map[string]any{
		"GetVariables":   gets,
		"SetVariables":   sets,
		"UnsetVariables": unsets,
		"Variables":      variables,
	}
	if err := generateBaseImpl(vars); err != nil {
		return err
	}
	if err := generatePredefinedMap(vars); err != nil {
		return err
	}
	return nil
}

func generateBaseImpl(vars map[string]any) error {
	out := new(bytes.Buffer)
	tpl := template.New("vintage.baseImpl")
	tpl.Funcs(template.FuncMap{
		"toUpper": func(text string) string {
			return strings.ToUpper(strings.ReplaceAll(text, ".", "_"))
		},
	})
	tpl = template.Must(tpl.Parse(tmplVariables))
	if err := tpl.Execute(out, vars); err != nil {
		return err
	}

	ret, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Println(out.String())
		return err
	}
	f, err := os.OpenFile("../../transformer/variable/base.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(ret); err != nil {
		return err
	}

	return nil
}

func generatePredefinedMap(vars map[string]any) error {
	out := new(bytes.Buffer)
	tpl := template.New("vintage.predefined")
	tpl.Funcs(template.FuncMap{
		"toUpper": func(text string) string {
			return strings.ToUpper(strings.ReplaceAll(text, ".", "_"))
		},
	})
	tpl = template.Must(tpl.Parse(tmplPredefinedMap))
	if err := tpl.Execute(out, vars); err != nil {
		return err
	}

	ret, err := format.Source(out.Bytes())
	if err != nil {
		fmt.Println(out.String())
		return err
	}
	f, err := os.OpenFile("../../transformer/variable/predefined.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(ret); err != nil {
		return err
	}

	return nil
}
