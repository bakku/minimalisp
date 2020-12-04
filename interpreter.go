package tinylisp

import (
	"fmt"
)

// Interpreter is an implementation of the AST visitor.
type Interpreter struct {
	globals *Environment
	current *Environment
}

// NewInterpreter is a factory function to create a new Interpreter.
func NewInterpreter() *Interpreter {
	global := NewEnvironment()
	setupStdlib(global)
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

func (i *Interpreter) execute(expression Expression, env *Environment) (interface{}, error) {
	prevEnv := i.current
	i.current = env
	ret, err := expression.Accept(i)
	i.current = prevEnv
	return ret, err
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

func (i *Interpreter) visitDefunExpr(defunExpr *DefunExpr) (interface{}, error) {
	fun := NewTinyLispFunction(defunExpr.Params, defunExpr.Body, i.current)

	if err := i.current.Define(defunExpr.Name, fun); err != nil {
		return nil, err
	}

	return fun, nil
}

func (i *Interpreter) visitFuncCallExpr(funcCallExpr *FuncCallExpr) (interface{}, error) {
	fun, err := i.current.Get(funcCallExpr.Name)
	if err != nil {
		return nil, err
	}

	callableFun, ok := fun.(Function)
	if !ok {
		return nil, &executionError{funcCallExpr.Name.Line, fmt.Sprintf("%s is not a function", funcCallExpr.Name.Lexeme)}
	}

	var arguments []interface{}

	for _, arg := range funcCallExpr.Arguments {
		val, err := arg.Accept(i)
		if err != nil {
			return nil, err
		}

		arguments = append(arguments, val)
	}

	if len(arguments) != callableFun.Arity() && callableFun.Arity() != infiniteArity {
		return nil, &executionError{funcCallExpr.Name.Line, fmt.Sprintf("Expected %d arguments but got %d", callableFun.Arity(), len(arguments))}
	}

	return callableFun.Call(funcCallExpr.Name.Line, i, arguments)
}

func (i *Interpreter) visitListExpr(listExpr *ListExpr) (interface{}, error) {
	var elements []interface{}

	for _, element := range listExpr.Elements {
		val, err := element.Accept(i)
		if err != nil {
			return nil, err
		}

		elements = append(elements, val)
	}

	return NewArrayList(elements), nil
}

func (i *Interpreter) visitLetExpr(letExpr *LetExpr) (interface{}, error) {
	letEnv := NewEnvironmentWithEnclosing(i.current)

	for n, name := range letExpr.Names {
		val, err := letExpr.Values[n].Accept(i)
		if err != nil {
			return nil, err
		}

		if err = letEnv.Define(name, val); err != nil {
			return nil, err
		}
	}

	ret, err := i.execute(letExpr.Body, letEnv)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *Interpreter) visitLambdaExpr(lambdaExpr *LambdaExpr) (interface{}, error) {
	return NewTinyLispFunction(lambdaExpr.Params, lambdaExpr.Body, i.current), nil
}

func isTruthy(val interface{}) bool {
	if val == false || val == nil {
		return false
	}

	return true
}
