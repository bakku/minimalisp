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

// Map applies a function to each element of a list.
type Map struct{}

// Arity returns 2.
func (f *Map) Arity() int {
	return 2
}

// Call implements map for a list.
func (f *Map) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	fun, ok := arguments[0].(Function)
	if !ok {
		return nil, &executionError{line, "<map> expects a function as first parameter"}
	}

	if fun.Arity() != 1 {
		return nil, &executionError{line, "<map> expects a function which accepts one argument"}
	}

	list, ok := arguments[1].(List)
	if !ok {
		return nil, &executionError{line, "<map> expects a list as second parameter"}
	}

	var mappedElements []interface{}

	restOfList := list

	for restOfList.Len() > 0 {
		var args []interface{}
		args = append(args, restOfList.First())

		newEl, err := fun.Call(line, i, args)
		if err != nil {
			return nil, err
		}

		mappedElements = append(mappedElements, newEl)

		restOfList = restOfList.Rest()
	}

	return NewArrayList(mappedElements), nil
}

func (f *Map) String() string {
	return "<map>"
}

// Filter returns a new list with only the elements for which the given function returned true.
type Filter struct{}

// Arity returns 2.
func (f *Filter) Arity() int {
	return 2
}

// Call implements filter for a list.
func (f *Filter) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	fun, ok := arguments[0].(Function)
	if !ok {
		return nil, &executionError{line, "<filter> expects a function as first parameter"}
	}

	if fun.Arity() != 1 {
		return nil, &executionError{line, "<filter> expects a function which accepts one argument"}
	}

	list, ok := arguments[1].(List)
	if !ok {
		return nil, &executionError{line, "<filter> expects a list as second parameter"}
	}

	var filteredElements []interface{}

	restOfList := list

	for restOfList.Len() > 0 {
		var args []interface{}
		args = append(args, restOfList.First())

		response, err := fun.Call(line, i, args)
		if err != nil {
			return nil, err
		}

		if isTruthy(response) {
			filteredElements = append(filteredElements, restOfList.First())
		}

		restOfList = restOfList.Rest()
	}

	return NewArrayList(filteredElements), nil
}

func (f *Filter) String() string {
	return "<filter>"
}
