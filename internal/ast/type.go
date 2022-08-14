package ast

import "fmt"

const (
	Program NodeType = iota

	IntDeclaration
	ExpressionStmt
	AssignmentStmt

	Primary
	Multiplicative
	Additive

	Identifier
	IntLiteral
)

type NodeType int

func (t NodeType) String() string {
	switch t {
	case Program:
		return "Program"
	case IntDeclaration:
		return "IntDeclaration"
	case ExpressionStmt:
		return "ExpressionStmt"
	case AssignmentStmt:
		return "AssignmentStmt"
	case Primary:
		return "Primary"
	case Multiplicative:
		return "Multiplicative"
	case Additive:
		return "Additive"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	default:
		return fmt.Sprintf("Undefined type: %d", t)
	}
}
