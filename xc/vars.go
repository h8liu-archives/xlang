package xc

import (
	"github.com/h8liu/xlang/ir"
)

// newVar allocates a named variable from the current function's stack
func (b *builder) newVar(name string, t *xtype) *enode {
	ret := b.newTemp(t)

	ret.v.Name = name
	ret.name = name

	return ret
}

// newTemp allocates an unamed temp variable from the current function's stack
func (b *builder) newTemp(t *xtype) *enode {
	ret := new(enode)
	ret.t = t
	ret.v = b.f.StackAlloc(t.size())

	return ret
}

// newConst creates a constant enode out of thin air
func newConst(t *xtype, v int32) *enode {
	ret := new(enode)
	ret.t = t
	if t.isInt {
		ret.v = ir.NewConst(v)
	} else {
		panic("todo")
	}
	return ret
}

// newZero creates a zero value constant enode out of thin air
func newZero(t *xtype) *enode {
	return newConst(t, 0)
}
