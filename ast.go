package tinylisp

// Expression is the interface all types of expressions must fulfil.
type Expression interface {
	Accept(visitor visitor) (interface{}, error)
}

// Visitor is the interface which an interpreter has to fulfil.
type visitor interface {
	visitLiteralExpr(literalExpr *LiteralExpr) (interface{}, error)
	visitDefvarExpr(defvarExpr *DefvarExpr) (interface{}, error)
	visitVarExpr(varExpr *VarExpr) (interface{}, error)
	visitIfExpr(ifExpr *IfExpr) (interface{}, error)
	visitDefunExpr(defunExpr *DefunExpr) (interface{}, error)
	visitFuncCallExpr(funcCallExpr *FuncCallExpr) (interface{}, error)
	visitListExpr(listExpr *ListExpr) (interface{}, error)
	visitLetExpr(letExpr *LetExpr) (interface{}, error)
	visitLambdaExpr(lambdaExpr *LambdaExpr) (interface{}, error)
}

// LiteralExpr is a literal such as a string or a number.
type LiteralExpr struct {
	Value interface{}
}

// Accept visits the literal expression.
func (e *LiteralExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitLiteralExpr(e)
}

// DefvarExpr is a definition of a variable.
type DefvarExpr struct {
	Name        Token
	Initializer Expression
}

// Accept visits the defvar expression.
func (e *DefvarExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitDefvarExpr(e)
}

// VarExpr is a reference to a variable.
type VarExpr struct {
	Name Token
}

// Accept visits the var expression.
func (e *VarExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitVarExpr(e)
}

// IfExpr is an if expression :).
type IfExpr struct {
	Condition  Expression
	ThenBranch Expression
	ElseBranch Expression
}

// Accept visits the if expression.
func (e *IfExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitIfExpr(e)
}

// DefunExpr is a definition of a function.
type DefunExpr struct {
	Name   Token
	Params []Token
	Body   Expression
}

// Accept visits the function definition.
func (e *DefunExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitDefunExpr(e)
}

// FuncCallExpr is a function call.
type FuncCallExpr struct {
	Name      Token
	Arguments []Expression
}

// Accept visits the function call.
func (e *FuncCallExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitFuncCallExpr(e)
}

// ListExpr represents a list collection.
type ListExpr struct {
	Elements []Expression
}

// Accept visits the list expression.
func (e *ListExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitListExpr(e)
}

// LetExpr is a let expression to define local variables.
type LetExpr struct {
	Names  []Token
	Values []Expression
	Body   Expression
}

// Accept visits the let expression.
func (e *LetExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitLetExpr(e)
}

// LambdaExpr is a lambda expression to define anonymous functions.
type LambdaExpr struct {
	Params []Token
	Body   Expression
}

// Accept visits the lambda expression.
func (e *LambdaExpr) Accept(visitor visitor) (interface{}, error) {
	return visitor.visitLambdaExpr(e)
}
