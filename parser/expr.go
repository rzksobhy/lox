package parser

import "github.com/rzksobhy/lox/lexer"

type Expr[R any] interface {
	accept(visitor Visitor[R]) R
}

type Visitor[R any] interface {
	VisitAssignExpr(expr *Assign[R]) R
	VisitBinaryExpr(expr *Binary[R]) R
	VisitCallExpr(expr *Call[R]) R
	VisitGetExpr(expr *Get[R]) R
	VisitGroupingExpr(expr *Grouping[R]) R
	VisitLiteralExpr(expr *Literal[R]) R
	VisitLogicalExpr(expr *Logical[R]) R
	VisitSetExpr(expr *Set[R]) R
	VisitSuperExpr(expr *Super[R]) R
	VisitThisExpr(expr *This[R]) R
	VisitUnaryExpr(expr *Unary[R]) R
	VisitVariableExpr(expr *Variable[R]) R
}

type Assign[R any] struct {
	token lexer.Token
	value Expr[any]
}

func (self *Assign[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitAssignExpr(self)
}

type Binary[R any] struct {
	left     Expr[any]
	operator lexer.Token
	right    Expr[any]
}

func (self *Binary[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitBinaryExpr(self)
}

type Call[R any] struct {
	callee Expr[any]
	paran  lexer.Token
	args   []Expr[any]
}

func (self *Call[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitCallExpr(self)
}

type Get[R any] struct {
	object Expr[any]
	name   lexer.Token
}

func (self *Get[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitGetExpr(self)
}

type Grouping[R any] struct {
	expression Expr[any]
}

func (self *Grouping[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitGroupingExpr(self)
}

type Literal[R any] struct {
	value any
}

func (self *Literal[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitLiteralExpr(self)
}

type Logical[R any] struct {
	left     Expr[any]
	operator lexer.Token
	right    Expr[any]
}

func (self *Logical[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitLogicalExpr(self)
}

type Set[R any] struct {
	object Expr[any]
	name   lexer.Token
	value  Expr[any]
}

func (self *Set[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitSetExpr(self)
}

type Super[R any] struct {
	keyword lexer.Token
	method  lexer.Token
}

func (self *Super[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitSuperExpr(self)
}

type This[R any] struct {
	keyword lexer.Token
}

func (self *This[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitThisExpr(self)
}

type Unary[R any] struct {
	operator lexer.Token
	right    Expr[any]
}

func (self *Unary[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitUnaryExpr(self)
}

type Variable[R any] struct {
	name lexer.Token
}

func (self *Variable[R]) accept(visitor Visitor[R]) R {
	return visitor.VisitVariableExpr(self)
}
