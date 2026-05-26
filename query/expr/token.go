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
	"strings"
)

// Token is a lexical token of the InfluxQL language.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	literal_beg
	// Literals
	IDENT     // main
	NUMBER    // 12345.67
	STRING    // "abc"
	BADSTRING // "abc
	BADESCAPE // \q
	NULL      // NULL
	UNKNOWN   // UNKNOWN
	TRUE      // true
	FALSE     // false
	literal_end

	operator_beg
	// Operators
	unary_operator_beg
	EXCLAMATION // !
	UNARY_MINUS // -
	NOT         // NOT
	BITWISE_NOT // ~0

	// Date operators
	GET_WEEK_START
	GET_MONTH_START
	GET_QUARTER_START
	GET_YEAR_START
	GET_DAY_OF_MONTH
	GET_DAY_OF_YEAR
	GET_MONTH_OF_YEAR
	GET_QUARTER_OF_YEAR
	// hll operator
	GET_HLL_VALUE
	// array operator
	ARRAY_LENGTH
	unary_operator_end

	derived_unary_operator_beg
	IS_NULL     // IS NULL
	IS_NOT_NULL // IS NOT NULL
	IS_TRUE     // IS TRUE
	IS_FALSE    // IS FALSE
	derived_unary_operator_end

	binary_operator_beg
	ADD        // +
	SUB        // -
	MUL        // *
	DIV        // /
	MOD        // %
	FLOOR      // floor
	CONVERT_TZ // convert_tz

	BITWISE_AND         // &
	BITWISE_OR          // |
	BITWISE_XOR         // ^
	BITWISE_LEFT_SHIFT  // <<
	BITWISE_RIGHT_SHIFT // >>

	AND // AND
	OR  // OR

	IN     // IN
	NOT_IN // NOT_IN
	IS     // IS
	NEQ    // !=
	// Do not modify the order of the following 5 operators
	EQ  // =
	LT  // <
	LTE // <=
	GT  // >
	GTE // >=

	// Geo intersects
	GEOGRAPHY_INTERSECTS
	// Array functions
	ARRAY_CONTAINS
	ARRAY_ELEMENT_AT
	binary_operator_end
	operator_end

	LPAREN // (
	RPAREN // )
	COMMA  // ,
	DOT    // .

	keyword_beg
	// Keywords
	ALL
	AS
	ASC
	BEGIN
	BY
	CASE
	DEFAULT
	DELETE
	DESC
	DISTINCT
	DROP
	ELSE
	END
	EXISTS
	FIELD
	FOR
	FROM
	GROUP
	IF
	INF
	INNER
	INSERT
	KEY
	KEYS
	LIMIT
	OFFSET
	ON
	ORDER
	SELECT
	THEN
	TO
	VALUES
	WHEN
	WHERE
	WITH
	keyword_end
)

var tokens = map[Token]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	IDENT:     "IDENT",
	NUMBER:    "NUMBER",
	STRING:    "STRING",
	BADSTRING: "BADSTRING",
	BADESCAPE: "BADESCAPE",
	NULL:      "NULL",
	UNKNOWN:   "UNKNOWN",
	TRUE:      "TRUE",
	FALSE:     "FALSE",

	EXCLAMATION: "!",
	UNARY_MINUS: "-",
	NOT:         "NOT",
	IS_NULL:     "IS NULL",
	IS_NOT_NULL: "IS NOT NULL",
	IS_TRUE:     "IS TRUE",
	IS_FALSE:    "IS FALSE",

	GET_WEEK_START:      "GET_WEEK_START",
	GET_MONTH_START:     "GET_MONTH_START",
	GET_QUARTER_START:   "GET_QUARTER_START",
	GET_YEAR_START:      "GET_YEAR_START",
	GET_DAY_OF_MONTH:    "GET_DAY_OF_MONTH",
	GET_DAY_OF_YEAR:     "GET_DAY_OF_YEAR",
	GET_MONTH_OF_YEAR:   "GET_MONTH_OF_YEAR",
	GET_QUARTER_OF_YEAR: "GET_QUARTER_OF_YEAR",
	GET_HLL_VALUE:       "GET_HLL_VALUE",

	ADD:        "+",
	SUB:        "-",
	MUL:        "*",
	DIV:        "/",
	MOD:        "%",
	FLOOR:      "FLOOR",
	CONVERT_TZ: "CONVERT_TZ",

	BITWISE_AND:         "&",
	BITWISE_OR:          "|",
	BITWISE_NOT:         "~",
	BITWISE_XOR:         "^",
	BITWISE_LEFT_SHIFT:  "<<",
	BITWISE_RIGHT_SHIFT: ">>",

	AND: "AND",
	OR:  "OR",

	IN:     "IN",
	NOT_IN: "NOT IN",
	IS:     "IS",
	EQ:     "=",
	NEQ:    "!=",
	LT:     "<",
	LTE:    "<=",
	GT:     ">",
	GTE:    ">=",

	GEOGRAPHY_INTERSECTS: "GEOGRAPHY_INTERSECTS",

	ARRAY_LENGTH:     "ARRAY_LENGTH",
	ARRAY_CONTAINS:   "ARRAY_CONTAINS",
	ARRAY_ELEMENT_AT: "ARRAY_ELEMENT_AT",

	LPAREN: "(",
	RPAREN: ")",
	COMMA:  ",",
	DOT:    ".",

	ALL:      "ALL",
	AS:       "AS",
	ASC:      "ASC",
	BEGIN:    "BEGIN",
	BY:       "BY",
	CASE:     "CASE",
	DEFAULT:  "DEFAULT",
	DELETE:   "DELETE",
	DESC:     "DESC",
	DROP:     "DROP",
	DISTINCT: "DISTINCT",
	ELSE:     "ELSE",
	END:      "END",
	EXISTS:   "EXISTS",
	FIELD:    "FIELD",
	FOR:      "FOR",
	FROM:     "FROM",
	GROUP:    "GROUP",
	IF:       "IF",
	INF:      "INF",
	INNER:    "INNER",
	INSERT:   "INSERT",
	KEY:      "KEY",
	KEYS:     "KEYS",
	LIMIT:    "LIMIT",
	OFFSET:   "OFFSET",
	ON:       "ON",
	ORDER:    "ORDER",
	SELECT:   "SELECT",
	THEN:     "THEN",
	TO:       "TO",
	VALUES:   "VALUES",
	WHEN:     "WHEN",
	WHERE:    "WHERE",
	WITH:     "WITH",
}

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token)
	for tok := keyword_beg + 1; tok < keyword_end; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	for _, tok := range []Token{AND, OR, IN, IS, NOT} {
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	keywords["null"] = NULL
	keywords["unknown"] = UNKNOWN
	keywords["true"] = TRUE
	keywords["false"] = FALSE
}

// String returns the string representation of the token.
func (tok Token) String() string { _ = "STUB: not implemented"; return "" }

func (tok Token) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Precedence returns the operator precedence of the binary operator token.
func (tok Token) Precedence() int { _ = "STUB: not implemented"; return 0 }

func (tok Token) isUnaryOperator() bool { _ = "STUB: not implemented"; return false }

func (tok Token) isDerivedUnaryOperator() bool { _ = "STUB: not implemented"; return false }

func (tok Token) isBinaryOperator() bool { _ = "STUB: not implemented"; return false }

// tokstr returns a literal if provided, otherwise returns the token string.
func tokstr(tok Token, lit string) string { _ = "STUB: not implemented"; return "" }

// Lookup returns the token associated with a given string.
func Lookup(ident string) Token { _ = "STUB: not implemented"; return *new(Token) }

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}
