package lexer

const (
	Initial DFAState = iota
	ID
	IDi
	IDin
	IDint
	GT
	LT
	GE
	LE
	Plus
	Minus
	Star
	Slash
	Assignment
	LeftParen
	RightParen
	IntLiteral

	Undefined
)

type DFAState uint

type Transform func(b byte) DFAState

var transMap = map[DFAState]Transform{
	Initial: fromInitial,
}

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
	case IsLetter(b):
		return ID
	case b == '+':
		return Plus
	case b == '-':
		return Minus
	case b == '*':
		return Star
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
	default:
		return Undefined
	}
}

func fromID(b byte) DFAState {
	switch {
	case IsDigit(b):
		return ID
	case IsLetter():

	default:
		return Undefined
	}
}
