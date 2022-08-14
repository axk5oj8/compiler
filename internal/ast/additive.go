package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/token"
)

// additive parse additive expression to Node
// Add -> Mul | Mul + Add;
func additive(reader token.Reader) (Node, error) {
	n, err := multiplicative(reader)
	if err != nil {
		return nil, err
	}

	if n == nil {
		return nil, nil
	}

	if !reader.PeekType(token.Plus) && !reader.PeekType(token.Minus) {
		return n, nil
	}

	tk := reader.Read() // take `+` or `-`
	rightChild, err := additive(reader)
	if err != nil {
		return n, err
	}
	if rightChild == nil {
		return n, errors.New("invalid additive expression, expect the right part")
	}

	leftChild := n
	n = NewNode(Additive, tk.Text())
	n.AddChild(leftChild)
	n.AddChild(rightChild)	

	return n, nil
}
