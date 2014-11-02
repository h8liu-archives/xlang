package xc

type irBuilder struct {
	stackSize int32

	insts []interface{}
}

type irOp struct {
	dest *enode
	a    *enode
	b    *enode
	op   string
}

type irCall struct {
	dest *enode
	f    *enode
	args []*enode
}

func (b *irBuilder) addAssign(dest, src *enode) {
	panic("todo")
}

func (b *irBuilder) addInst(i interface{}) {
	b.insts = append(b.insts, i)
}

func (b *irBuilder) addCall(dest, f *enode, args ...*enode) {
	b.addInst(&irCall{
		dest: dest,
		f:    f,
		args: args,
	})
}

func (b *irBuilder) addUnaryOp(dest *enode, op string, en *enode) {
	b.addInst(&irOp{
		dest: dest,
		a:    nil,
		b:    en,
		op:   op,
	})
}

func (b *irBuilder) addBinaryOp(dest, x *enode, op string, y *enode) {
	b.addInst(&irOp{
		dest: dest,
		a:    x,
		b:    y,
		op:   op,
	})
}

func (b *irBuilder) stackAlloc(size int32, en *enode) int32 {
	if size <= 0 {
		panic("bug")
	}
	ret := b.stackSize
	b.stackSize += size
	return ret
}

func (b *irBuilder) stackReset() {
	b.stackSize = 0
}
