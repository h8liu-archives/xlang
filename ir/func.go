package ir

import (
	"github.com/h8liu/xlang/prt"
)

// Func is an IR function unit.
type Func struct {
	blocks []*Block

	rets []*Var
	args []*Var
	vars []*Var
}

// NewFunc creates a new IR function unit.
func NewFunc() *Func {
	ret := new(Func)
	return ret
}

// NewBlock creates a new basic block out of this function.
func (f *Func) NewBlock() *Block {
	ret := new(Block)
	ret.f = f
	ret.index = len(f.blocks)

	f.blocks = append(f.blocks, ret)
	return ret
}

// StackAlloc allocates a new variable on the stack.
// It returns the frame offset.
func (f *Func) StackAlloc(size uint32) *Var {
	ret := f.stackAlloc(size)
	f.vars = append(f.vars, ret)
	return ret
}

func (f *Func) stackAlloc(size uint32) *Var {
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

	return ret
}

// Print prints the function.
func (f *Func) Print(p *prt.Printer) {
	for _, b := range f.blocks {
		b.Print(p)
	}
}
