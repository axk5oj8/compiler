package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/token"
)

// intDeclareStmt parse int variable declaration statement like
// `int x;` or `int x = 1+2*3;` into Node
func intDeclareStmt(reader token.Reader) (Node, error) {
	// check if next token is `int`
	if !reader.PeekType(token.Int) {
		return nil, nil
	}

	reader.Read() // take `int`
		
	if !reader.PeekType(token.Identifier) {
		return nil, errors.New("variable name is expected")
	}

	tk := reader.Read() // take `identifier`
	n := NewNode(IntDeclaration, tk.Text())

	// check if next token is `=`
	if reader.PeekType(token.Assignment) {
		reader.Read()
		child, err := additive(reader)
		if err != nil {
			return nil, err
		}
		if child == nil {
			return nil, errors.New("invalid variable initialization, expect an expression")
		}
		n.AddChild(child)
	}

	// check if end with `;`
	if !reader.PeekType(token.SemiColon) {
		return n, errors.New("invalid statement, expect semicolon")
	}
	reader.Read() // take `;`

	return n, nil
}
