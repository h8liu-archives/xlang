package mem

// Page defines a general page inteface.
type Page interface {
	// Writes a byte at a particular page offset.
	Write(offset uint32, b uint8)

	// Reads a byte at a particular page offset.
	Read(offset uint32) uint8
}

const (
	// PageOffset is the number of bits for page offset.
	PageOffset = 12

	// PageSize is the number of bytes a page has.
	PageSize = 1 << PageOffset

	// PageMask is the mask for masking the page offset.
	PageMask = PageSize - 1
)

// PageStart returns the starting address of a page with a particular page id.
func PageStart(i uint32) uint32 { return i << PageOffset }

// PageID returns the page id of a memory address.
func PageID(i uint32) uint32 { return i >> PageOffset }
