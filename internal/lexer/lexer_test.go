package lexer

import (
	"testing"
)

func Test_lexer_Tokenize(t *testing.T) {
	l := NewLexer("age >?;")
	reader, err := l.Tokenize()
	if err != nil {
		t.Fail()
	}
	reader.Dump()
}
