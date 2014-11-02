package parser

// EntryScanner is helps on parsing a statement
type EntryScanner struct {
	s     Stmt
	index int

	accepted *Tok
	last     *Tok
}

// NewEntryScanner creates a new EntryScanner
func NewEntryScanner(s Stmt) *EntryScanner {
	ret := new(EntryScanner)
	ret.s = s
	return ret
}

func (s *EntryScanner) next() {
	s.index++

	if s.index >= len(s.s) {
		if s.index > len(s.s) {
			panic("EntryScanner over running")
		}

		entry := s.s[s.index-1]
		if entry.Block != nil {
			s.last = entry.Block.Lbrace
		} else {
			s.last = entry.Tok
		}
	}
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

// SeeKeyword returns true if the current token is one of the keywords.
func (s *EntryScanner) SeeKeyword(kws ...string) bool {
	if !s.See(TypeKeyword) {
		return false
	}

	t := s.Tok()
	for _, kw := range kws {
		if t.Lit == kw {
			return true
		}
	}

	return false
}

// SeeIdent is just a shortcut for s.See(TypeIdent).
func (s *EntryScanner) SeeIdent() bool {
	return s.See(TypeIdent)
}

// AcceptOp returns true and accepts the current token if the
// current token is the exact op. It returns false and is
// an nop otherwise.
func (s *EntryScanner) AcceptOp(op string) bool {
	if !s.SeeOp(op) {
		return false
	}

	s.accepted = s.Entry().Tok
	s.next()
	return true
}

// AcceptKeyword returns true and accepts the current token if
// the current token is the exact keyword. It returns false and
// is an nop otherwise.
func (s *EntryScanner) AcceptKeyword(kw string) bool {
	if !s.SeeKeyword(kw) {
		return false
	}

	s.accepted = s.Entry().Tok
	s.next()
	return true
}

// Accept returns and accepts the current token entry.
// It panics if the current entry is a block.
func (s *EntryScanner) Accept() *Tok {
	ret := s.Tok()
	s.accepted = ret
	s.next()
	return ret
}

// Accepted returns the token that is just accepted.
func (s *EntryScanner) Accepted() *Tok {
	return s.accepted
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

// SeeBlock returns true if the current entry is a block.
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
	entry := s.Entry()
	if entry != nil {
		if entry.Block != nil {
			return entry.Block.Lbrace.Pos
		}
		return entry.Tok.Pos
	}
	if s.last != nil {
		return s.last.Pos
	}
	return nil
}
