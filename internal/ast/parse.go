package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/lexer"
	"github.com/axk5oj8/compiler/internal/token"
)

type Parser func(token.Reader) (Node, error)

var stmtParsers []Parser

func Parse(script string) (Node, error) {
	lxr := lexer.NewLexer(script)
	reader, err := lxr.Tokenize()
	if err != nil {
		return nil, err
	}

	root, err := prog(reader)
	if err != nil {
		return nil, err
	}

	return root, nil
}

func prog(reader token.Reader) (Node, error) {
	n := NewNode(Program, "pwc")

	for reader.Peek() != nil {
		var child Node
		var err error
		for _, sp := range stmtParsers {
			child, err = sp(reader)
			if err != nil {
				return nil, err
			}
			if child != nil {
				n.AddChild(child)
				break
			}
		}
		if child == nil {
			return n, errors.New("unknown statement")
		}
	}
	return n, nil
}

func init() {
	stmtParsers = append(stmtParsers, intDeclareStmt, expressionStmt, assignmentStmt)
}