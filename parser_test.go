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

func TestParse_ShouldReturnCorrectExpressionsForDefuns(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Defun, "defun", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "say-hello", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "first", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "last", 1, nil},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "first", 1, nil},
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

	_, ok := expressions[0].(*tinylisp.DefunExpr)
	if !ok {
		t.Fatalf("Expected defun expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForNestedCalls(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "first", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "rest", 1, nil},
		tinylisp.Token{tinylisp.Quote, "'", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Number, "1", 1, 1},
		tinylisp.Token{tinylisp.Number, "2", 1, 2},
		tinylisp.Token{tinylisp.Number, "3", 1, 3},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
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

	_, ok := expressions[0].(*tinylisp.FuncCallExpr)
	if !ok {
		t.Fatalf("Expected function call expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForLet(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Let, "let", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "n", 1, nil},
		tinylisp.Token{tinylisp.Number, "1", 1, 1},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "+", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "n", 1, nil},
		tinylisp.Token{tinylisp.Number, "1", 1, 1},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
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

	_, ok := expressions[0].(*tinylisp.LetExpr)
	if !ok {
		t.Fatalf("Expected let expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForLambda(t *testing.T) {
	tokens := []tinylisp.Token{
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Defvar, "defvar", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "f", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Lambda, "lambda", 1, 1},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "name", 1, nil},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.LeftParen, "(", 1, nil},
		tinylisp.Token{tinylisp.Identifier, "println", 1, nil},
		tinylisp.Token{tinylisp.Str, "\"Hello\"", 1, "Hello"},
		tinylisp.Token{tinylisp.Identifier, "name", 1, nil},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
		tinylisp.Token{tinylisp.RightParen, ")", 1, nil},
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
