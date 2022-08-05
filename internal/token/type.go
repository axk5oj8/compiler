package token

const (
	Plus Type = iota
	GE
	GT
	EQ
	Identifier
)

type Type uint
