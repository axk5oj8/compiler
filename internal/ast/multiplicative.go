package ast

import (
	"errors"

	"github.com/axk5oj8/compiler/internal/token"
)

// multiplicative parse multiplicative expression to Node
// Mul -> Pri | Pri * Mul
func multiplicative(reader token.Reader) (Node, error) {
	n, err := primary(reader)
	if err != nil {
		return nil, err
	}

	if !reader.PeekType(token.Asterisk) && !reader.PeekType(token.Slash) {
		return n, nil
	}

	tk := reader.Read() // take `*` or `/`
	rightChild, err := multiplicative(reader)
	if err != nil {
		return n, err
	}
	if rightChild == nil {
		return n, errors.New("invalid multiplicative expression, expect the right part")
	}

	leftChild := n
	n = NewNode(Multiplicative, tk.Text())
	n.AddChild(leftChild)
	n.AddChild(rightChild)
	
	return n, nil
}