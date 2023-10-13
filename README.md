# vintage

`vintage` is a Fastly VCL runtime in a Edge.
This project provides runtime and transformer, vintage transpiles your VCLs into executable code on the edge (e.g Fastly Compute).

This project is subset of [falco](https://github.com/ysugimoto/falco), which is VCL linter and interpreter.
falco also includes this project to transpile VCLs inside that tool, then vintage provides the runtimes on execution.
Before the transpilation, vintage checks your VCLs by falco linter so you need to have a valid VCLs on the falco.

## CLI

Donwload cli command from [release]() page and place it at your `$PATH`.
Example of transpilation is the following:

```shell
vintage transpile --target compute --package main --output vintage.go
```

Describe CLI option:

| option name   | required | default                  | description                                 |
|:-------------:|:--------:|:------------------------:|:--------------------------------------------|
| -t, --target  | no       | compute (Fastly Compute) | Transpile target of edge platform           |
| -p, --package | no       | main                     | Go package name of transpiled program       |
| -o, --output  | no       | ./vintage.go             | Output filename (should have .go extention) |

### Supported Runtimes

Supprted runtimes, which can specify on `-t, --target` cli option are following:

- `compute (default)` : Fastly Compute Runtime, the generated code could run in Compute@Edge
- `native` : Generates raw Golang code, could run in common platforms that Golang can compile to

## Run Application With Generated Code

After transpilation succeeded, you can get single go file that at `--output` cli option.
The generated file exposes `VclHandler` function that implements server handler corresponds to target platform.
For example, `fsthttp.Handler` for Fastly Compute or `http.Handler` for native.

To work application correctly, you need to do following steps:

1. The generated code has some dependecies so needs to run `go mod tidy` to install dependencies
2. If you transpiled for Fastly Compute, need to set up `fastly.toml` for Golang. see [documentation](https://developer.fastly.com/learning/compute/go/) in detail
3. You need to implement a little to start server. Fastly compute example is the following:

```go
package main

import (
	"github.com/fastly/compute-sdk-go/fsthttp"
)

func main() {
	fsthttp.Serve(VclHandler())
}
```

On native:

```go
package main

import (
	"net/http"
)

func main() {
    http.Handle("/", VclHandler())
    http.ListenAndServe(":8888", nil)
}
```

It's very tiny implementation! The handle has common Golang HTTP handler interface, of course you can combinate with HTTP middlewares, your favorite HTTP framework.

Note that on Fastly Compute Runtime, some packages are limited to use so ensure your framework or library can use in Compute Runtime.

## Welcome Your Feedback!

We tested many kind of VCLs (production use, massive lines and include modules), they could work fine but it's not enough.
If you can try `vintage` in your VCLs and found any problems or weired behavior, we welcome your feedback on an issue.

## Contribution

- Fork this repository
- Customize / Fix problem
- Send PR :-)
- Or feel free to create issues for us. We'll look into it

## License

MIT License

## Contributors

- [@ysugimoto](https://github.com/ysugimoto)
