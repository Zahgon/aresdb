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
	"io"
)

// Parser represents an InfluxQL parser.
type Parser struct {
	s *bufScanner
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser { _ = "STUB: not implemented"; return nil }

// ParseExpr parses an expression string and returns its AST representation.
func ParseExpr(s string) (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

// parseInt parses a string and returns an integer literal.
func (p *Parser) parseInt(min, max int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Return an error if the number has a fractional part.

// Convert string to int.

// parseUInt32 parses a string and returns a 32-bit unsigned integer literal.
func (p *Parser) parseUInt32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// Convert string to unsigned 32-bit integer

// parseUInt64 parses a string and returns a 64-bit unsigned integer literal.
func (p *Parser) parseUInt64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Convert string to unsigned 64-bit integer

// parseIdent parses an identifier.
func (p *Parser) parseIdent() (string, error) { _ = "STUB: not implemented"; return "", nil }

// parseIdentList parses a comma delimited list of identifiers.
func (p *Parser) parseIdentList() ([]string, error) {
	_ = "STUB: not implemented"
	// Parse first (required) identifier.
	return nil, nil
}

// Parse remaining (optional) identifiers.

// parseSegmentedIdents parses a segmented identifiers.
// e.g.,  "db"."rp".measurement  or  "db"..measurement
func (p *Parser) parseSegmentedIdents() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse remaining (optional) identifiers.

// No more segments so we're done.

// Next segment is a regex so we're done.

// Add an empty identifier.

// Parse the next identifier.

// parserString parses a string.
func (p *Parser) parseString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// peekRune returns the next rune that would be read by the scanner.
func (p *Parser) peekRune() rune { _ = "STUB: not implemented"; return 0 }

// parseOptionalTokenAndInt parses the specified token followed
// by an int, if it exists.
func (p *Parser) parseOptionalTokenAndInt(t Token) (int, error) {
	_ = "STUB: not implemented"
	// Check if the token exists.
	return 0, nil
}

// Scan the number.

// Return an error if the number has a fractional part.

// Parse number.

// parseVarRef parses a reference to a measurement or field.
func (p *Parser) parseVarRef() (*VarRef, error) {
	_ = "STUB: not implemented"
	// Parse the segments of the variable ref.
	return nil, nil
}

func rewriteIsOp(expr Expr) (Token, error) { _ = "STUB: not implemented"; return *new(Token), nil }

func rewriteIsExpr(expr Expr) (Expr, error) { _ = "STUB: not implemented"; return *new(Expr), nil }

// ParseExpr parses an expression.
// binOpPrcdncLb: binary operator precedence lower bound.
// Any binary operator with a lower precedence than that will cause ParseExpr to stop.
// This is used for parsing binary operators following a unary operator.
func (p *Parser) ParseExpr(binOpPrcdncLb int) (Expr, error) {
	_ = "STUB: not implemented"

	// Dummy root node.
	return *new(Expr), nil
}

// Parse a non-binary expression type to start.
// This variable will always be the root of the expression tree.

// Loop over operations and unary exprs and build a tree based on precendence.

// If the next token is NOT an operator then return the expression.

// Otherwise parse the next expression.

// Find the right spot in the tree to add the new expression by
// descending the RHS of the expression tree until we reach the last
// BinaryExpr or a BinaryExpr whose RHS has an operator with
// precedence >= the operator being added.

// Add the new expression here and break.

// parseUnaryExpr parses an non-binary expression.
// TODO: shz@ revisit inclusion parameter when open sourcing
func (p *Parser) parseUnaryExpr(inclusion bool) (Expr, error) {
	_ = "STUB: not implemented"
	// If the first token is a LPAREN then parse it as its own grouped expression.
	return *new(Expr), nil
}

// Expect an RPAREN at the end.

// Parse a tuple as a function call with empty name.

// Parse an expression argument.

// If there's not a comma next then stop parsing arguments.

// There should be a right parentheses at the end.

// Read next token.

// If the next immediate token is a left parentheses, parse as function call.
// Otherwise parse as a variable reference.

// unscan the last token (wasn't an LPAREN)
// unscan the IDENT token

// Parse it as a VarRef.

// If the next immediate token is a left parentheses, parse as function call.
// Otherwise parse as a Distinct expression.

// Assumes CASE token has been scanned.
func (p *Parser) parseCase() (*Case, error) { _ = "STUB: not implemented"; return nil, nil }

// parseCall parses a function call.
// This function assumes the function name and LPAREN have been consumed.
func (p *Parser) parseCall(name string) (*Call, error) {
	_ = "STUB: not implemented"
	return nil,

		// If there's a right paren then just return immediately.
		nil
}

// Otherwise parse function call arguments.

// Parse an expression argument.

// If there's not a comma next then stop parsing arguments.

// There should be a right parentheses at the end.

// scan returns the next token from the underlying scanner.
func (p *Parser) scan() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"

	// scanIgnoreWhitespace scans the next non-whitespace token.
	return *new(Token), *new(Pos), ""
}

func (p *Parser) scanIgnoreWhitespace() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), *new(Pos), ""
}

// consumeWhitespace scans the next token if it's whitespace.
func (p *Parser) consumeWhitespace() { _ = "STUB: not implemented"; return }

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() {
	_ = "STUB: not implemented"

	// QuoteString returns a quoted string.
	return
}

func QuoteString(s string) string { _ = "STUB: not implemented"; return "" }

// QuoteIdent returns a quoted identifier from multiple bare identifiers.
func QuoteIdent(segments ...string) string { _ = "STUB: not implemented"; return "" }

// not last segment && not ""

// IdentNeedsQuotes returns true if the ident string given would require quotes.
func IdentNeedsQuotes(ident string) bool {
	_ = "STUB: not implemented"
	// check if this identifier is a keyword
	return false
}

// split splits a string into a slice of runes.
func split(s string) (a []rune) { _ = "STUB: not implemented"; return nil }

// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Message  string
	Found    string
	Expected []string
	Pos      Pos
}

// newParseError returns a new instance of ParseError.
func newParseError(found string, expected []string, pos Pos) *ParseError {
	_ = "STUB: not implemented"
	return nil
}

// Error returns the string representation of the error.
func (e *ParseError) Error() string { _ = "STUB: not implemented"; return "" }
