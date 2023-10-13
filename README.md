# Vintage

`Vintage` is a Fastly VCL runtime inside a Edge.
This project provides runtime and transformer, vintage transpiles your VCLs into executable code on the edge (e.g Fastly Compute).

This project is subset of [falco](https://github.com/ysugimoto/falco), which is VCL linter and interpreter.
falco also includes this project to transpile VCLs inside that tool, then vintage provides the runtimes on execution.
Before the transpilation, vintage checks your VCLs by falco linter so you need to have a valid VCLs on the falco.

## How to use

### CLI

Donwload cli command from [release]() page and place it at your `$PATH`.
Example of transpilation is the following:

```shell
vintage transpile --target compute --package main --output vintage.go
```

Describe CLI option:

| option name   | required | default                  | description                               |
|:=============:|:========:|:========================:|:==========================================|
| -t, --target  | no       | compute (Fastly Compute) | Transpile target of edge platform         |
| -p, --package | no       | main                     | Go package name of transpiled program     |
| -o, --output  | no       | vintage.go               | Output filename (must have .go extention) |

## Programmable

vintage depends on `falco` so you need to install as dependency.
Below example tranpiles with Fastly remote snippets:

```go
package main

import (
	"os"
	"time"

	"github.com/ysugimoto/falco/remote"
	"github.com/ysugimoto/falco/resolver"
	"github.com/ysugimoto/falco/snippets"
	"github.com/ysugimoto/vintage/transformer/core"
	"github.com/ysugimoto/vintage/transformer/fastly"
)

func main() {
    // Set up VCL resolver to include another modules
	rslv, err := resolver.NewFileResolvers("/path/to/entrypoint.vcl", []string{})
	if err != nil {
		panic(err)
	}

    // Initialize remote resource fetcher.
    // You need to specify Fastly service id and api token to fetch resources
	fetcher := remote.NewFastlyApiFetcher(
		os.Getenv("FASTLY_SERVICE_ID"),
		os.Getenv("FASTLY_API_TOKEN"),
		10*time.Second,
	)
    // Do fetch remote resources
	s, err := snippets.Fetch(fetcher)
	if err != nil {
		panic(err)
	}
	if err := s.FetchLoggingEndpoint(fetcher); err != nil {
		panic(err)
	}

    // Let's transform - buf variable is []byte of excutable go code
	buf, err := fastly.NewFastlyTransformer(core.WithSnippets(s)).Transform(rslv[0])
	if err != nil {
		panic(err)
	}

    // Write buffer to the file
	fp, err := os.OpenFile("/path/to/compute-project/vintage.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	fp.Write(buf)
}

