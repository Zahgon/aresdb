// Modifications Copyright (c) 2017-2018 Uber Technologies, Inc.
// Copyright (c) 2013-2016 Errplane Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package expr

import (
	memCom "github.com/uber/aresdb/memstore/common"
)

// Type defines data types for expression evaluation.
// Expression types are determined at query compilation time, type castings are
// generated when apprioperiate. Notice that word widths are not specified here.
type Type int

const (
	UnknownType Type = iota
	Boolean
	Unsigned
	Signed
	Float
	GeoPoint
	GeoShape
	UUID
)

var typeNames = map[Type]string{
	UnknownType: "Unknown",
	Boolean:     "Boolean",
	Unsigned:    "Unsigned",
	Signed:      "Signed",
	Float:       "Float",
	GeoPoint:    "GeoPoint",
	GeoShape:    "GeoShape",
	UUID:        "UUID",
}

// constants for call names.
const (
	ConvertTzCallName           = "convert_tz"
	CountCallName               = "count"
	DayOfWeekCallName           = "dayofweek"
	FromUnixTimeCallName        = "from_unixtime"
	GeographyIntersectsCallName = "geography_intersects"
	HexCallName                 = "hex"
	// hll aggregation function applies to hll columns
	HllCallName = "hll"
	// countdistincthll aggregation function applies to all columns, hll value is computed on the fly
	CountDistinctHllCallName = "countdistincthll"
	HourCallName             = "hour"
	ListCallName             = ""
	MaxCallName              = "max"
	MinCallName              = "min"
	SumCallName              = "sum"
	AvgCallName              = "avg"
	// array functions
	LengthCallName    = "length"
	ContainsCallName  = "contains"
	ElementAtCallName = "element_at"
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func (t Type) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Expr represents an expression that can be evaluated to a value.
type Expr interface {
	expr()
	String() string
	Type() Type
}

func (*BinaryExpr) expr()      { _ = "STUB: not implemented"; return }
func (*BooleanLiteral) expr()  { _ = "STUB: not implemented"; return }
func (*Call) expr()            { _ = "STUB: not implemented"; return }
func (*Case) expr()            { _ = "STUB: not implemented"; return }
func (*Distinct) expr()        { _ = "STUB: not implemented"; return }
func (*NullLiteral) expr()     { _ = "STUB: not implemented"; return }
func (*NumberLiteral) expr()   { _ = "STUB: not implemented"; return }
func (*ParenExpr) expr()       { _ = "STUB: not implemented"; return }
func (*StringLiteral) expr()   { _ = "STUB: not implemented"; return }
func (*UnaryExpr) expr()       { _ = "STUB: not implemented"; return }
func (*UnknownLiteral) expr()  { _ = "STUB: not implemented"; return }
func (*VarRef) expr()          { _ = "STUB: not implemented"; return }
func (*Wildcard) expr()        { _ = "STUB: not implemented"; return }
func (*GeopointLiteral) expr() { _ = "STUB: not implemented"; return }
func (*UUIDLiteral) expr() {
	_ = "STUB: not implemented"

	// walkNames will walk the Expr and return the database fields
	return
}

func walkNames(exp Expr) []string { _ = "STUB: not implemented"; return nil }

// walkFunctionCalls walks the Field of a query for any function calls made
func walkFunctionCalls(exp Expr) []*Call { _ = "STUB: not implemented"; return nil }

// VarRef represents a reference to a variable.
type VarRef struct {
	Val      string
	ExprType Type

	// The following fields are populated for convenience after the query is
	// validated against the schema.

	// ID of the table in the query scope (0 for the main table, 1+ for foreign
	// tables).
	TableID int
	// ID of the column in the schema.
	ColumnID int
	// Enum dictionary for enum typed column. Can only be accessed while holding
	// the schema lock.
	EnumDict map[string]int `json:"-"`
	// Setting enum reverse dict requires holding the schema lock,
	// while reading from it does not require holding the schema lock.
	EnumReverseDict []string `json:"-"`

	DataType memCom.DataType

	// Whether this column is hll column (can run hll directly)
	IsHLLColumn bool
}

// Type returns the type.
func (r *VarRef) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the variable reference.
	return *new(Type)
}

func (r *VarRef) String() string {
	_ = "STUB: not implemented"

	// Call represents a function call.
	return ""
}

type Call struct {
	Name     string
	Args     []Expr
	ExprType Type
}

// Type returns the type.
func (c *Call) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the call.
	return *new(Type)
}

func (c *Call) String() string {
	_ = "STUB: not implemented"
	// Join arguments.
	return ""
}

// Write function name and args.

// WhenThen represents a when-then conditional expression pair in a case expression.
type WhenThen struct {
	When Expr
	Then Expr
}

// Case represents a CASE WHEN .. THEN .. ELSE .. THEN expression.
type Case struct {
	WhenThens []WhenThen
	Else      Expr
	ExprType  Type
}

// Type returns the type.
func (c *Case) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the expression.
	return *new(Type)
}

func (c *Case) String() string { _ = "STUB: not implemented"; return "" }

// Distinct represents a DISTINCT expression.
type Distinct struct {
	// Identifier following DISTINCT
	Val string
}

// Type returns the type.
func (d *Distinct) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the expression.
	return *new(Type)
}

func (d *Distinct) String() string { _ = "STUB: not implemented"; return "" }

// NewCall returns a new call expression from this expressions.
func (d *Distinct) NewCall() *Call { _ = "STUB: not implemented"; return nil }

// NumberLiteral represents a numeric literal.
type NumberLiteral struct {
	Val      float64
	Int      int
	Expr     string
	ExprType Type
}

