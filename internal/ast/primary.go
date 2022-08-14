package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/token"
)

// primary parse primary expression to Node
// Pri -> Id | Num | (Add)
func primary(reader token.Reader) (Node, error) {
	tk := reader.Peek()
	if tk == nil {
		return nil, nil
	}

	switch tk.Type() {
	case token.IntLiteral:
		reader.Read()
		return NewNode(IntLiteral, tk.Text()), nil
	case token.Identifier:
		reader.Read()
		return NewNode(Identifier, tk.Text()), nil
	case token.LeftParen:
		reader.Read()
		n, err := additive(reader)
		if err != nil || n == nil {
			return nil, errors.New("expect an additive expression inside parenthesis")
		}

		if reader.PeekType(token.RightParen) {
			return nil, errors.New("expect right parenthesis") 
		}

		reader.Read()
		return n, nil
	default:
		return nil, nil
	}
}