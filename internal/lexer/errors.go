package lexer

import "errors"

var (
	ErrForbiddenChar = errors.New("character is forbidden")
	ErrInvalidChar   = errors.New("character is invalid for current state")
	ErrEOF           = errors.New("end of file")
	ErrBlankChar     = errors.New("character is blank")
)
