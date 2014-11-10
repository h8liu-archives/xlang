package mem

// DataPage is a regular memoery page that saves data.
type DataPage struct {
	bytes []byte
}

var _ Page = new(DataPage)

// NewPage creates a regular memory page that saves data.
func NewPage() *DataPage {
	ret := new(DataPage)
	ret.bytes = make([]byte, PageSize)
	return ret
}

// Read reads a byte at a particuar offset.
func (p *DataPage) Read(offset uint32) uint8 {
	return p.bytes[offset]
}

// Write writes a byte at a particular offset.
func (p *DataPage) Write(offset uint32, b uint8) {
	p.bytes[offset] = b
}

// Bytes returns the content of the page.
func (p *DataPage) Bytes() []byte {
	return p.bytes
}
