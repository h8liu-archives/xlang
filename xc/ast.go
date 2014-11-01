package xc

import (
	"github.com/h8liu/xlang/parser"
)

// ASTNode basically could be anything
type ASTNode interface{}

// AST presents the xlang abstract syntax tree.
type AST struct {
	errs *parser.ErrList
	s    *parser.EntryScanner

	root  ASTNode
	scope *scope
	ir    *irBuilder
	obj   *Object
}

// ASTBlock is a scoped block.
type ASTBlock struct {
	Nodes []ASTNode
}

// ASTModule is a module that consists of top-level declarations.
type ASTModule struct {
}

func newAST() *AST {
	ret := new(AST)
	ret.errs = parser.NewErrList()
	return ret
}

func newStmtsAST(b *parser.Block) *AST {
	ret := newAST()
	root := new(ASTBlock)
	ret.parseStmts(root, b)
	ret.root = root

	return ret
}

func newExprsAST(b *parser.Block) *AST {
	ret := newAST()
	root := new(ASTBlock)
	ret.parseExprs(root, b)
	ret.root = root
	return ret
}

func (ast *AST) parseProg(ret *ASTModule, b *parser.Block) {
	panic("todo")
}

func (ast *AST) parseStmts(ret *ASTBlock, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty statement
		}

		ast.s = parser.NewEntryScanner(s)
		expr := ast.parseStmt()
		if ast.s.Entry() != nil {
			ast.errs.Log(ast.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}

func (ast *AST) parseExprs(ret *ASTBlock, b *parser.Block) {
	for _, s := range b.Stmts {
		if len(s) == 0 {
			continue // empty expr
		}

		ast.s = parser.NewEntryScanner(s)
		expr := ast.parseExpr()
		if ast.s.Entry() != nil {
			ast.errs.Log(ast.s.Pos(), "expect end of statement")
			continue
		}

		ret.Nodes = append(ret.Nodes, expr)
	}
}

func (ast *AST) buildFunc() {
	ast.scope = newScope()
	ast.scope.push()

	b := ast.root.(*ASTBlock)
	for _, s := range b.Nodes {
		ast.buildStmt(s)
	}

	ast.scope.pop()
}

func (ast *AST) buildExpr(s ASTNode) *enode {
	return nil
}

func (ast *AST) buildStmt(s ASTNode) {
	switch n := s.(type) {
	case *ASTAssign:
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

			destType := dest.Type()
			src := temps[i]
			srcType := src.Type()
			if !srcType.canAssignTo(destType) {
				ast.errs.Log(n.Pos, "cannot assign %s to %s", srcType, destType)
				return
			}

			ast.ir.addAssign(dest, src)
		}

	case *ASTVarDecl:
		var src *enode
		if n.Expr != nil {
			src = ast.buildExpr(n.Expr)
		}

		varName := n.Name.Lit
		pre := ast.scope.findTop(varName)
		if pre != nil {
			ast.errs.Log(n.Name.Pos, "%s already declared", n.Name.Lit)
			ast.errs.Log(pre.pos, "  previously declared here")
			return
		}

		typ := newBasicType("int")
		v := ast.newVar(varName, n.Name.Pos, typ)
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
			srcType := src.Type()
			destType := v.Type()

			if !srcType.canAssignTo(destType) {
				ast.errs.Log(n.Name.Pos, "cannot assign %s to %s", srcType, destType)
				return
			}

			ast.ir.addAssign(v, src)
		} else {
			ast.ir.addAssign(v, ast.makeZero(v.Type()))
		}
	}
}

func (ast *AST) newVar(name string, p *parser.Pos, t *xtype) *enode {
	panic("todo")
}

func (ast *AST) makeZero(t *xtype) *enode {
	panic("todo")
}
