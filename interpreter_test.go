package tinylisp_test

import (
	"testing"

	. "bakku.dev/tinylisp"
)

func TestInterpret_ShouldCorrectlyInterpretCode1(t *testing.T) {
	expressions := []Expression{
		&DefvarExpr{Token{Identifier, "name", 1, nil}, &LiteralExpr{"Steven"}},
		&VarExpr{Token{Identifier, "name", 2, nil}},
	}

	interpreter := NewInterpreter()
	ret, err := interpreter.Interpret(expressions)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ret != "Steven" {
		t.Fatalf("Expected 'Steven' as result, got '%v'", ret)
	}
}

func TestInterpret_ShouldCorrectlyInterpretCode2(t *testing.T) {
	expressions := []Expression{
		&IfExpr{
			&LiteralExpr{nil},
			&LiteralExpr{"yes"},
			&LiteralExpr{"no"},
		},
	}

	interpreter := NewInterpreter()
	ret, err := interpreter.Interpret(expressions)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ret != "no" {
		t.Fatalf("Expected 'no' as result, got '%v'", ret)
	}
}

func TestInterpret_ShouldCorrectlyInterpretCode3(t *testing.T) {
	expressions := []Expression{
		&DefvarExpr{
			Token{Identifier, "outer-name", 1, nil},
			&LiteralExpr{"Steven"},
		},
		&DefunExpr{
			Token{Identifier, "give-outer", 2, nil},
			[]Token{},
			&VarExpr{Token{Identifier, "outer-name", 2, nil}},
		},
		&FuncCallExpr{
			Token{Identifier, "give-outer", 3, nil},
			[]Expression{},
		},
	}

	interpreter := NewInterpreter()
	ret, err := interpreter.Interpret(expressions)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ret != "Steven" {
		t.Fatalf("Expected 'no' as result, got '%v'", ret)
	}
}
