package img

import (
	"io"
)

// Header describes the size and starting address of an image section.
type Header struct {
	addr uint32
	size uint32
}

func u32(buf []byte) uint32 {
	ret := uint32(buf[0])
	ret |= uint32(buf[1]) << 8
	ret |= uint32(buf[2]) << 16
	ret |= uint32(buf[3]) << 24
	return ret
}

func pu32(buf []byte, i uint32) {
	buf[0] = uint8(i)
	buf[1] = uint8(i >> 8)
	buf[2] = uint8(i >> 16)
	buf[3] = uint8(i >> 24)
}

// ReadIn unmarshalls a header from an input stream.
func (h *Header) ReadIn(in io.Reader) error {
	buf := make([]byte, 8)
	_, e := io.ReadFull(in, buf)
	if e != nil {
		return e
	}

	h.addr = u32(buf[0:4])
	h.size = u32(buf[4:8])
	return nil
}

// WriteTo writes a header out to an output stream.
func (h *Header) WriteTo(out io.Writer) error {
	buf := make([]byte, 8)
	pu32(buf[0:4], h.addr)
	pu32(buf[4:8], h.size)

	_, e := out.Write(buf)
	return e
}

// Start returns the starting address of the section.
func (h *Header) Start() uint32 { return h.addr }
