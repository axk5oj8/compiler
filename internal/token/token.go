package token

type Token interface {
	Type() Type
	String() string
}

func NewToken(typ Type, txt string) Token {
	return &token{typ, txt}
}

type token struct {
	typ Type
	txt string
}

func (t *token) Type() Type {
	return t.typ
}

func (t *token) String() string {
	return t.txt
}
