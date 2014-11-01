package xc

type nameIndex struct {
	all map[string]interface{}
}

func (ind *nameIndex) put(name string, thing interface{}) bool {
	if thing == nil {
		panic("cannot put nil")
	}
	_, found := ind.all[name]
	if found {
		return false
	}
	ind.all[name] = thing
	return true
}

func (ind *nameIndex) find(name string) interface{} {
	return ind.all[name]
}

type scope struct {
	stack []*nameIndex
}

func newScope() *scope {
	ret := new(scope)
	return ret
}

func (s *scope) push() {
	index := new(nameIndex)
	s.stack = append(s.stack, index)
}

func (s *scope) pop() {
	if len(s.stack) == 0 {
		panic("nothing to pop")
	}

	s.stack = s.stack[:len(s.stack)-1]
}

func (s *scope) top() *nameIndex {
	if len(s.stack) == 0 {
		panic("no top yet")
	}

	return s.stack[len(s.stack)-1]
}

func (s *scope) put(name string, thing interface{}) bool {
	top := s.top()
	return top.put(name, thing)
}

func (s *scope) find(name string) interface{} {
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
