package img

import (
	"fmt"
	"io"
)

// Writer wraps an output stream for creating
// virtual machine images.
type Writer struct {
	io.Writer
}

// NewWriter creates a virtual machine image writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

// Write writes a section with a particular starting address.
func (w *Writer) Write(addr uint32, bytes []byte) (e error) {
	return Write(w.Writer, addr, bytes)
}

// Write writes a section with a particualr starting address
// to an output stream.
func Write(out io.Writer, addr uint32, bytes []byte) (e error) {
	n := uint64(len(bytes))
	if n > (1 << 31) {
		// this is almost impossible to happen
		return fmt.Errorf("too many bytes")
	}
	if uint64(addr)+n > (1 << 32) {
		return fmt.Errorf("out of memory space")
	}

	header := new(Header)
	header.addr = addr
	header.size = uint32(n)
	if e = header.WriteTo(out); e != nil {
		return e
	}

	if _, e = out.Write(bytes); e != nil {
		return e
	}

	return
}
