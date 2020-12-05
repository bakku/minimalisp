package minimalisp

// And implements logical and for Minimalisp.
type And struct{}

// Arity returns infiniteArity for or.
func (f *And) Arity() int {
	return infiniteArity
}

// Call implements the and operation.
func (f *And) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<and> requires at least two arguments"}
	}

	for _, arg := range arguments {
		if !isTruthy(arg) {
			return false, nil
		}
	}

	return arguments[len(arguments)-1], nil
}

func (f *And) String() string {
	return "<and>"
}

// Or implements logical or for Minimalisp.
type Or struct{}

// Arity returns infiniteArity for or.
func (f *Or) Arity() int {
	return infiniteArity
}

// Call implements the or operation.
func (f *Or) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<or> requires at least two arguments"}
	}

	for _, arg := range arguments {
		if isTruthy(arg) {
			return arg, nil
		}
	}

	return false, nil
}

func (f *Or) String() string {
	return "<or>"
}

// Lt implements less than for Minimalisp.
type Lt struct{}

// Arity returns infiniteArity for less than.
func (f *Lt) Arity() int {
	return infiniteArity
}

// Call implements the less than operation.
func (f *Lt) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<<> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if num, ok := arg.(float64); ok {
			if num2, ok := arguments[i+1].(float64); ok {
				if !(num < num2) {
					return false, nil
				}
			} else {
				return nil, &executionError{line, "<<> is only defined for numbers"}
			}
		} else {
			return nil, &executionError{line, "<<> is only defined for numbers"}
		}
	}

	return true, nil
}

func (f *Lt) String() string {
	return "<<>"
}

// Lte implements less than equal for Minimalisp.
type Lte struct{}

// Arity returns infiniteArity for less than equal.
func (f *Lte) Arity() int {
	return infiniteArity
}

// Call implements the less than equal operation.
func (f *Lte) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<<=> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if num, ok := arg.(float64); ok {
			if num2, ok := arguments[i+1].(float64); ok {
				if !(num <= num2) {
					return false, nil
				}
			} else {
				return nil, &executionError{line, "<<=> is only defined for numbers"}
			}
		} else {
			return nil, &executionError{line, "<<=> is only defined for numbers"}
		}
	}

	return true, nil
}

func (f *Lte) String() string {
	return "<<=>"
}

// Gt implements greater than for Minimalisp.
type Gt struct{}

// Arity returns infiniteArity for greater than.
func (f *Gt) Arity() int {
	return infiniteArity
}

// Call implements the greater than operation.
func (f *Gt) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<>> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if num, ok := arg.(float64); ok {
			if num2, ok := arguments[i+1].(float64); ok {
				if !(num > num2) {
					return false, nil
				}
			} else {
				return nil, &executionError{line, "<>> is only defined for numbers"}
			}
		} else {
			return nil, &executionError{line, "<>> is only defined for numbers"}
		}
	}

	return true, nil
}

func (f *Gt) String() string {
	return "<>>"
}

// Gte implements greater than equal for Minimalisp.
type Gte struct{}

// Arity returns infiniteArity for greater than equal.
func (f *Gte) Arity() int {
	return infiniteArity
}

// Call implements the greater than equal operation.
func (f *Gte) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<>=> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if num, ok := arg.(float64); ok {
			if num2, ok := arguments[i+1].(float64); ok {
				if !(num >= num2) {
					return false, nil
				}
			} else {
				return nil, &executionError{line, "<>=> is only defined for numbers"}
			}
		} else {
			return nil, &executionError{line, "<>=> is only defined for numbers"}
		}
	}

	return true, nil
}

func (f *Gte) String() string {
	return "<>=>"
}

// Eq implements equal for Minimalisp.
type Eq struct{}

// Arity returns infiniteArity for equal.
func (f *Eq) Arity() int {
	return infiniteArity
}

// Call implements the equal operation.
func (f *Eq) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<=> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if arg != arguments[i+1] {
			return false, nil
		}
	}

	return true, nil
}

func (f *Eq) String() string {
	return "<=>"
}

// NotEq implements not equal for Minimalisp.
type NotEq struct{}

// Arity returns infiniteArity for not equal.
func (f *NotEq) Arity() int {
	return infiniteArity
}

// Call implements the not equal operation.
func (f *NotEq) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	if len(arguments) < 2 {
		return nil, &executionError{line, "<!=> requires at least two arguments"}
	}

	for i, arg := range arguments {
		if i == (len(arguments) - 1) {
			break
		}

		if arg == arguments[i+1] {
			return false, nil
		}
	}

	return true, nil
}

func (f *NotEq) String() string {
	return "<!=>"
}

// Not implements not for Minimalisp.
type Not struct{}

// Arity returns 1 for not.
func (f *Not) Arity() int {
	return 1
}

// Call implements the not operation.
func (f *Not) Call(line int, i *Interpreter, arguments []interface{}) (interface{}, error) {
	return !isTruthy(arguments[0]), nil
}

func (f *Not) String() string {
	return "<!>"
}
