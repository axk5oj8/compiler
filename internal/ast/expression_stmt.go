package ast

import "github.com/axk5oj8/compiler/internal/token"

// expressionStmt takes expression statment into Node
// an expression ends with `;`
func expressionStmt(reader token.Reader) (Node, error) {
	pos := reader.Where() // remember where we start
	n, err := additive(reader)	
	if err != nil {
		return nil, err
	}

	if n == nil {
		return nil, nil
	}

	if !reader.PeekType(token.SemiColon) {
		reader.GoTo(pos)
		return nil, nil
	}
	reader.Read()

	return n, nil
}