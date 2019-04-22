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
	e := &Encoder{}
	if wr, ok := w.(*bufio.Writer); ok {
		e.w = wr
	} else {
		e.w = bufio.NewWriter(w)
	}
	return e
}

// Token write token
func (e *Encoder) Token(token Token) error {
	return nil
}

// Encode encode struct to writer
func (e *Encoder)Encode(v interface{}) error {
	return nil
}
