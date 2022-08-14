package main

import (
	"fmt"
	"os"

	"github.com/axk5oj8/compiler/internal/ast"
)

func main() {
	script := os.Args[1]

	tree, err := ast.Parse(script)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ast.Dump(tree, "")
}