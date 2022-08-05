package lexer

func IsForbidden(b byte) bool {
	return false
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func IsBlank(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n'
}

func IsLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
