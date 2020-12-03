package tinylisp

import "fmt"

// Parser generates the AST using the tokens from the scanner.
type Parser struct {
	tokens []Token
	curr   int

	expressions []Expression
}

// NewParser is a factory function to create a new Parser.
func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens}
}

// Parse parses all tokens and returns a list of expressions.
func (p *Parser) Parse() ([]Expression, error) {
	for !p.isAtEnd() {
		expr, err := p.declaration()
		if err != nil {
			return nil, err
		}

		p.expressions = append(p.expressions, expr)
	}

	return p.expressions, nil
}

func (p *Parser) declaration() (Expression, error) {
	if p.match(LeftParen) {
		if p.matchN(Defvar, 1) {
			return p.varDef()
		}
	}

	return p.primary()
}

func (p *Parser) varDef() (Expression, error) {
	if _, err := p.consume(LeftParen, "Expect '(' before variable definition"); err != nil {
		return nil, err
	}

	if _, err := p.consume(Defvar, "Expect identifier after '('"); err != nil {
		return nil, err
	}

	ident, err := p.consume(Identifier, "Expect identifier after 'defvar'")
	if err != nil {
		return nil, err
	}

	expr, err := p.expression()
	if err != nil {
		return nil, err
	}

	if _, err := p.consume(RightParen, "Expect ')' after expression"); err != nil {
		return nil, err
	}

	return &DefvarExpr{ident, expr}, nil
}

func (p *Parser) expression() (Expression, error) {
	return p.primary()
}

func (p *Parser) primary() (Expression, error) {
	if p.match(False) {
		p.curr++
		return &LiteralExpr{false}, nil
	} else if p.match(True) {
		p.curr++
		return &LiteralExpr{true}, nil
	} else if p.match(Str) {
		p.curr++
		return &LiteralExpr{p.peekN(-1).Value}, nil
	} else if p.match(Number) {
		p.curr++
		return &LiteralExpr{p.peekN(-1).Value}, nil
	} else if p.match(Nil) {
		p.curr++
		return &LiteralExpr{nil}, nil
	} else if p.match(Identifier) {
		p.curr++
		return &VarExpr{p.peekN(-1)}, nil
	}

	return nil, &executionError{p.peek().Line, fmt.Sprintf("Expression expected.")}
}

func (p *Parser) peek() Token {
	return p.tokens[p.curr]
}

func (p *Parser) peekN(n int) Token {
	if p.curr+n >= len(p.tokens) {
		return Token{EOF, "", 0, nil}
	}

	return p.tokens[p.curr+n]
}

func (p *Parser) match(tokenType int) bool {
	if p.isAtEnd() || (p.peek().TokenType != tokenType) {
		return false
	}

	return true
}

func (p *Parser) matchN(tokenType, n int) bool {
	return p.peekN(n).TokenType == tokenType
}

func (p *Parser) consume(tokenType int, msg string) (Token, error) {
	if !p.match(tokenType) {
		return Token{}, &executionError{p.peek().Line, msg}
	}

	ret := p.peek()
	p.curr++

	return ret, nil
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == EOF
}
