package tinylisp

// Interpreter is an implementation of the AST visitor.
type Interpreter struct {
	globals *Environment
	current *Environment
}

// NewInterpreter is a factory function to create a new Interpreter.
func NewInterpreter() *Interpreter {
	global := NewEnvironment()
	return &Interpreter{global, global}
}

// Interpret takes a slice of expressions and interprets them.
func (i *Interpreter) Interpret(expressions []Expression) (interface{}, error) {
	var ret interface{}
	var err error = nil

	for _, expr := range expressions {
		if ret, err = expr.Accept(i); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (i *Interpreter) visitLiteralExpr(literalExpr *LiteralExpr) (interface{}, error) {
	return literalExpr.Value, nil
}

func (i *Interpreter) visitDefvarExpr(defvarExpr *DefvarExpr) (interface{}, error) {
	val, err := defvarExpr.Initializer.Accept(i)
	if err != nil {
		return nil, err
	}

	err = i.current.Define(defvarExpr.Name, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (i *Interpreter) visitVarExpr(varExpr *VarExpr) (interface{}, error) {
	return i.current.Get(varExpr.Name)
}

func (i *Interpreter) visitIfExpr(ifExpr *IfExpr) (interface{}, error) {
	cond, err := ifExpr.Condition.Accept(i)
	if err != nil {
		return nil, err
	}

	if isTruthy(cond) {
		return ifExpr.ThenBranch.Accept(i)
	}

	return ifExpr.ElseBranch.Accept(i)
}

func isTruthy(val interface{}) bool {
	if val == false || val == nil {
		return false
	}

	return true
}
