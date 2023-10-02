package main

import (
	"fmt"

	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/vintage/transformer/fastly"
)

func main() {
	rslv, err := resolver.NewFileResolvers("../../examples/default01.vcl", []string{})
	if err != nil {
		panic(err)
	}
	buf, err := fastly.NewFastlyTransformer().Transform(rslv[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
