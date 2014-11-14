package jasm

import (
	"fmt"
)

func spf(f string, args ...interface{}) string {
	return fmt.Sprintf(f, args...)
}

func srcReg(i int) string {
	if i == 0 {
		return "(0|0)"
	}
	return spf("(%s|0)", destReg(i))
}

func rvar(i int) string {
	switch i {
	case 0:
		return "r0"
	case 1:
		return "r1"
	case 2:
		return "r2"
	case 3:
		return "r3"
	case 4:
		return "err"
	case 5:
		return "sp"
	case 6:
		return "ret"
	case 7:
		return "pc"
	default:
		panic("bug")
	}
}

func destReg(i int) string {
	return rvar(i)
}

func riGen1(i *ri, op string) string {
	return spf("%s = (%s %s %s) |0;", destReg(i.dest),
		srcReg(i.src1), op, srcReg(i.src2),
	)
}

func riGen2(i *ri, op string) string {
	return spf("%s = (%s %s (%s & 0xff)) |0;", destReg(i.dest),
		srcReg(i.src1), op, srcReg(i.src2),
	)
}

func riGen(i *ri) string {
	switch i.op {
	case "add":
		return riGen1(i, "+")
	case "sub":
		return riGen1(i, "-")
	case "and":
		return riGen1(i, "&")
	case "or":
		return riGen1(i, "|")
	case "xor":
		return riGen1(i, "^")
	case "slt":
		return riGen1(i, "<")
	case "mul":
		return spf("%s = stdlib.Math.imul(%s, %s);",
			destReg(i.dest), srcReg(i.src1), srcReg(i.src2))
	case "nor":
		return spf("%s = ((0xffffffff|0) ^ ((%s|%s)|0)) |0;",
			destReg(i.dest), srcReg(i.src1), srcReg(i.src2))
	case "sllv":
		return riGen2(i, "<<")
	case "srlv":
		return riGen2(i, ">>>")
	case "srav":
		return riGen2(i, ">>")
	default:
		panic("bug")
	}
}

func miGen(i *mi) string {
	switch i.op {
	case "lw":
		return spf("%s = memU32[((%s + %d)|0) >> 2] |0;",
			destReg(i.reg), srcReg(i.base), i.off)
	case "lb":
		return spf("%s = (((memU8[((%s + %d)|0)]|0) << 24) >>> 24) |0;",
			destReg(i.reg), srcReg(i.base), i.off)
	case "lbu":
		return spf("%s = memU8[((%s + %d)|0)]|0;",
			destReg(i.reg), srcReg(i.base), i.off)
	case "sw":
		return spf("memU32[((%s + %d)|0) >> 2] = %s|0",
			srcReg(i.base), i.off, srcReg(i.reg))
	case "sb":
		return spf("memU8[((%s + %d)|0)] = %s &0xff",
			srcReg(i.base), i.off, srcReg(i.reg))
	default:
		panic("bug")
	}
}

func biGen(i *bi) string {
	switch i.op {
	case "bne":
		return spf("if (%s != %s) { pc = (pc + %d)|0; }",
			srcReg(i.src1), srcReg(i.src2), i.off)
	case "beq":
		return spf("if (%s == %s) { pc = (pc + %d)|0; }",
			srcReg(i.src1), srcReg(i.src2), i.off)
	default:
		panic("bug")
	}
}

func jiGen(i *ji) string {
	switch i.op {
	case "j":
		return spf("pc = (pc + %d)|0;", i.off)
	case "jal":
		return spf("ret = pc|0; pc = (pc + %d)|0;", i.off)
	default:
		panic("bug")
	}
}
