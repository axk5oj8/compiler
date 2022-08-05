package lexer

import (
	"errors"
	"strings"

	"github.com/axk5oj8/compiler/internal/token"
)

func NewLexer(code string) Lexer {
	return &lexer{
		buffer: []byte(code),
	}
}

type Lexer interface {
	Tokenize() (token.Reader, error)
}

type lexer struct {
	buffer []byte
	tokens []token.Token

	pos   int
	state DFAState
}

func (l *lexer) Tokenize() (token.Reader, error) {
	for {
		tk, err := l.read()
		if tk != nil {
			l.tokens = append(l.tokens, tk)
			l.state = Initial
		}

		if err == nil {
			continue
		}

		if errors.Is(err, ErrBlankChar) {
			l.pos++
			continue
		}

		if errors.Is(err, ErrForbiddenChar) || (tk == nil && errors.Is(err, ErrInvalidChar)) {
			return nil, err
		}

		if errors.Is(err, ErrEOF) {
			return token.NewReader(l.tokens), nil
		}
	}
}

func (l *lexer) read() (token.Token, error) {
	if l.pos >= len(l.buffer) {
		return nil, ErrEOF
	}

	var err error
	var ns DFAState
	var sb strings.Builder

	for l.pos < len(l.buffer) {
		b := l.buffer[l.pos]
		ns, err = l.state.take(b)
		if err != nil {
			break
		}

		sb.WriteByte(b)

		l.state = ns
		l.pos++
	}

	if ttype, ok := typeMap[l.state]; ok {
		return token.NewToken(ttype, sb.String()), err
	}

	return nil, err
}
