package tinylisp

// First returns the first element of a list.
type First struct{}

// Arity returns 1.
func (f *First) Arity() int {
	return 1
}

// Call implements the extraction of the first element.
func (f *First) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	list, ok := arguments[0].(List)
	if !ok {
		return nil, &executionError{line, "'first' is only defined for lists"}
	}

	if list.Len() == 0 {
		return nil, nil
	}

	return list.First(), nil
}

func (f *First) String() string {
	return "<first>"
}

// Rest returns all except the first element of a list.
type Rest struct{}

// Arity returns 1.
func (f *Rest) Arity() int {
	return 1
}

// Call implements returning all except the first element of a list.
func (f *Rest) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	list, ok := arguments[0].(List)
	if !ok {
		return nil, &executionError{line, "'rest' is only defined for lists"}
	}

	if list.Len() == 0 {
		return nil, nil
	}

	return list.Rest(), nil
}

func (f *Rest) String() string {
	return "<rest>"
}

// Add adds an element to a list and returns a new list.
type Add struct{}

// Arity returns 2.
func (f *Add) Arity() int {
	return 2
}

// Call implements the addition of an element to a list.
func (f *Add) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	list, ok := arguments[0].(List)
	if !ok {
		return nil, &executionError{line, "'len' is only defined for lists"}
	}

	return list.Add(arguments[1]), nil
}

func (f *Add) String() string {
	return "<add>"
}

// Len returns the amount of elements in a list.
type Len struct{}

// Arity returns 1.
func (f *Len) Arity() int {
	return 1
}

// Call implements the counting of the elements in a list.
func (f *Len) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	list, ok := arguments[0].(List)
	if !ok {
		return nil, &executionError{line, "'len' is only defined for lists"}
	}

	return list.Len(), nil
}

func (f *Len) String() string {
	return "<len>"
}
