package main

import (
	"os"

	"github.com/axk5oj8/compiler/internal/lexer"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		l := lexer.NewLexer(arg)
		if reader, err := l.Tokenize(); err == nil {
			reader.Dump()
		}
	}
}