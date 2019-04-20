package css

import (
	"bufio"
	"io"
)

// Encoder css encoder
type Encoder struct {
	w *bufio.Writer
}


// NewDecoder
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w: bufio.NewWriter(w),
	}
}

// Token write token
func (e *Encoder) Token(token Token) error {
	return nil
}

// Encode encode struct to writer
func (e *Encoder)Encode(v interface{}) error {
	return nil
}
