package xc

type symTable struct {
	all map[string]*symbol
}

func newSymTable() *symTable {
	ret := new(symTable)
	ret.all = make(map[string]*symbol)
	return ret
}

func (t *symTable) put(s *symbol) bool {
	if s == nil {
		panic("cannot put nil")
	}

	name := s.name
	_, found := t.all[name]
	if found {
		return false
	}
	t.all[name] = s
	return true
}

func (t *symTable) find(name string) *symbol {
	return t.all[name]
}
