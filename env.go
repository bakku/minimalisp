package tinylisp

import "fmt"

// Environment acts as a map to store and lookup values.
type Environment struct {
	values map[string]interface{}
}

// NewEnvironment is a factory function to create a new environment.
func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

// Define adds a new variable to the environment.
func (e *Environment) Define(token Token, value interface{}) error {
	_, ok := e.values[token.Lexeme]
	if ok {
		return &executionError{token.Line, fmt.Sprintf("Variable '%s' already defined", token.Lexeme)}
	}

	e.values[token.Lexeme] = value
	return nil
}

// Get returns a variable from the environment.
func (e *Environment) Get(token Token) (interface{}, error) {
	val, ok := e.values[token.Lexeme]
	if !ok {
		return nil, &executionError{token.Line, fmt.Sprintf("Undefined variable '%s'.", token.Lexeme)}
	}

	return val, nil
}
