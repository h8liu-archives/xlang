package xc

import (
	"fmt"
)

func (ast *AST) newVar(name string, t *xtype) *enode {
	ret := new(enode)
	ret.name = name
	ret.t = t
	ret.addr = ast.ir.stackAlloc(t.size(), ret)

	return ret
}

func (ast *AST) newTemp(t *xtype) *enode {
	ret := new(enode)
	ret.t = t
	ret.addr = ast.ir.stackAlloc(t.size(), ret)
	ret.name = fmt.Sprintf("#%d", ret.addr)

	return ret
}

func (ast *AST) newConst(t *xtype, v int32) *enode {
	ret := new(enode)
	ret.t = t
	if t.isInt {
		ret.isConst = true
		ret.value = v
	} else {
		panic("todo")
	}
	return ret
}

func (ast *AST) newZero(t *xtype) *enode {
	return ast.newConst(t, 0)
}
