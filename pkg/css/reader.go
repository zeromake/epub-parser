package css

import (
	"bufio"
	"fmt"
	"io"
)

// Token
type Token interface {}

// A ParseError is returned for parsing errors.
// Line numbers are 1-indexed and columns are 0-indexed.
type ParseError struct {
	Line      int   // Line where the error occurred
	Column    int   // Column (rune index) where the error occurred
	Err       error // The actual error
}

func (e *ParseError) Error() string {
	// github/pkg/errors
	return fmt.Sprintf("parse error on line %d, column %d:\n%+v\n", e.Line, e.Column, e.Err)
}

// Decoder css reader
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r:     bufio.NewReader(r),
	}
}
// Token get next token
func (d *Decoder) Token() (Token, error) {
	return nil, nil
}