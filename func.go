package tinylisp

// infiniteArity is a constant which allows a function to have an infinite arity.
const infiniteArity = -1

// Function is a tiny lisp function
type Function interface {
	Arity() int
	Call(line int, interpreter *Interpreter, args []interface{}) (interface{}, error)
}

// TinyLispFunction is the standard function which is used in the tiny lisp interpreter.
type TinyLispFunction struct {
	name    string
	params  []Token
	body    Expression
	closure *Environment
}

// NewTinyLispFunction is a factory function to create a new function.
func NewTinyLispFunction(name string, params []Token, body Expression, closure *Environment) Function {
	return &TinyLispFunction{name, params, body, closure}
}

// Arity returns the amount of params which are expected for a function call.
func (f *TinyLispFunction) Arity() int {
	return len(f.params)
}

// Call calls the function.
func (f *TinyLispFunction) Call(line int, interpreter *Interpreter, args []interface{}) (interface{}, error) {
	env := NewEnvironmentWithEnclosing(f.closure)

	for i, p := range f.params {
		if err := env.Define(p, args[i]); err != nil {
			return nil, err
		}
	}

	return interpreter.execute(f.body, env)
}

func (f *TinyLispFunction) String() string {
	return "<" + f.name + ">"
}
