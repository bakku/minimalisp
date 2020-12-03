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
