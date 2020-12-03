package tinylisp

import "fmt"

type executionError struct {
	line int
	msg  string
}

func (e *executionError) Error() string {
	return fmt.Sprintf("[line %d] %s", e.line, e.msg)
}
