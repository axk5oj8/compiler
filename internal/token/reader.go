package token

import "fmt"

type Reader interface {
	// Read will take next Token from token stream and then return. Nil if stream is empty.
	Read() Token
	// Peek will return next Token from token stream, but not take. Nil if stream is empty.
	Peek() Token

	UnRead()

	Where() int

	GoTo(int)

	Dump()
}

func NewReader(tokens []Token) Reader {
	return &reader{
		pos:    0,
		tokens: tokens,
	}
}

type reader struct {
	pos    int
	tokens []Token
}

func (r *reader) Read() Token {
	if r.pos >= len(r.tokens) {
		return nil
	}
	r.pos++
	return r.tokens[r.pos-1]
}

func (r *reader) Peek() Token {
	if r.pos >= len(r.tokens) {
		return nil
	}

	return r.tokens[r.pos]
}

func (r *reader) UnRead() {
	if r.pos > 0 {
		r.pos--
	}
}

func (r *reader) Where() int {
	return r.pos
}

func (r *reader) GoTo(i int) {
	if len(r.tokens) > i && i >= 0 {
		r.pos = i
	}
}

func (r *reader) Dump() {
	fmt.Println("text\t\ttype")

	var t Token

	for {
		if t = r.Read(); t == nil {
			return
		}
		fmt.Printf("%s\t\t%s\n", t, t.Type())
	}
}
