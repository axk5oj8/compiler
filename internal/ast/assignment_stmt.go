package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/token"
)

// assignmentStmt takes assignment statement into Node
// x = 3*4;
func assignmentStmt(reader token.Reader) (Node, error) {
	if !reader.PeekType(token.Identifier) {
		return nil, nil
	}
	
	tk := reader.Read()

	if !reader.PeekType(token.Assignment) {
		reader.UnRead()
		return nil, nil
	}

	n := NewNode(AssignmentStmt, tk.Text())
	reader.Read()

	child, err := additive(reader)
	if err != nil {
		return n, err
	}
	if child == nil {
		return n, errors.New("invalid assignment satement, expect an expression")
	}

	n.AddChild(child)

	if !reader.PeekType(token.SemiColon) {
		return n, errors.New("invalid statement, expect semicolon")
	}
	reader.Read()

	return n, nil
}