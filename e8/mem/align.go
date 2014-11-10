package mem

import (
	"e8vm.net/e8/align"
)

// Align is a page wrapper that takes a page and perform aligned reads and
// writes in a page.  Byte order is little endian.  Read and writes are
// performed byte by byte,
// where lower bytes are written/read first.
//
// If the offset is not properly aligned, it will be aligned down
// automatically.
type Align struct {
	Page
}

func maskOffset(offset uint32) uint32 { return offset & PageMask }

func offset8(offset uint32) uint32 {
	return maskOffset(offset)
}

func offset16(offset uint32) uint32 {
	return align.A16(maskOffset(offset))
}

func offset32(offset uint32) uint32 {
	return align.A32(maskOffset(offset))
}

func offset64(offset uint32) uint32 {
	return align.A64(maskOffset(offset))
}

// WriteU8 writes an uint8.
func (a *Align) WriteU8(offset uint32, value uint8) {
	a.writeU8(offset8(offset), value)
}

// WriteU16 writes an uint16.
func (a *Align) WriteU16(offset uint32, value uint16) {
	a.writeU16(offset16(offset), value)
}

// WriteU32 writes an uint32.
func (a *Align) WriteU32(offset uint32, value uint32) {
	a.writeU32(offset32(offset), value)
}

// WriteF64 writes a float64.
func (a *Align) WriteF64(offset uint32, value float64) {
	panic("todo")
}

// ReadU8 reads an uint8.
func (a *Align) ReadU8(offset uint32) uint8 {
	return a.readU8(offset8(offset))
}

// ReadU16 reads an uint16.
func (a *Align) ReadU16(offset uint32) uint16 {
	return a.readU16(offset16(offset))
}

// ReadU32 reads an uint32.
func (a *Align) ReadU32(offset uint32) uint32 {
	return a.readU32(offset32(offset))
}

// ReadF64 reads an float64.
func (a *Align) ReadF64(offset uint32) float64 {
	panic("todo")
}

func (a *Align) writeU8(offset uint32, value uint8) {
	a.Page.Write(offset, value)
}

func (a *Align) writeU16(offset uint32, value uint16) {
	a.Page.Write(offset, uint8(value))
	a.Page.Write(offset+1, uint8(value>>8))
}

func (a *Align) writeU32(offset uint32, value uint32) {
	a.Page.Write(offset, uint8(value))
	a.Page.Write(offset+1, uint8(value>>8))
	a.Page.Write(offset+2, uint8(value>>16))
	a.Page.Write(offset+3, uint8(value>>24))
}

func (a *Align) readU8(offset uint32) uint8 {
	return a.Page.Read(offset)
}

func (a *Align) readU16(offset uint32) uint16 {
	ret := uint16(a.Page.Read(offset))
	ret |= uint16(a.Page.Read(offset+1)) << 8
	return ret
}

func (a *Align) readU32(offset uint32) uint32 {
	ret := uint32(a.Page.Read(offset))
	ret |= uint32(a.Page.Read(offset+1)) << 8
	ret |= uint32(a.Page.Read(offset+2)) << 16
	ret |= uint32(a.Page.Read(offset+3)) << 24
	return ret
}
