// Package defines a parse tree intended to be created during parsing by a basic parser
package tree

import (
	"./../scanner/scanner"
)

type Node struct {
	Type int
	Children []*node
	Token *scanner.Token
}

func NewNode(term bool, tok *scanner.Token) *node {
	return &node{term, []*node{}, tok}
}

func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) eval() {
	if n.Type == NON_TERMINAL {
		return n.Token
	}
	
}
