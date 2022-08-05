package lexer

import (
	"errors"
	"github.com/axk5oj8/compiler/internal/token"
	"strings"
)

type Lexer interface {
	Tokenize(code string) ([]token.Token, error)
}

type lexer struct {
	buffer []byte
	tokens []token.Token

	pos   int
	state DFAState
}

func (l *lexer) Tokenize(code string) ([]token.Token, error) {
	l.buffer = []byte(code)

	for {
		tk, err := l.read()
		if tk != nil {
			l.tokens = append(l.tokens, tk)
			l.state = Initial
		}

		if errors.Is(err, ErrBlankChar) {
			l.pos++
		}

		if errors.Is(err, ErrForbiddenChar) {
			return l.tokens, err
		}

		if errors.Is(err, ErrEOF) {
			break
		}
	}
	return l.tokens, nil
}

func (l *lexer) read() (token.Token, error) {
	if l.pos >= len(l.buffer) {
		return nil, ErrEOF
	}

	var err error
	var sb strings.Builder

	for l.pos < len(l.buffer) {
		b := l.buffer[l.pos]
		newState, err := l.state.take(b)
		if err != nil {
			break
		}

		sb.WriteByte(b)

		l.state = newState
		l.pos++
	}

	return token.NewToken(l.state, sb.String()), err
}
