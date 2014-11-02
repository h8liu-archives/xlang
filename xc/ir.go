package xc

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/prt"
)

// this is now like a basic block inside a function
// for no optimization, this is okay
// but if we want to add optimization later, we will need to change this
type irBlock struct {
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

func newIrBlock() *irBlock {
	ret := new(irBlock)
	return ret
}

func (b *irBlock) addInst(i interface{}) {
	b.insts = append(b.insts, i)
}

func (b *irBlock) addCall(dest, f *enode, args ...*enode) {
	b.addInst(&irCall{
		dest: dest,
		f:    f,
		args: args,
	})
}

func (b *irBlock) addAssign(dest, src *enode) {
	b.addUnaryOp(dest, "", src)
}

func (b *irBlock) addUnaryOp(dest *enode, op string, en *enode) {
	b.addInst(&irOp{
		dest: dest,
		a:    nil,
		b:    en,
		op:   op,
	})
}

func (b *irBlock) addBinaryOp(dest, x *enode, op string, y *enode) {
	b.addInst(&irOp{
		dest: dest,
		a:    x,
		b:    y,
		op:   op,
	})
}

func (b *irBlock) stackAlloc(size int32, en *enode) int32 {
	if size <= 0 {
		panic("bug")
	}
	ret := b.stackSize
	b.stackSize += size
	return ret
}

func (b *irBlock) stackReset() {
	b.stackSize = 0
}

func (b *irBlock) instStr(i interface{}) string {
	switch i := i.(type) {
	case *irOp:
		if i.a == nil {
			// unary op
			if i.op == "" {
				// just copying
				return fmt.Sprintf("%s = %s", i.dest, i.b)
			}
			return fmt.Sprintf("%s = %s %s", i.dest, i.op, i.b)
		} else {
			return fmt.Sprintf("%s = %s %s %s", i.dest, i.a, i.op, i.b)
		}

	case *irCall:
		ret := new(bytes.Buffer)
		if i.dest != nil {
			fmt.Fprintf(ret, "%s = ", i.dest)
		}

		fmt.Fprintf(ret, "call %s (", i.f)

		for i, arg := range i.args {
			if i > 0 {
				fmt.Fprintf(ret, ",")
			}
			fmt.Fprintf(ret, "%s", arg)
		}

		fmt.Fprintf(ret, ")")

		return ret.String()
	default:
		panic("bug")
	}
}

func (b *irBlock) PrintInsts(p *prt.Printer) {
	for _, i := range b.insts {
		p.Println(b.instStr(i))
	}
}
