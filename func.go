package minimalisp

// infiniteArity is a constant which allows a function to have an infinite arity.
const infiniteArity = -1

// Function is a minimalisp function
type Function interface {
	Arity() int
	Call(line int, interpreter *Interpreter, args []interface{}) (interface{}, error)
}

// MinimalispFunction is the standard function which is used in the minimalisp interpreter.
type MinimalispFunction struct {
	name    string
	params  []Token
	body    Expression
	closure *Environment
}

// NewMinimalispFunction is a factory function to create a new function.
func NewMinimalispFunction(name string, params []Token, body Expression, closure *Environment) Function {
	return &MinimalispFunction{name, params, body, closure}
}

// Arity returns the amount of params which are expected for a function call.
func (f *MinimalispFunction) Arity() int {
	return len(f.params)
}

// Call calls the function.
func (f *MinimalispFunction) Call(line int, interpreter *Interpreter, args []interface{}) (interface{}, error) {
	env := NewEnvironmentWithEnclosing(f.closure)

	for i, p := range f.params {
		if err := env.Define(p, args[i]); err != nil {
			return nil, err
		}
	}

	return interpreter.execute(f.body, env)
}

func (f *MinimalispFunction) String() string {
	return "<" + f.name + ">"
}
