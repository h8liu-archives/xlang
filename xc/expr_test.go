package xc_test

import (
	"github.com/h8liu/xlang/xc"

	"testing"
)

func TestExpr(t *testing.T) {
	o := func(s string, e string) {
		src := xc.NewStrSource("test.xpr", s)
		b, errs := src.BuildExprsAST()
		if errs != nil {
			t.Errorf("error on parsing: %q", s)
		} else if len(b.Nodes) != 1 {
			t.Errorf("not single expr: %q", s)
		} else {
			n := b.Nodes[0]
			res := xc.ExprStr(n)
			if res != e {
				t.Errorf("parsing expr %q, got %q, expect %q",
					s, res, e,
				)
			}
		}
	}

	e := func(s string) {
		src := xc.NewStrSource("test.xpr", s)
		_, errs := src.BuildExprsAST()
		if errs == nil {
			t.Errorf("expect error on parsing: %q", s)
		}
	}

	o("3", "3")
	o("3 // some comment", "3")
	o("3+4", "(3+4)")
	o("3 + 4", "(3+4)")
	o("3 + /* hello */ 4", "(3+4)")
	o("a", "a")
	o("a;", "a")
	o("(a)", "a")
	o("(((a)))", "a")
	o("a-b-c", "((a-b)-c)")
	o("a-(b-c)", "(a-(b-c))")
	o("_3", "_3")

	e("/*")
	e("(")
	e("3a")
	e("())")
	e(")")
	e("3 a")
	e("a b")
	e("_ 3")
}
