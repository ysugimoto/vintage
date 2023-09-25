package core

type Package struct {
	Alias string
}

type Packages map[string]Package

func (p Packages) Add(name, alias string) {
	p[name] = Package{
		Alias: alias,
	}
}
