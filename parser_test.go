package minimalisp_test

import (
	"testing"

	. "bakku.dev/minimalisp"
)

func TestParse_ShouldReturnCorrectExpressionsForLiterals(t *testing.T) {
	tokens := []Token{
		Token{False, "false", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*LiteralExpr)
	if !ok {
		t.Fatalf("Expected literal expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForDefvars(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{Defvar, "defvar", 1, nil},
		Token{Identifier, "hello", 1, nil},
		Token{Str, "\"hello\"", 1, "hello"},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*DefvarExpr)
	if !ok {
		t.Fatalf("Expected defvar expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForIfs(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{If, "if", 1, nil},
		Token{True, "true", 1, nil},
		Token{Str, "\"yes\"", 1, "yes"},
		Token{Str, "\"no\"", 1, "no"},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*IfExpr)
	if !ok {
		t.Fatalf("Expected if expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForDefuns(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{Defun, "defun", 1, nil},
		Token{Identifier, "say-hello", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "first", 1, nil},
		Token{Identifier, "last", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{Identifier, "first", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*DefunExpr)
	if !ok {
		t.Fatalf("Expected defun expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForNestedCalls(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "first", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "rest", 1, nil},
		Token{Quote, "'", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Number, "1", 1, 1},
		Token{Number, "2", 1, 2},
		Token{Number, "3", 1, 3},
		Token{RightParen, ")", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*FuncCallExpr)
	if !ok {
		t.Fatalf("Expected function call expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForLet(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{Let, "let", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "n", 1, nil},
		Token{Number, "1", 1, 1},
		Token{RightParen, ")", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "+", 1, nil},
		Token{Identifier, "n", 1, nil},
		Token{Number, "1", 1, 1},
		Token{RightParen, ")", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*LetExpr)
	if !ok {
		t.Fatalf("Expected let expression")
	}
}

func TestParse_ShouldReturnCorrectExpressionsForLambda(t *testing.T) {
	tokens := []Token{
		Token{LeftParen, "(", 1, nil},
		Token{Defvar, "defvar", 1, nil},
		Token{Identifier, "f", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Lambda, "lambda", 1, 1},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "name", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{LeftParen, "(", 1, nil},
		Token{Identifier, "println", 1, nil},
		Token{Str, "\"Hello\"", 1, "Hello"},
		Token{Identifier, "name", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{RightParen, ")", 1, nil},
		Token{EOF, "", 1, nil},
	}

	parser := NewParser(tokens)
	expressions, err := parser.Parse()

	if err != nil {
		t.Fatalf("Expected err to be nil, got %v", err)
	}

	if len(expressions) != 1 {
		t.Fatalf("Expected %d expressions, got %d", 1, len(expressions))
	}

	_, ok := expressions[0].(*DefvarExpr)
	if !ok {
		t.Fatalf("Expected defvar expression")
	}
}
