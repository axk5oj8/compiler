package ast

import "fmt"

func NewNode(ntype NodeType, text string) Node {
	return &node{
		ntype: ntype,
		text:  text,
	}
}

type Node interface {
	Parent() Node
	SetParent(Node)

	Children() []Node
	AddChild(Node)

	Type() NodeType

	Text() string
}

type node struct {
	parent   Node
	children []Node

	ntype NodeType
	text  string
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) SetParent(p Node) {
	n.parent = p
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) Type() NodeType {
	return n.ntype
}

func (n *node) Text() string {
	return n.text
}

func (n *node) AddChild(child Node) {
	n.children = append(n.children, child)
	child.SetParent(n)
}

func Dump(node Node, indent string) {
	fmt.Println(indent + node.Type().String() + " " + node.Text())
	for _, child := range node.Children() {
		Dump(child, indent+"\t")
	}
}
