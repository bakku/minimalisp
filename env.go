package minimalisp

import "fmt"

// Environment acts as a map to store and lookup values.
type Environment struct {
	values    map[string]interface{}
	enclosing *Environment
}

// NewEnvironment is a factory function to create a new environment.
func NewEnvironment() *Environment {
	return &Environment{
		values: make(map[string]interface{}),
	}
}

// NewEnvironmentWithEnclosing create a new environment with a parent environment.
func NewEnvironmentWithEnclosing(enclosing *Environment) *Environment {
	return &Environment{
		values:    make(map[string]interface{}),
		enclosing: enclosing,
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
		if e.enclosing != nil {
			return e.enclosing.Get(token)
		}

		return nil, &executionError{token.Line, fmt.Sprintf("Undefined variable '%s'.", token.Lexeme)}
	}

	return val, nil
}
