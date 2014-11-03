package xc

// a scope is just a stack of symbol lookup tables
type scope struct {
	stack []*symTable
}

func newScope() *scope {
	ret := new(scope)
	return ret
}

func (s *scope) push() {
	index := newSymTable()
	s.stack = append(s.stack, index)
}

func (s *scope) pop() {
	if len(s.stack) == 0 {
		panic("nothing to pop")
	}

	s.stack = s.stack[:len(s.stack)-1]
}

func (s *scope) top() *symTable {
	if len(s.stack) == 0 {
		panic("no top yet")
	}

	return s.stack[len(s.stack)-1]
}

func (s *scope) put(sym *symbol) bool {
	top := s.top()
	return top.put(sym)
}

func (s *scope) findTop(name string) *symbol {
	return s.top().find(name)
}

func (s *scope) find(name string) *symbol {
	n := len(s.stack)
	for i := n - 1; i >= 0; i-- {
		ind := s.stack[i]
		ret := ind.find(name)
		if ret != nil {
			return ret
		}
	}

	return nil
}
