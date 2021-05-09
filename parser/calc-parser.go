package parser

import (
	"errors"

	"github.com/gmmads/Calculator/constants"
	"github.com/gmmads/Calculator/lexer"
)

type CalcParser struct{}

// AST-node types
const (
	BINOP  = "BINOP"
	NUMBER = "NUM"
	UMINUS = "UMINUS"
)

func NewCalcParser() Parser {
	return &CalcParser{}
}

var (
	i      int
	tokens []lexer.Token
)

// Returns the root of the tree
func (*CalcParser) Parse(tokenizedExpr []lexer.Token) (*AstNode, error) {
	i = 0
	tokens = tokenizedExpr
	root, err := expr()
	if err != nil {
		return nil, err
	}
	if i < len(tokens) {
		return nil, errors.New("parsing error: Failed to parse entire expression")
	}
	return root, nil
}

func eat(tokenType string) error {
	if i >= len(tokens) {
		return errors.New("parsing error: Reached end of expression without finishing parsing")
	}
	if tokens[i].TokenType == tokenType {
		i++
		return nil
	}
	return errors.New("parsing error: Unexpected token: " + tokenType)
}

func factor() (*AstNode, error) {
	token := tokens[i]
	if token.TokenType == constants.NUMBER {
		err := eat(constants.NUMBER)
		if err != nil {
			return nil, err
		}

		return &AstNode{NodeType: constants.NUMBER, Token: token}, nil
	} else if token.TokenType == constants.LPAREN {
		err1 := eat(constants.LPAREN)
		if err1 != nil {
			return nil, err1
		}

		node, err2 := expr()
		if err2 != nil {
			return nil, err2
		}

		err3 := eat(constants.RPAREN)
		if err3 != nil {
			return nil, err3
		}

		return node, nil
	} else if token.TokenType == constants.MINUS {
		err1 := eat(constants.MINUS)
		if err1 != nil {
			return nil, err1
		}

		child, err2 := expr()
		if err2 != nil {
			return nil, err2
		}
		return &AstNode{NodeType: constants.UMINUS, Token: token, Children: []*AstNode{child}}, nil
	}
	return nil, errors.New("parsing error: Unexpected token: " + token.TokenType)
}

func term() (*AstNode, error) {
	node, err := factor()
	if err != nil {
		return nil, err
	}

	for i < len(tokens) &&
		(tokens[i].TokenType == constants.MULT || tokens[i].TokenType == constants.DIV) {
		token := tokens[i]
		err1 := eat(token.TokenType)
		if err1 != nil {
			return nil, err1
		}

		var children []*AstNode
		children = append(children, node)

		right_child, err2 := factor()
		if err2 != nil {
			return nil, err2
		}

		children = append(children, right_child)
		node = &AstNode{NodeType: constants.BINOP, Token: token, Children: children}
	}

	return node, nil
}

func expr() (*AstNode, error) {
	node, err := term()
	if err != nil {
		return nil, err
	}

	for i < len(tokens) &&
		(tokens[i].TokenType == constants.PLUS || tokens[i].TokenType == constants.MINUS) {
		token := tokens[i]
		err1 := eat(token.TokenType)
		if err1 != nil {
			return nil, err1
		}

		var children []*AstNode
		children = append(children, node)

		right_child, err2 := term()
		if err2 != nil {
			return nil, err2
		}

		children = append(children, right_child)
		node = &AstNode{NodeType: constants.BINOP, Token: token, Children: children}
	}

	return node, nil
}
