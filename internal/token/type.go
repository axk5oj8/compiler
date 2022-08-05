package token

const (
	Plus Type = iota
	GE
	GT
	EQ
	Identifier
)

type Type uint

func (t Type) String() string {
	switch t {
	case Plus:
		return "Plus"
	case GE:
		return "GE"
	case GT:
		return "GT"
	case EQ:
		return "EQ"
	case Identifier:
		return "Identifier"
	default:
		return "Unknown"
	}
}
