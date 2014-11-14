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
	return destReg(i)
}

func destReg(i int) string {
	return spf("(%s|0)", rname(i))
}

func riGen1(i *ri, op string) string {
	return spf("%s = (%s %s %s) |0;", destReg(i.dest),
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
	default:
		panic("bug")
	}
}
