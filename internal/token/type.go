package token

const (
	Plus       Type = iota // +
	Minus                  // -
	Asterisk               // *
	Slash                  // /
	GE                     // >=
	GT                     // >
	EQ                     // ==
	LE                     // <=
	LT                     // <
	SemiColon              // ;
	LeftParen              // (
	RightParen             // )
	Assignment             // =
	If
	Else
	Int
	Identifier
	IntLiteral
	StringLiteral
)

type Type uint

func (t Type) String() string {
	switch t {
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Asterisk:
		return "Asterisk"
	case Slash:
		return "Slash"
	case GE:
		return "GE"
	case GT:
		return "GT"
	case EQ:
		return "EQ"
	case LE:
		return "LE"
	case LT:
		return "LT"
	case SemiColon:
		return "SemiColon"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Assignment:
		return "Assignment"
	case If:
		return "If"
	case Else:
		return "Else"
	case Int:
		return "Int"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	case StringLiteral:
		return "StringLiteral"
	default:
		return "Unknown"
	}
}
