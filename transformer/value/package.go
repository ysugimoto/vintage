package value

import (
	"sort"
	"strings"
)

type Package struct {
	Alias string
}

type Packages map[string]Package

func (p Packages) Add(name, alias string) {
	p[name] = Package{
		Alias: alias,
	}
}

func (p Packages) Sorted() []string {
	var sorted []string
	for key, val := range p {
		pkg := `"` + key + `"`
		if val.Alias != "" {
			pkg = val.Alias + " " + pkg
		}
		sorted = append(sorted, pkg)
	}

	sort.Slice(sorted, func(i, j int) bool {
		a := strings.Split(sorted[i], "/")
		b := strings.Split(sorted[j], "/")
		if len(a) == len(b) {
			return len(sorted[i]) < len(sorted[j])
		}
		return len(a) < len(b)
	})

	return sorted
}
