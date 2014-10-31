package xc_test

import (
	"github.com/h8liu/xlang/xc"

	"testing"
)

func TestStmt(t *testing.T) {
	o := func(s string, e string) {
		src := xc.NewStrSource("test.xpr", s)
		b, errs := src.BuildStmtsAST()
		if errs != nil {
			t.Errorf("error on parsing: %q", s)
		} else if len(b.Nodes) != 1 {
			t.Errorf("not single expr: %q", s)
		} else {
			n := b.Nodes[0]
			res := xc.StmtStr(n)
			if res != e {
				t.Errorf("parsing expr %q, got %q, expect %q",
					s, res, e,
				)
			}
		}
	}

	e := func(s string) {
		src := xc.NewStrSource("test.xpr", s)
		_, errs := src.BuildStmtsAST()
		if errs == nil {
			t.Errorf("expect error on parsing: %q", s)
		}
	}

	o("var x = 3", "var x = 3")
	o("var x", "var x")
	o("var x = 3 + 4", "var x = (3+4)")
	o("println(3 + 4)", "println((3+4))")
	o("x = 3 + 4", "x = (3+4)")
	o("x = x", "x = x")

	e("var")
	e("var x =")
	e("var = ")
	e("var x = (3")
	e("x =")
	e("x = var")
}
