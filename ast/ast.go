package ast

type ExprType int

const (
	Expr_String ExprType = iota
	Expr_FuncCall
)

type Expr struct {
	Type ExprType
	Id   string
	Args []*Expr
}
