package ir

type Func struct {
	blocks []*Block
	vars   []*Var
}

func NewFunc() *Func {
	ret := new(Func)
	return ret
}

// NewBlock creates a new basic block out of this function.
func (f *Func) NewBlock() *Block {
	ret := new(Block)
	ret.f = f

	f.blocks = append(f.blocks, ret)
	return ret
}

// StackAlloc allocates a new variable on the stack.
// It returns the frame offset.
func (f *Func) StackAlloc(size uint32) *Var {
	if size <= 0 {
		panic("bug")
	}

	ret := new(Var)
	ret.onHeap = false
	ret.size = size
	ret.index = len(f.vars)

	// TODO: stack alloc alignment
	if size != 4 {
		panic("not implemented")
	}

	f.vars = append(f.vars, ret)
	return ret
}
