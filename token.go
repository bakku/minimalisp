package tinylisp

const (
	LeftParen = iota
	RightParen
	Semicolon
	Quote
	Identifier
	Str
	Number
	Lambda
	Defvar
	Defun
	If
	Let
	Nil
	EOF
)

var keywords = map[string]int{
	"lamdba": Lambda,
	"let":    Let,
	"defvar": Defvar,
	"defun":  Defun,
	"if":     If,
	"nil":    Nil,
}

// Token represents a certain token at a specific location
// in a piece of source code.
type Token struct {
	TokenType int
	Lexeme    string
	Line      int
	value     interface{}
}
