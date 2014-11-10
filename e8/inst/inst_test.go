package inst_test

import (
	"testing"

	. "e8vm.net/e8/inst"
	"e8vm.net/e8/mem"
	"e8vm.net/e8/vm"
)

func TestSingleInst(t *testing.T) {
	c := vm.NewCore()
	p := mem.NewPage()
	c.Map(0, p)

	s := func(i Inst) {
		c.WriteReg(1, 0x1)
		c.WriteReg(2, 0x20)
		c.WriteReg(3, 0x300)
		c.WriteReg(4, 0)

		c.WriteReg(5, 0x31)
		c.WriteReg(6, 0xfffffff0)

		Exec(c, i)
	}

	r := Rinst
	rs := RinstShamt

	c4 := func(i Inst, v uint32) {
		s(i)
		if c.ReadReg(4) != v {
			t.Fail()
		}

		if c.ReadReg(0) != 0 {
			t.Fail()
		}
	}

	c4(r(0, 0, 4, FnAdd), 0)
	c4(r(2, 3, 4, FnAdd), 0x320)
	c4(r(3, 2, 4, FnAdd), 0x320)
	c4(r(3, 3, 4, FnAdd), 0x600)
	c4(r(2, 6, 4, FnAdd), 0x10)

	c4(r(0, 0, 4, FnSub), 0)
	c4(r(2, 0, 4, FnSub), 0x20)
	c4(r(2, 3, 4, FnSub), 0xfffffd20)
	c4(r(3, 2, 4, FnSub), 0x300-0x20)
	c4(r(3, 3, 4, FnSub), 0)
	c4(r(2, 6, 4, FnSub), 0x30)

	c4(r(0, 0, 4, FnAnd), 0)
	c4(r(2, 2, 4, FnAnd), 0x20)
	c4(r(5, 2, 4, FnAnd), 0x20)
	c4(r(2, 3, 4, FnAnd), 0)

	c4(rs(0, 2, 4, 2, FnSll), 0x20<<2)
	c4(rs(0, 2, 4, 0, FnSll), 0x20)
	c4(rs(0, 2, 4, 2, FnSrl), 0x20>>2)
	c4(rs(0, 2, 4, 0, FnSrl), 0x20)
	c4(rs(0, 1, 4, 2, FnSrl), 0)
	c4(rs(0, 6, 4, 2, FnSrl), 0x3ffffffc)
	c4(rs(0, 6, 4, 2, FnSra), 0xfffffffc)

	c4(r(1, 2, 4, FnSllv), 0x40)
	c4(r(2, 2, 4, FnSllv), 0)
	c4(r(0, 2, 4, FnSllv), 0x20)
	c4(r(1, 2, 4, FnSrlv), 0x10)
	c4(r(0, 2, 4, FnSrlv), 0x20)
	c4(r(1, 6, 4, FnSrlv), 0x7ffffff8)
	c4(r(1, 6, 4, FnSrav), 0xfffffff8)

	c4(r(2, 3, 4, FnXor), 0x320)
	c4(r(5, 2, 4, FnXor), 0x11)

	c4(r(2, 3, 4, FnOr), 0x320)
	c4(r(5, 2, 4, FnOr), 0x31)

	c4(r(2, 3, 4, FnNor), 0xffffffff-0x320)
	c4(r(5, 2, 4, FnNor), 0xffffffff-0x31)

	c4(r(5, 2, 4, FnSlt), 0)
	c4(r(2, 5, 4, FnSlt), 1)

	c4(r(2, 3, 4, FnMul), 0x6000)
}
