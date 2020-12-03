package tinylisp

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// Scanner is responsible for scanning TinyLisp source code and returning
// it as a list of tokens.
type Scanner struct {
	// start marks where the current lexeme starts.
	start int
	// end marks where the current lexeme ends.
	end int

	line   int
	src    string
	tokens []Token

	// to specify where errors are written.
	out io.Writer
}

// NewScanner is a factory function to create a new Scanner.
func NewScanner(src string, out io.Writer) *Scanner {
	return &Scanner{
		src:  src,
		line: 1,
		out:  out,
	}
}

// Scan scans the source code and returns a list of tokens.
func (s *Scanner) Scan() ([]Token, bool) {
	var ok bool = true

	for !s.isAtEnd() {
		err := s.nextToken()
		if err != nil {
			// Ignore errors returned by Fprintf.
			_, _ = fmt.Fprintf(s.out, "%v\n", err)
			ok = false
		}

		s.end++
		s.start = s.end
	}

	s.tokens = append(s.tokens, Token{EOF, "", s.line, nil})

	return s.tokens, ok
}

func (s *Scanner) nextToken() error {
	c := string(s.src[s.end])

	switch c {
	case "(":
		s.tokens = append(s.tokens, Token{LeftParen, s.src[s.start : s.end+1], s.line, nil})
		return nil
	case ")":
		s.tokens = append(s.tokens, Token{RightParen, s.src[s.start : s.end+1], s.line, nil})
		return nil
	case ";":
		s.end++

		for s.peek() != "\n" && !s.isAtEnd() {
			s.end++
		}

		return nil
	case "'":
		s.tokens = append(s.tokens, Token{Quote, s.src[s.start : s.end+1], s.line, nil})
		return nil
	case " ":
		return nil
	case "\t":
		return nil
	case "\r":
		return nil
	case "\n":
		s.line++
		return nil
	case "\"":
		if err := s.string(); err != nil {
			return err
		}

		return nil
	default:
		if isDigit(c) {
			if err := s.number(); err != nil {
				return err
			}

			return nil
		} else if isAlpha(c) {
			if err := s.identifier(); err != nil {
				return err
			}

			return nil
		} else {
			return &executionError{s.line, fmt.Sprintf("Unexpected character: %s", c)}
		}
	}
}

func (s *Scanner) string() error {
	s.end++

	for s.peek() != "\"" && !s.isAtEnd() {
		s.end++
	}

	if s.isAtEnd() {
		return &executionError{s.line, "Unterminated string"}
	}

	s.tokens = append(s.tokens, Token{Str, s.src[s.start : s.end+1], s.line, s.src[s.start+1 : s.end]})

	return nil
}

func (s *Scanner) number() error {
	for isDigit(s.peekN(1)) {
		s.end++
	}

	if s.peekN(1) == "." && isDigit(s.peekN(2)) {
		s.end++

		for isDigit(s.peekN(1)) {
			s.end++
		}
	}

	num, err := strconv.ParseFloat(s.src[s.start:s.end+1], 64)

	if err != nil {
		return &executionError{s.line, fmt.Sprintf("error while parsing float: %v", err)}
	}

	s.tokens = append(s.tokens, Token{Number, s.src[s.start : s.end+1], s.line, num})

	return nil
}

func (s *Scanner) identifier() error {
	for isAlphaNumeric(s.peekN(1)) {
		s.end++
	}

	text := s.src[s.start : s.end+1]

	tokenType, ok := keywords[text]

	if ok {
		s.tokens = append(s.tokens, Token{tokenType, s.src[s.start : s.end+1], s.line, nil})
	} else {
		s.tokens = append(s.tokens, Token{Identifier, s.src[s.start : s.end+1], s.line, nil})
	}

	return nil
}

func (s *Scanner) peek() string {
	return string(s.src[s.end])
}

func (s *Scanner) peekN(n int) string {
	if s.end+n >= len(s.src) {
		return " "
	}

	return string(s.src[s.end+n])
}

func isAlphaNumeric(c string) bool {
	return isDigit(c) || isAlpha(c)
}

func isDigit(c string) bool {
	matched, err := regexp.MatchString("[0-9]", c)
	if err != nil {
		return false
	}

	return matched
}

func isAlpha(c string) bool {
	matched, err := regexp.MatchString("[a-zA-Z\\+-<>!=/*%_?]", c)
	if err != nil {
		return false
	}

	return matched
}

func (s *Scanner) isAtEnd() bool {
	return s.end >= len(s.src)
}
