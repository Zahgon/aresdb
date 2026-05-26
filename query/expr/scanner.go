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
	"errors"
	"io"
)

// Scanner represents a lexical scanner for InfluxQL.
type Scanner struct {
	r              *reader
	lastNonWSToken Token
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner { _ = "STUB: not implemented"; return nil }

// Scan returns the next token and position from the underlying reader.
// Also returns the literal text read for strings, numbers, and duration tokens
// since these token types can have different literal representations.
func (s *Scanner) Scan() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), *new(Pos), ""
}

func (s *Scanner) scan() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	// Read next code point.
	return *new(Token), *new(Pos), ""
}

// If we see whitespace then consume all contiguous whitespace.
// If we see a letter, or certain acceptable special characters, then consume
// as an ident or reserved word.

// Otherwise parse individual characters.

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	// Create a buffer and read the current character into it.
	return *new(Token), *new(Pos), ""
}

// Read every subsequent whitespace character into the buffer.
// Non-whitespace characters and EOF will cause the loop to exit.

func (s *Scanner) scanIdent() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	// Save the starting position of the identifier.
	return *new(Token), *new(Pos), ""
}

// If the literal matches a keyword then return that keyword.

// scanString consumes a contiguous string of non-quote characters.
// Quote characters can be consumed if they're first escaped with a backslash.
func (s *Scanner) scanString() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), *new(Pos), ""
}

// scanNumber consumes anything that looks like the start of a number.
// Numbers start with a digit, full stop, plus sign or minus sign.
// This function can return non-number tokens if a scan is a false positive.
// For example, a minus sign followed by a letter will just return a minus sign.
func (s *Scanner) scanNumber() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *

	// Check if the initial rune is a "+" or "-".
	new(Token), *new(Pos), ""
}

// Read as many digits as possible.

// scanDigits consume a contiguous series of digits.
func (s *Scanner) scanDigits() string { _ = "STUB: not implemented"; return "" }

func (s *Scanner) scanHexChars() string { _ = "STUB: not implemented"; return "" }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { _ = "STUB: not implemented"; return false }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { _ = "STUB: not implemented"; return false }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { _ = "STUB: not implemented"; return false }

// isIdentChar returns true if the rune can be used in an unquoted identifier.
func isIdentChar(ch rune) bool { _ = "STUB: not implemented"; return false }

func isHexChar(ch rune) bool { _ = "STUB: not implemented"; return false }

// isIdentFirstChar returns true if the rune can be used as the first char in an unquoted identifer.
func isIdentFirstChar(ch rune) bool { _ = "STUB: not implemented"; return false }

// bufScanner represents a wrapper for scanner to add a buffer.
// It provides a fixed-length circular buffer that can be unread.
type bufScanner struct {
	s   *Scanner
	i   int // buffer index
	n   int // buffer size
	buf [3]struct {
		tok Token
		pos Pos
		lit string
	}
}

// newBufScanner returns a new buffered scanner for a reader.
func newBufScanner(r io.Reader) *bufScanner { _ = "STUB: not implemented"; return nil }

// Scan reads the next token from the scanner.
func (s *bufScanner) Scan() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), *new(Pos), ""
}

// scanFunc uses the provided function to scan the next token.
func (s *bufScanner) scanFunc(scan func() (Token, Pos, string)) (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	// If we have unread tokens then read them off the buffer first.
	return *new(Token), *new(Pos), ""
}

// Move buffer position forward and save the token.

// Unscan pushes the previously token back onto the buffer.
func (s *bufScanner) Unscan() {
	_ = "STUB: not implemented"

	// curr returns the last read token.
	return
}

func (s *bufScanner) curr() (tok Token, pos Pos, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), *new(Pos), ""
}

// reader represents a buffered rune reader used by the scanner.
// It provides a fixed-length circular buffer that can be unread.
type reader struct {
	r   io.RuneScanner
	i   int // buffer index
	n   int // buffer char count
	pos Pos // last read rune position
	buf [3]struct {
		ch  rune
		pos Pos
	}
	eof bool // true if reader has ever seen eof.
}

// ReadRune reads the next rune from the reader.
// This is a wrapper function to implement the io.RuneReader interface.
// Note that this function does not return size.
func (r *reader) ReadRune() (ch rune, size int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// UnreadRune pushes the previously read rune back onto the buffer.
// This is a wrapper function to implement the io.RuneScanner interface.
func (r *reader) UnreadRune() error { _ = "STUB: not implemented"; return nil }

// read reads the next rune from the reader.
func (r *reader) read() (ch rune, pos Pos) {
	_ = "STUB: not implemented"
	// If we have unread characters then read them off the buffer first.
	return 0, *new(Pos)
}

// Read next rune from underlying reader.
// Any error (including io.EOF) should return as EOF.

// nop

// Save character and position to the buffer.

// Update position.
// Only count EOF once.

// Mark the reader as EOF.
// This is used so we don't double count EOF characters.

// unread pushes the previously read rune back onto the buffer.
func (r *reader) unread() {
	_ = "STUB: not implemented"

	// curr returns the last read character and position.
	return
}

func (r *reader) curr() (ch rune, pos Pos) { _ = "STUB: not implemented"; return 0, *new(Pos) }

// eof is a marker code point to signify that the reader can't read any more.
const eof = rune(0)

func ScanDelimited(r io.RuneScanner, start, end rune, escapes map[rune]rune, escapesPassThru bool) ([]byte, error) {
	_ = "STUB: not implemented"
	// Scan start delimiter.
	return nil, nil
}

// If the next character is an escape then write the escaped char.
// If it's not a valid escape then return an error.

// Unread ch1 (char after the \)

// Write ch0 (\) to the output buffer.

// ScanString reads a quoted string from a rune reader.
func ScanString(r io.RuneScanner) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If the next character is an escape then write the escaped char.
// If it's not a valid escape then return an error.

var errBadString = errors.New("bad string")
var errBadEscape = errors.New("bad escape")

// ScanBareIdent reads bare identifier from a rune reader.
func ScanBareIdent(r io.RuneScanner) string {
	_ = "STUB: not implemented"
	// Read every ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	return ""
}
