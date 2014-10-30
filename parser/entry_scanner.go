package parser

// EntryScanner is helps on parsing a statement
type EntryScanner struct {
	s     Stmt
	index int

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
	entry := s.Entry()
	if entry == nil {
		return false
	}

	if entry.Block != nil {
		return false
	}

	for _, t := range types {
		if entry.Tok.Type == t {
			return true
		}
	}

	return false
}

// SeeOp returns true if the current token is one of the ops.
func (s *EntryScanner) SeeOp(ops ...string) bool {
	if !s.See(TypeOperator) {
		return false
	}

	t := s.Tok()
	for _, op := range ops {
		if t.Lit == op {
			return true
		}
	}

	return false
}

// AcceptOp returns true and accepts the current token if the
// current token is the exact op. It returns false and is
// an noop otherwise.
func (s *EntryScanner) AcceptOp(op string) bool {
	if !s.SeeOp(op) {
		return false
	}

	s.Accept()
	return true
}

// Accept returns and accepts the current token entry.
// It panics if the current entry is a block.
func (s *EntryScanner) Accept() *Tok {
	ret := s.Tok()
	s.index++
	return ret
}

// Tok returns the current token entry.
// It panics if the current entry is a block.
func (s *EntryScanner) Tok() *Tok {
	entry := s.Entry()
	if entry == nil {
		return nil
	}
	if entry.Block != nil {
		panic("current entry is a block")
	}

	return entry.Tok
}

// IsBlock returns true if the current entry is a block.
func (s *EntryScanner) SeeBlock() bool {
	entry := s.Entry()
	if entry == nil {
		return false
	}
	return entry.Block != nil
}

// Entry returns the current entry
func (s *EntryScanner) Entry() *Entry {
	if s.index < len(s.s) {
		return s.s[s.index]
	}

	return nil
}

// Pos returns the pos
func (s *EntryScanner) Pos() *Pos {
	if s.SeeBlock() {
		panic("todo")
	}

	t := s.Tok()
	if t != nil {
		return t.Pos
	}
	if s.last != nil {
		return s.last.Pos
	}
	return nil
}
