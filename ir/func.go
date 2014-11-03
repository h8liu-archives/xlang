package ir

type Func struct {
	frameSize uint32

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
	ret.addr = f.frameSize

	// TODO: stack alloc alignment
	if size != 4 {
		panic("not implemented")
	}

	f.frameSize += size
	f.vars = append(f.vars, ret)
	return ret
}
