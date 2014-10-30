package parser

// EntryScanner is helps on parsing a statement
type EntryScanner struct {
	s Stmt

	last *Tok
}

// NewEntryScanner creates a new EntryScanner
func NewEntryScanner(s Stmt) *EntryScanner {
	ret := new(EntryScanner)
	ret.s = s
	return ret
}

// See return true if the current token is of one the types.
func (s *EntryScanner) See(types ...Type) bool {
	panic("todo")
}

// SeeOp returns true if the current token is one of the ops.
func (s *EntryScanner) SeeOp(ops ...string) bool {
	panic("todo")
}

// AcceptOp returns true and accepts the current token if the
// current token is the exact op. It returns false and is
// an noop otherwise.
func (s *EntryScanner) AcceptOp(op string) bool {
	panic("todo")
}

// Accept returns and accepts the current token entry.
// It panics if the current entry is a block.
func (s *EntryScanner) Accept() *Tok {
	panic("todo")
}

// Tok returns the current token entry.
// It panics if the current entry is a block.
func (s *EntryScanner) Tok() *Tok {
	panic("todo")
}

// IsBlock returns true if the current entry is a block.
func (s *EntryScanner) IsBlock() bool {
	panic("todo")
}

func (s *EntryScanner) Pos() *Pos {
	t := s.Tok()
	if t != nil {
		return t.Pos
	}
	if s.last != nil {
		return s.last.Pos
	}
	return nil
}
