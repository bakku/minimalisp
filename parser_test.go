package tinylisp_test

import (
	"testing"

	"bakku.dev/tinylisp"
)

func TestParse_ShouldReturnCorrectExpressionsForLiterals(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.False, "false", 1, nil},
		tinylisp.Token{tinylisp.EOF, "", 1, nil},
	}

	parser := tinylisp.NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*tinylisp.LiteralExpr)
	if !ok {
		t.Fatalf("Expected literal expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForDefvars(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Defvar, "defvar", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "hello", 1, nil},
		tinylisp.Token{tinylisp.Str, "\"hello\"", 1, "hello"},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.EOF, "", 1, nil},
	}

	parser := tinylisp.NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*tinylisp.DefvarExpr)
	if !ok {
		t.Fatalf("Expected defvar expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForIfs(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.If, "if", 1, nil},
		tinylisp.Token{tinylisp.True, "true", 1, nil},
		tinylisp.Token{tinylisp.Str, "\"yes\"", 1, "yes"},
		tinylisp.Token{tinylisp.Str, "\"no\"", 1, "no"},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.EOF, "", 1, nil},
	}

	parser := tinylisp.NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*tinylisp.IfExpr)
	if !ok {
		t.Fatalf("Expected if expression")
	}
}
