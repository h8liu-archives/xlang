package xc

import (
	"math"
	"strconv"

	"github.com/h8liu/xlang/ir"
	"github.com/h8liu/xlang/parser"
	"github.com/h8liu/xlang/xast"
)

var voidNode = &enode{t: typeVoid, v: ir.Void}

type builder struct {
	t     *xast.Tree
	errs  *parser.ErrList
	scope *scope
	f     *ir.Func
	b     *ir.Block
	obj   *Object
}

func (b *builder) prepare() {
	b.errs = parser.NewErrList()
	b.f = ir.NewFunc()
	b.b = b.f.NewBlock()
	b.scope = newScope()
	b.scope.push() // buildin scope

	// TODO: fix this
	t := newFuncType(typeVoid, typeInt)
	v := &enode{
		name: "print",
		t:    t,
		v:    ir.NewConstSym("<builtin>", "print"),
	}
	s := &symbol{
		name: "print",
		pos:  nil,
		v:    v,
	}
	b.scope.put(s)
}

// builds a function
func (b *builder) buildFunc() {
	b.scope.push()

	block := b.t.Root().(*xast.Block)
	for _, s := range block.Nodes {
		b.buildStmt(s)
	}

	b.scope.pop()

	b.obj = new(Object)
	b.obj.f = b.f
	b.obj.b = b.b
}

// builds an expression
func (b *builder) buildExpr(s xast.Node) *enode {
	switch n := s.(type) {
	case *xast.OpExpr:
		return b.buildOp(n)
	case *xast.Call:
		return b.buildCall(n)
	case *parser.Tok:
		if n.Type == parser.TypeIdent {
			return b.buildVarRef(n)
		} else if n.Type == parser.TypeInt {
			return b.buildIntConst(n)
		}
	}
	return nil
}

// builds an operation
func (b *builder) buildOp(n *xast.OpExpr) *enode {
	const arithErr = "arithmetic operation on non-number"
	const typeErr = "type mismatch on operation %q"

	if n.A == nil {
		// unary op
		switch n.Op.Lit {
		case "+":
			ret := b.buildExpr(n.B)
			if ret == nil {
				return nil
			}
			if !ret.typ().isNum() {
				b.errs.Log(n.Op.Pos, arithErr)
				return nil
			}
			return ret
		case "-":
			y := b.buildExpr(n.B)
			if y == nil {
				return nil
			}
			if !y.typ().isNum() {
				b.errs.Log(n.Op.Pos, arithErr)
				return nil
			}
			ret := b.newTemp(y.typ())
			b.b.AddUnaryOp(ret.v, "-", y.v)
			return ret
		default:
			panic("unknown op")
		}
	}

	switch n.Op.Lit {
	case "+", "-":
		x := b.buildExpr(n.A)
		y := b.buildExpr(n.B)
		if x == nil || y == nil {
			return nil
		}
		if !x.typ().isNum() {
			b.errs.Log(n.Op.Pos, arithErr)
			return nil
		}
		if !x.typ().numEquals(y.typ()) {
			b.errs.Log(n.Op.Pos, typeErr, n.Op.Lit)
			return nil
		}

		ret := b.newTemp(x.typ())
		b.b.AddBinaryOp(ret.v, x.v, n.Op.Lit, y.v)
		return ret
	default:
		panic("unknown op")
	}
}

// builds a function call
func (b *builder) buildCall(n *xast.Call) *enode {
	f := b.buildExpr(n.Func)
	if f == nil {
		return nil
	}

	var args []*ir.Var
	for _, p := range n.Paras {
		r := b.buildExpr(p)
		if r == nil {
			return nil
		}
		args = append(args, r.v)
	}

	// TODO: function signature type check
	// we now assume it is always print with one parameter
	if len(n.Paras) != 1 {
		b.errs.Log(n.Lparen.Pos, "print takes exactly one paramter")
		return nil
	}

	b.b.AddCall(ir.Void, f.v, args...)

	return voidNode
}

