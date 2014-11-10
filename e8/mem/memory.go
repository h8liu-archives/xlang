// Package mem defines the memory structure for E8 virtual machine.
package mem

// Memory defines a paged memory structure
type Memory struct {
	// Weather automatically allocate a new page on page miss
	NoAutoAlloc bool

	pages  map[uint32]Page
	_align *Align
}

var noopPage = NewNoopPage()

// New creates an empty memory space.
func New() *Memory {
	ret := new(Memory)
	ret.pages = make(map[uint32]Page)
	ret._align = new(Align)
	return ret
}

// Get fetches the page associated with the address.  If the page is missing,
// and auto allocation is on, a new page will be auto allocated
func (m *Memory) Get(addr uint32) Page {
	id := PageID(addr)
	ret := m.pages[id]
	if ret == nil {
		if m.NoAutoAlloc {
			return noopPage
		}
		p := NewPage()
		m.pages[id] = p
		return p
	}

	return ret
}

// Check checks if a page exists for the address
func (m *Memory) Check(addr uint32) bool {
	return m.pages[PageID(addr)] != nil
}

func (m *Memory) align(addr uint32) *Align {
	m._align.Page = m.Get(addr)
	return m._align
}

// WriteU8 writes a byte at addr.
// If the page is missing, and auto allocation is off, this is an noop.
func (m *Memory) WriteU8(addr uint32, value uint8) {
	m.align(addr).WriteU8(addr, value)
}

// WriteU16 writes a half word at addr, the address will be automatically
// aligned down.  Byte order is little endian, and the lower bytes will be
// written first.  If the page is missing and auto allocation is off, this is a
// noop.
func (m *Memory) WriteU16(addr uint32, value uint16) {
	m.align(addr).WriteU16(addr, value)
}

// WriteU32 writes a word at addr, the address will be automatically aligned
// down.  Byte order is little endian, and the lower bytes will be written
// first.  If the page is missing and auto allocation is off, this is a noop.
func (m *Memory) WriteU32(addr uint32, value uint32) {
	m.align(addr).WriteU32(addr, value)
}

// WriteF64 writes a double precision floating point at addr, the address will
// be automatically aligned down.
func (m *Memory) WriteF64(addr uint32, value float64) {
	m.align(addr).WriteF64(addr, value)
}

// ReadU8 reads a byte at addr.
// If the page is missing and auto allocation is off, 0 is returned.
func (m *Memory) ReadU8(addr uint32) uint8 {
	return m.align(addr).ReadU8(addr)
}

// ReadU16 reads a half word at addr.
// Byte order is little endian, and the lower bytes will be read first.
// If the page is missing and auto allocation is off, 0 is returned.
func (m *Memory) ReadU16(addr uint32) uint16 {
	return m.align(addr).ReadU16(addr)
}

// ReadU32 reads a word at addr.
// Byte order is little endian, and the lower bytes will be read first.
// If the page is missing and auto allocation is off, 0 is returned.
func (m *Memory) ReadU32(addr uint32) uint32 {
	return m.align(addr).ReadU32(addr)
}

// ReadF64 reads a double precision floating point at addr.
func (m *Memory) ReadF64(addr uint32) float64 {
	return m.align(addr).ReadF64(addr)
}

// Map maps a page for the address, the address will be auto aligned down to
// page boundaries.
func (m *Memory) Map(addr uint32, page Page) {
	m.pages[PageID(addr)] = page
}

// Unmap unmaps a page for the address. The unmapped page is returned.
func (m *Memory) Unmap(addr uint32) Page {
	id := PageID(addr)
	ret := m.pages[id]
	delete(m.pages, id)
	return ret
}
