package xc

import (
	"math"
	"strconv"

	"github.com/h8liu/xlang/parser"
)

var voidNode = &enode{isVoid: true, t: typeVoid}

func (ast *AST) prepareBuild() {
	ast.ir = newIrBlock()
	ast.scope = newScope()
	ast.scope.push() // buildin scope

	// TODO: fix this
	t := &xtype{isFunc: true}
	v := &enode{
		name:   "print",
		t:      t,
		onHeap: true,
		addr:   0x8000,
	}
	s := &symbol{
		name: "print",
		pos:  nil,
		typ:  t,
		v:    v,
	}
	ast.scope.put(s)
}

// builds a function
func (ast *AST) buildFunc() {
	ast.scope.push()

	b := ast.root.(*ASTBlock)
	for _, s := range b.Nodes {
		ast.buildStmt(s)
	}

	ast.scope.pop()

	ast.obj = new(Object)
	ast.obj.ir = ast.ir
}

// builds an expression
func (ast *AST) buildExpr(s ASTNode) *enode {
	switch n := s.(type) {
	case *ASTOpExpr:
		return ast.buildOp(n)
	case *ASTCall:
		return ast.buildCall(n)
	case *parser.Tok:
		if n.Type == parser.TypeIdent {
			return ast.buildVarRef(n)
		} else if n.Type == parser.TypeInt {
			return ast.buildIntConst(n)
		}
	}
	return nil
}

// builds an operation
func (ast *AST) buildOp(n *ASTOpExpr) *enode {
	const arithErr = "arithmetic operation on non-number"
	const typeErr = "type mismatch on operation %q"

	if n.A == nil {
		// unary op
		switch n.Op.Lit {
		case "+":
			ret := ast.buildExpr(n.B)
			if ret == nil {
				return nil
			}
			if !ret.typ().isNum() {
				ast.errs.Log(n.Op.Pos, arithErr)
				return nil
			}
			return ret
		case "-":
			b := ast.buildExpr(n.B)
			if b == nil {
				return nil
			}
			if !b.typ().isNum() {
				ast.errs.Log(n.Op.Pos, arithErr)
				return nil
			}
			ret := ast.newTemp(b.typ())
			ast.ir.addUnaryOp(ret, "-", b)
			return ret
		default:
			panic("unknown op")
		}
	}

	switch n.Op.Lit {
	case "+", "-":
		a := ast.buildExpr(n.A)
		b := ast.buildExpr(n.B)
		if a == nil || b == nil {
			return nil
		}
		if !a.typ().isNum() {
			ast.errs.Log(n.Op.Pos, arithErr)
			return nil
		}
		if !a.typ().numEquals(b.typ()) {
			ast.errs.Log(n.Op.Pos, typeErr, n.Op.Lit)
			return nil
		}

		ret := ast.newTemp(a.typ())
		ast.ir.addBinaryOp(ret, a, n.Op.Lit, b)
		return ret
	default:
		panic("unknown op")
	}
}

// builds a function call
func (ast *AST) buildCall(n *ASTCall) *enode {
	f := ast.buildExpr(n.Func)
	if f == nil {
		return nil
	}

	var args []*enode
	for _, p := range n.Paras {
		r := ast.buildExpr(p)
		if r == nil {
			return nil
		}
		args = append(args, r)
	}

	// TODO: function signature type check
	// we now assume it is always print with one parameter

	if len(n.Paras) != 1 {
		ast.errs.Log(n.Lparen.Pos, "print only accepts one paramter")
		return nil
	}

	ast.ir.addCall(voidNode, f, args...)
	return voidNode
}

// builds a variable reference
func (ast *AST) buildVarRef(t *parser.Tok) *enode {
	found := ast.scope.find(t.Lit)
	if found == nil {
		ast.errs.Log(t.Pos, "%s not defined", t.Lit)
		return nil
	}

	return found.v
}

// builds a integer constant.
// for integer within int32 range, the type is int32
// otherwise, for integer within uint32 range, the type is uint32
// otherwise, it is out of range and invalid
func (ast *AST) buildIntConst(t *parser.Tok) *enode {
	v, e := strconv.ParseInt(t.Lit, 0, 64)
	if e != nil {
		ast.errs.Log(t.Pos, "invalid integer")
		return nil
	}

	if v > math.MaxUint32 || v < math.MinInt32 {
		ast.errs.Log(t.Pos, "integer out of range")
		return nil
	}

	if v > math.MaxInt32 {
		return ast.newConst(typeUint, int32(v))
	}

	return ast.newConst(typeInt, int32(v))
}

// build assignment
func (ast *AST) buildAssign(n *ASTAssign) {
	nleft := len(n.LHS)
	nright := len(n.RHS)
	if nleft != nright {
		ast.errs.Log(n.Pos, "expect %d on left hand side, got %d",
			nright, nleft,
		)
		return
	}

	var temps []*enode
	for _, expr := range n.RHS {
		t := ast.buildExpr(expr)
		if t == nil {
			return
		}
		temps = append(temps, t)
	}

	for i, d := range n.LHS {
		dest := ast.buildExpr(d)
		if dest == nil {
			return
		}

		if !dest.addressable() {
			ast.errs.Log(n.Pos, "assigning to not addressable")
			return
		}

		destType := dest.typ()
		src := temps[i]
		srcType := src.typ()
		if !srcType.canAssignTo(destType) {
			ast.errs.Log(n.Pos, "cannot assign %s to %s", srcType, destType)
			return
		}

		ast.ir.addAssign(dest, src)
	}
}

// build variable declaration
func (ast *AST) buildVarDecl(n *ASTVarDecl) {
	var src *enode
	if n.Expr != nil {
		src = ast.buildExpr(n.Expr)
		// still build the symbol if src is nil
	}

	varName := n.Name.Lit
	pre := ast.scope.findTop(varName)
	if pre != nil {
		ast.errs.Log(n.Name.Pos, "%s already declared", n.Name.Lit)
		ast.errs.Log(pre.pos, "  previously declared here")
		return
	}

	typ := typeInt // TODO: parse the type
	v := ast.newVar(varName, typ)
	sym := &symbol{
		name: varName,
		pos:  n.Name.Pos,
		typ:  typ,
		v:    v,
	}
	ast.scope.put(sym)

	if n.Expr != nil {
		if src == nil {
			return
		}
		srcType := src.typ()
		destType := v.typ()

		if !srcType.canAssignTo(destType) {
			ast.errs.Log(n.Name.Pos, "cannot assign %s to %s", srcType, destType)
			return
		}

		ast.ir.addAssign(v, src)
	} else {
		ast.ir.addAssign(v, ast.newZero(v.typ()))
	}
}

func (ast *AST) buildExprStmt(n *ASTExprStmt) {
	ast.buildExpr(n.Expr)
}

// build a statement
func (ast *AST) buildStmt(s ASTNode) {
	switch n := s.(type) {
	case *ASTAssign:
		ast.buildAssign(n)
	case *ASTVarDecl:
		ast.buildVarDecl(n)
	case *ASTExprStmt:
		ast.buildExprStmt(n)
	default:
		panic("invalid statement")
	}
}