// Type returns the type.
func (l *NumberLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the literal.
	return *new(Type)
}

func (l *NumberLiteral) String() string { _ = "STUB: not implemented"; return "" }

// BooleanLiteral represents a boolean literal.
type BooleanLiteral struct {
	Val bool
}

// Type returns the type.
func (l *BooleanLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the literal.
	return *new(Type)
}

func (l *BooleanLiteral) String() string { _ = "STUB: not implemented"; return "" }

// isTrueLiteral returns true if the expression is a literal "true" value.
func isTrueLiteral(expr Expr) bool { _ = "STUB: not implemented"; return false }

// isFalseLiteral returns true if the expression is a literal "false" value.
func isFalseLiteral(expr Expr) bool { _ = "STUB: not implemented"; return false }

// StringLiteral represents a string literal.
type StringLiteral struct {
	Val string
}

// Type returns the type.
func (l *StringLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the literal.
	return *new(Type)
}

func (l *StringLiteral) String() string { _ = "STUB: not implemented"; return "" }

// GeopointLiteral represents a literal for GeoPoint
type GeopointLiteral struct {
	Val [2]float32
}

// Type returns the type.
func (l *GeopointLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the literal.
	return *new(Type)
}

func (l *GeopointLiteral) String() string { _ = "STUB: not implemented"; return "" }

// UUIDLiteral represents a literal for UUID
type UUIDLiteral struct {
	Val [2]uint64
}

// Type returns the type.
func (l *UUIDLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the literal.
	return *new(Type)
}

func (l *UUIDLiteral) String() string { _ = "STUB: not implemented"; return "" }

// NullLiteral represents a NULL literal.
type NullLiteral struct{}

// Type returns the type.
func (l *NullLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns "NULL".
	return *new(Type)
}

func (l *NullLiteral) String() string {
	_ = "STUB: not implemented"

	// UnknownLiteral represents an UNKNOWN literal.
	return ""
}

type UnknownLiteral struct{}

// Type returns the type.
func (l *UnknownLiteral) Type() Type {
	_ = "STUB: not implemented"

	// String returns "UNKNOWN".
	return *new(Type)
}

func (l *UnknownLiteral) String() string {
	_ = "STUB: not implemented"

	// UnaryExpr represents an operation on a single expression.
	return ""
}

type UnaryExpr struct {
	Op       Token
	Expr     Expr
	ExprType Type
}

// Type returns the type.
func (e *UnaryExpr) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the unary expression.
	return *new(Type)
}

func (e *UnaryExpr) String() string { _ = "STUB: not implemented"; return "" }

// BinaryExpr represents an operation between two expressions.
type BinaryExpr struct {
	Op       Token
	LHS      Expr
	RHS      Expr
	ExprType Type
}

// Type returns the type.
func (e *BinaryExpr) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the binary expression.
	return *new(Type)
}

func (e *BinaryExpr) String() string { _ = "STUB: not implemented"; return "" }

// ParenExpr represents a parenthesized expression.
type ParenExpr struct {
	Expr     Expr
	ExprType Type // used for type casting
}

// Type returns the type.
func (e *ParenExpr) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

// String returns a string representation of the parenthesized expression.
func (e *ParenExpr) String() string { _ = "STUB: not implemented"; return "" }

// Wildcard represents a wild card expression.
type Wildcard struct{}

// Type returns the type.
func (e *Wildcard) Type() Type {
	_ = "STUB: not implemented"

	// String returns a string representation of the wildcard.
	return *new(Type)
}

func (e *Wildcard) String() string {
	_ = "STUB: not implemented"

	// CloneExpr returns a deep copy of the expression.
	return ""
}

func CloneExpr(expr Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Visitor can be called by Walk to traverse an AST hierarchy.
// The Visit() function is called once per expression.
type Visitor interface {
	Visit(Expr) Visitor
}

// Walk traverses an expression hierarchy in depth-first order.
func Walk(v Visitor, expr Expr) { _ = "STUB: not implemented"; return }

// WalkFunc traverses an expression hierarchy in depth-first order.
func WalkFunc(e Expr, fn func(Expr)) { _ = "STUB: not implemented"; return }

type walkFuncVisitor func(Expr)

func (fn walkFuncVisitor) Visit(e Expr) Visitor {
	_ = "STUB: not implemented"

	// Rewriter can be called by Rewrite to replace nodes in the AST hierarchy.
	// The Rewrite() function is called once per expression.
	return *new(Visitor)
}

type Rewriter interface {
	Rewrite(Expr) Expr
}

// Rewrite recursively invokes the rewriter to replace each expression.
// Nodes are traversed depth-first and rewritten from leaf to root.
func Rewrite(r Rewriter, expr Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// RewriteFunc rewrites an expression hierarchy.
func RewriteFunc(e Expr, fn func(Expr) Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }

type rewriterFunc func(Expr) Expr

func (fn rewriterFunc) Rewrite(e Expr) Expr {
	_ = "STUB: not implemented"

	// IsUUIDColumn returns whether an Expr is UUID
	return *new(Expr)
}

func IsUUIDColumn(expression Expr) bool { _ = "STUB: not implemented"; return false }

// Cast returns an expression that casts the input to the desired type.
// The returned expression AST will be used directly for VM instruction
// generation of the desired types.
func Cast(e Expr, t Type) Expr {
	_ = "STUB: not implemented"
	// Input type is already desired.
	return *new(Expr)
}

// Type casting is only required if at least one side is float.
// We do not cast (or check for overflow) among boolean, signed and unsigned.

// Data type for NumberLiteral can be changed directly.

// Use ParenExpr to respresent a VM type cast.
