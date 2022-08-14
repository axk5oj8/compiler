package lexer

import "github.com/axk5oj8/compiler/internal/token"

const (
	Initial DFAState = iota
	ID
	IDi
	IDin
	IDint
	IDif
	IDe
	IDel
	IDels
	IDelse
	GT
	LT
	GE
	LE
	EQ
	Plus
	Minus
	Asterisk
	Slash
	Assignment
	LeftParen
	RightParen
	SemiColon
	IntLiteral

	Undefined
)

var (
	typeMap = map[DFAState]token.Type{
		Plus:       token.Plus,
		Minus:      token.Minus,
		Asterisk:   token.Asterisk,
		Slash:      token.Slash,
		GE:         token.GE,
		GT:         token.GT,
		EQ:         token.EQ,
		LE:         token.LE,
		SemiColon:  token.SemiColon,
		LeftParen:  token.LeftParen,
		RightParen: token.RightParen,
		IDif:       token.If,
		IDelse:     token.Else,
		IDint:      token.Int,
		ID:         token.Identifier,
		IntLiteral: token.IntLiteral,
		IDi:        token.Identifier,
		IDin:       token.Identifier,
		IDe:        token.Identifier,
		IDel:       token.Identifier,
		IDels:      token.Identifier,
		Assignment: token.Assignment,
	}

	transMap = map[DFAState]Transform{
		Initial:    fromInitial,
		ID:         fromID,
		IDi:        fromIDi,
		IDin:       fromIDin,
		IDint:      fromIDint,
		IDif:       fromIDif,
		IDe:        fromIDe,
		IDel:       fromIDel,
		IDels:      fromIDels,
		IDelse:     fromIDelse,
		GT:         fromGT,
		LT:         fromLT,
		LE:         fromLE,
		GE:         fromGE,
		EQ:         fromEQ,
		Plus:       fromPlus,
		Minus:      fromMinus,
		Asterisk:   fromAsterisk,
		Slash:      fromSlash,
		Assignment: fromAssignment,
		LeftParen:  fromLeftParen,
		RightParen: fromRightParen,
		SemiColon:  fromSemiColon,
		IntLiteral: fromIntLiteral,
	}
)

type DFAState uint

type Transform func(b byte) DFAState

func (s DFAState) take(b byte) (DFAState, error) {
	if IsForbidden(b) {
		return s, ErrForbiddenChar
	}

	if IsBlank(b) {
		return s, ErrBlankChar
	}

	if ns := transMap[s](b); ns != Undefined {
		return ns, nil
	}

	return s, ErrInvalidChar
}

func fromInitial(b byte) DFAState {
	switch {
	case IsDigit(b):
		return IntLiteral
	case b == 'i':
		return IDi
	case b == 'e':
		return IDe
	case IsLetter(b):
		return ID
	case b == '+':
		return Plus
	case b == '-':
		return Minus
	case b == '*':
		return Asterisk
	case b == '/':
		return Slash
	case b == '=':
		return Assignment
	case b == '>':
		return GT
	case b == '<':
		return LT
	case b == '(':
		return LeftParen
	case b == ')':
		return RightParen
	case b == ';':
		return SemiColon
	default:
		return Undefined
	}
}

func fromID(b byte) DFAState {
	if IsDigit(b) || IsLetter(b) {
		return ID
	}
	return Undefined
}

func fromIDi(b byte) DFAState {
	if b == 'n' {
		return IDin
	}

	if b == 'f' {
		return IDif
	}

	return fromID(b)
}

func fromIDin(b byte) DFAState {
	if b == 't' {
		return IDint
	}
	return fromID(b)
}

func fromIDint(b byte) DFAState {
	return fromID(b)
}

func fromIDif(b byte) DFAState {
	if b == 'n' {
		return IDin
	}
	return Undefined
}

func fromIDe(b byte) DFAState {
	if b == 'l' {
		return IDel
	}
	return fromID(b)
}

func fromIDel(b byte) DFAState {
	if b == 's' {
		return IDels
	}
	return fromID(b)
}

func fromIDels(b byte) DFAState {
	if b == 'e' {
		return IDelse
	}
	return fromID(b)
}
func fromIDelse(b byte) DFAState {
	return fromID(b)
}

func fromGT(b byte) DFAState {
	if b == '=' {
		return GE
	}
	return Undefined
}

func fromGE(b byte) DFAState {
	return Undefined
}
func fromLT(b byte) DFAState {
	if b == '=' {
		return LE
	}
	return Undefined
}
func fromLE(b byte) DFAState {
	return Undefined
}
func fromEQ(b byte) DFAState {
	return Undefined
}
func fromPlus(b byte) DFAState {
	return Undefined
}
func fromMinus(b byte) DFAState {
	return Undefined
}
func fromAsterisk(b byte) DFAState {
	return Undefined
}
func fromSlash(b byte) DFAState {
	return Undefined
}

func fromAssignment(b byte) DFAState {
	if b == '=' {
		return EQ
	}
	return Undefined
}

func fromLeftParen(b byte) DFAState {
	return Undefined
}

func fromRightParen(b byte) DFAState {
	return Undefined
}

func fromSemiColon(b byte) DFAState {
	return Undefined
}

func fromIntLiteral(b byte) DFAState {
	if IsDigit(b) {
		return IntLiteral
	}
	return Undefined
}
