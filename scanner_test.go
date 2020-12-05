package minimalisp_test

import (
	"bytes"
	"fmt"
	"testing"

	. "bakku.dev/minimalisp"
)

func TestScanSourceCode_ShouldReturnEOFForEmptyString(t *testing.T) {
	var buf bytes.Buffer
	scanner := NewScanner("", &buf)
	tokens, ok := scanner.Scan()

	if buf.String() != "" {
		t.Fatalf("Expected error messages to be empty, got %s", buf.String())
	}

	if !ok {
		t.Fatalf("Expected everything to be ok")
	}

	if len(tokens) != 1 {
		t.Fatalf("Expected token list size 1, got %v", len(tokens))
	}

	token := tokens[0]

	if token.TokenType != EOF {
		t.Fatalf("Expected EOF token, got %v", token.TokenType)
	}
}

func TestScanSourceCode_ShouldReturnCorrectSequenceOfTokens(t *testing.T) {
	sourceCode := `
        ; some comment
        (defun say-hello (name)
          (if name
            (let (hello-name (str "Hello, " name))
              (println hello-name))
            (println "No name given"))) 

        (defvar name "Steven")

        (println '(123.456 "two" say-hello))
        `

	var buf bytes.Buffer
	scanner := NewScanner(sourceCode, &buf)
	tokens, ok := scanner.Scan()

	if buf.String() != "" {
		fmt.Println(tokens)
		t.Fatalf("Expected error messages to be empty, got %s", buf.String())
	}

	if !ok {
		t.Fatalf("Expected everything to be ok")
	}

	if len(tokens) != 45 {
		t.Fatalf("Expected token list size 45, got %v", len(tokens))
	}
}
