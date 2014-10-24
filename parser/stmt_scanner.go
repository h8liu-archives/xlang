package parser

type stmtScanner struct {
	stmt stmt

	cur  int
	hold *entry
}

var _ Stmt = new(stmtScanner)

func newStmtScanner(s stmt) *stmtScanner {
	ret := new(stmtScanner)
	ret.stmt = s
	return ret
}

func (s *stmtScanner) Scan() bool {
	ret := s.cur < len(s.stmt)
	if ret {
		s.hold = &s.stmt[s.cur]
	} else {
		s.hold = nil
	}
	s.cur++

	return ret
}

func (s *stmtScanner) IsBlock() bool {
	return s.hold.block != nil
}

func (s *stmtScanner) Block() Block {
	return newBlockScanner(s.hold.block)
}

func (s *stmtScanner) Tok() *Tok {
	return s.hold.tok
}
