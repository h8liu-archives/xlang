package mem

// NoopPage is a page where reading and writing on it
// are always noops.
type NoopPage struct{}

// NewNoopPage creates a page where reading and writing on it
// are always noops.
func NewNoopPage() *NoopPage {
	return new(NoopPage)
}

var _ Page = new(NoopPage)

// Write writes a byte at a particular offset, but has no effect.
func (p *NoopPage) Write(offset uint32, b uint8) {}

// Read reads a byte at a particular offset, but has no effect.
func (p *NoopPage) Read(offset uint32) uint8 { return 0 }
