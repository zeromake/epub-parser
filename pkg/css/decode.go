package css

import (
	"bufio"
	"fmt"
	"io"
)

// A ParseError is returned for parsing errors.
// Line numbers are 1-indexed and columns are 0-indexed.
type ParseError struct {
	Line   int   // Line where the error occurred
	Column int   // Column (rune index) where the error occurred
	Err    error // The actual error
}

func (e *ParseError) Error() string {
	// github/pkg/errors
	return fmt.Sprintf("css parse error on line %d, column %d:\n%+v\n", e.Line, e.Column, e.Err)
}

// Decoder css reader
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder new a decoder
func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{}
	if rr, ok := r.(*bufio.Reader); ok {
		d.r = rr
	} else {
		d.r = bufio.NewReader(r)
	}
	return d
}

// Token get next token
func (d *Decoder) Token() (Token, error) {
	return nil, nil
}

// Decode decode reader to struct
func (d *Decoder) Decode(v interface{}) error {
	return nil
}
