package parser

type blockScanner struct {
	block block

	cur  int
	hold stmt
}

var _ Block = new(blockScanner)

func newBlockScanner(b block) *blockScanner {
	ret := new(blockScanner)
	ret.block = b

	return ret
}

func (s *blockScanner) Scan() bool {
	ret := s.cur < len(s.block)
	if ret {
		s.hold = s.block[s.cur]
	} else {
		s.hold = nil
	}
	s.cur++

	return ret
}

func (s *blockScanner) Stmt() Stmt {
	if s.hold == nil {
		panic("invalid scan")
	}
	return newStmtScanner(s.hold)
}
