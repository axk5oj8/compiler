package token

type Reader interface {
	// Read will take next Token from token stream and then return. Nil if stream is empty.
	Read() Token
	// Peek will return next Token from token stream, but not take. Nil if stream is empty.
	Peek() Token

	UnRead()

	Where() int

	At(int) error
}
