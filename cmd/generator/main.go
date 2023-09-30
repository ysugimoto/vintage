package main

func main() {
	if err := generatePredefinedScript(); err != nil {
		panic(err)
	}

	if err := generateBuiltinFunctionScript(); err != nil {
		panic(err)
	}
}
