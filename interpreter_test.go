package tinylisp_test

import (
	"testing"

	"bakku.dev/tinylisp"
)

func TestInterpret_ShouldCorrectlyInterpretCode1(t *testing.T) {
	expressions := []tinylisp.Expression{
		&tinylisp.DefvarExpr{tinylisp.Token{tinylisp.Identifier, "name", 1, nil}, &tinylisp.LiteralExpr{"Steven"}},
		&tinylisp.VarExpr{tinylisp.Token{tinylisp.Identifier, "name", 2, nil}},
	}

	interpreter := tinylisp.NewInterpreter()
	ret, err := interpreter.Interpret(expressions)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ret != "Steven" {
		t.Fatalf("Expected 'Steven' as result, got '%v'", ret)
	}
}