// builds a variable reference
func (b *builder) buildVarRef(t *parser.Tok) *enode {
	found := b.scope.find(t.Lit)
	if found == nil {
		b.errs.Log(t.Pos, "%s not defined", t.Lit)
		return nil
	}

	return found.v
}

// builds a integer constant.
// for integer within int32 range, the type is int32
// otherwise, for integer within uint32 range, the type is uint32
// otherwise, it is out of range and invalid
func (b *builder) buildIntConst(t *parser.Tok) *enode {
	v, e := strconv.ParseInt(t.Lit, 0, 64)
	if e != nil {
		b.errs.Log(t.Pos, "invalid integer")
		return nil
	}

	if v > math.MaxUint32 || v < math.MinInt32 {
		b.errs.Log(t.Pos, "integer out of range")
		return nil
	}

	if v > math.MaxInt32 {
		return newConst(typeUint, int32(v))
	}

	return newConst(typeInt, int32(v))
}

// build assignment
func (b *builder) buildAssign(n *xast.Assign) {
	nleft := len(n.LHS)
	nright := len(n.RHS)
	if nleft != nright {
		b.errs.Log(n.Pos, "expect %d on left hand side, got %d",
			nright, nleft,
		)
		return
	}

	var temps []*enode
	for _, expr := range n.RHS {
		t := b.buildExpr(expr)
		if t == nil {
			return
		}
		temps = append(temps, t)
	}

	for i, d := range n.LHS {
		dest := b.buildExpr(d)
		if dest == nil {
			return
		}

		if !dest.addressable() {
			b.errs.Log(n.Pos, "assigning to not addressable")
			return
		}

		destType := dest.typ()
		src := temps[i]
		srcType := src.typ()
		if !srcType.canAssignTo(destType) {
			b.errs.Log(n.Pos, "cannot assign %s to %s", srcType, destType)
			return
		}

		b.b.AddAssign(dest.v, src.v)
	}
}

// build variable declaration
func (b *builder) buildVarDecl(n *xast.VarDecl) {
	var srcs []*enode

	// evaluate the expressions first
	if n.Exprs != nil {
		if len(n.Exprs) != len(n.Names) {
			b.errs.Log(n.Pos, "number of expressions mismatch")
		} else {
			srcs = make([]*enode, len(n.Exprs))
			for i, expr := range n.Exprs {
				e := b.buildExpr(expr)
				if e == nil {
					srcs = nil
					break
				}
				srcs[i] = e
			}
		}
	}

	// now we declare the names
	for i, name := range n.Names {
		pre := b.scope.findTop(name.Lit)
		if pre != nil {
			if pre.pos == nil {
				panic("trying redeclare a builtin symbol?")
			}
			b.errs.Log(name.Pos, "%s already declared", name.Lit)
			b.errs.Log(pre.pos, "  previously declared here")
			return
		}

		typ := typeInt // TODO: parse the type
		v := b.newVar(name.Lit, typ)
		sym := &symbol{
			name: name.Lit,
			pos:  name.Pos,
			v:    v,
		}
		b.scope.put(sym)

		if srcs != nil {
			// have init list, then assign it
			src := srcs[i]
			srcType := src.typ()
			destType := v.typ()

			if !srcType.canAssignTo(destType) {
				b.errs.Log(name.Pos, "cannot assign %s to %s", srcType, destType)
				return
			}

			b.b.AddAssign(v.v, src.v)
		} else {
			// init with zero value
			b.b.AddAssign(v.v, newZero(v.typ()).v)
		}
	}
}

func (b *builder) buildExprStmt(n *xast.ExprStmt) {
	b.buildExpr(n.Expr)
	// TODO: check if the expression is a call
}

// build a statement
func (b *builder) buildStmt(s xast.Node) {
	switch n := s.(type) {
	case *xast.Assign:
		b.buildAssign(n)
	case *xast.VarDecl:
		b.buildVarDecl(n)
	case *xast.ExprStmt:
		b.buildExprStmt(n)
	default:
		panic("invalid statement")
	}
}
