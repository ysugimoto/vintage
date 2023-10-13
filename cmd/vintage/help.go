package main

import (
	"fmt"
	"strings"
)

func printHelp() {
	help := strings.TrimSpace(`
=====================================================
  _    ___       __
 | |  / (_)___  / /_____ _____ ____ 
 | | / / / __ \/ __/ __ ` + "`" + `/ __ ` + "`" + `/ _ \
 | |/ / / / / / /_/ /_/ / /_/ /  __/
 |___/_/_/ /_/\__/\__,_/\__, /\___/ 
                       /____/       VCL into a Edge
=====================================================

Usage:
    vintage [flags] [entry vcl file]

Flags:
    -I, --include_path : Add include path
    -h, --help         : Show this help
    -o, --output       : Specify output filename
    -t, --target       : Specify transpile target
    -p, --package      : Specify package name

Simple tranpilation example:
    vintage -o vintage.go -p main /path/to/vcl/main.vcl

`)
	fmt.Println(help)
}
