package main

import (
	"fmt"

	"github.com/ysugimoto/falco/resolver"
)

func main() {
	rslv, err := resolver.NewFileResolvers("../../examples/default01.vcl", []string{})
	if err != nil {
		panic(err)
	}
	buf, err := transformer.New(rslv).Transform()
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
