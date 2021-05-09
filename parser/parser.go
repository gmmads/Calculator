package parser

import (
	"github.com/gmmads/Calculator/lexer"
)

type AstNode struct {
	NodeType string
	Token    lexer.Token
	Children []*AstNode
}

type Parser interface {
	Parse(tokens []lexer.Token) (*AstNode, error)
}
