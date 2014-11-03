package xc

import (
	"github.com/h8liu/xlang/ir"
)

func (ast *AST) newVar(name string, t *xtype) *enode {
	ret := new(enode)
	ret.name = name
	ret.t = t

	ret.v = ast.f.StackAlloc(t.size())
	ret.v.Name = name

	return ret
}

func (ast *AST) newTemp(t *xtype) *enode {
	ret := new(enode)
	ret.t = t
	ret.v = ast.f.StackAlloc(t.size())

	return ret
}

func (ast *AST) newConst(t *xtype, v int32) *enode {
	ret := new(enode)
	ret.t = t
	if t.isInt {
		ret.v = ir.NewConst(v)
	} else {
		panic("todo")
	}
	return ret
}

func (ast *AST) newZero(t *xtype) *enode {
	return ast.newConst(t, 0)
}
