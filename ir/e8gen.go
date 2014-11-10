package ir

type E8Gen struct {
	retAddr   *Var
	frameSize uint32
}

func NewE8Gen() *E8Gen {
	ret := new(E8Gen)
	return ret
}

/*
E8 calling convention:

- $31 is program counter
- $30 is the return address
- $29 is the stack pointer
- $1-3 is the first three arguments and return values
- others are temps

when calling
[sp] = return address
[sp+4] = return value (if not void)
... 3-n other arguments

*/

const (
	e8AddrSize = 4
)

func (g *E8Gen) GenFunc(f *Func) {
	if len(f.blocks) == 0 {
		return
	}

	// $30 is stack counter
	g.frameSize = g.arrangeStack(f)
	g.genFuncPrologue(f)

	for _, b := range f.blocks {
		g.genBlock(b)
	}
}

func (g *E8Gen) arrangeStack(f *Func) uint32 {
	g.retAddr = f.StackAlloc(e8AddrSize) // allocate the return address

	offset := uint32(0)
	push := func(v *Var) {
		v.addr = offset
		offset += v.size
	}

	// parameters that were pushed on stack
	// by the caller
	push(g.retAddr)
	if len(f.rets) > 3 {
		for _, v := range f.rets[3:] {
			push(v)
		}
	}
	if len(f.args) > 3 {
		for _, v := range f.rets[3:] {
			push(v)
		}
	}

	for _, v := range f.rets[:3] {
		push(v)
	}
	for _, v := range f.args[:3] {
		push(v)
	}
	for _, v := range f.vars {
		push(v)
	}

	return offset
}

func (g *E8Gen) genFuncPrologue(f *Func) {
	panic("todo")
}

func (g *E8Gen) genBlock(b *Block) {
	for _, i := range b.insts {
		switch i := i.(type) {
		case *oper:
			g.genOp(i)
		case *call:
			g.genCall(i)
		default:
			panic("bug")
		}
	}
}

func (g *E8Gen) genOp(i *oper) {
	panic("todo")
}

func (g *E8Gen) genCall(i *call) {

}
