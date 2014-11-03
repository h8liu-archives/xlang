package xc

import (
	"github.com/h8liu/xlang/ir"
)

// this enode thing has two parts.
// first, it is a memory object on the stack or the heap.
// second, it has a type descriptor
// we sould move the memory object part to the ir, because the IR optimizer could
// futher change that
type enode struct {
	name string // this is just for debugging
	t    *xtype
	v    *ir.Var
}

func (n *enode) typ() *xtype {
	return n.t
}

func (n *enode) addressable() bool {
	return n.v.IsAddressable()
}
