package jasm

type inst interface{}
type symInst struct {
	inst
	sym *sym
}

type sym struct {
	labs []string
}

type ri struct {
	op         string
	dest       int
	src1, src2 int
}

type mi struct {
	op   string
	dest int
	base int
	off  int32
}

type ii struct {
	op   string
	dest int
	src  int
	im   int32
}

type bi struct {
	op         string
	src1, src2 int
	off        int32
}

type ji struct {
	op  string
	off int32
}

func rname(i int) string {
	switch i {
	case 0:
		return "$0"
	case 1:
		return "$1"
	case 2:
		return "$2"
	case 3:
		return "$3"
	case 4:
		return "$err"
	case 5:
		return "$sp"
	case 6:
		return "$ret"
	case 7:
		return "$pc"
	default:
		panic("bug")
	}
}

func riStr(i *ri) string {
	switch i.op {
	case "add", "sub", "and", "or", "xor", "nor", "slt",
		"mul", "sllv", "srlv", "srav":
		return spf("%s %s, %s, %s", i.op,
			rname(i.dest), rname(i.src1), rname(i.src2),
		)
	default:
		panic("bug")
	}
}

func miStr(i *mi) string {
	switch i.op {
	case "lw", "lb", "lbu", "sw", "sb":
		if i.off == 0 {
			return spf("%s %s, (%s)", i.op,
				rname(i.dest), rname(i.base),
			)
		} else {
			return spf("%s %s, %d(%s)", i.op,
				rname(i.dest), i.off, rname(i.base),
			)
		}
	default:
		panic("bug")
	}
}

func iiStr(i *ii) string {
	switch i.op {
	case "addi", "slti":
		return spf("%s %s, %s, %d", i.op,
			rname(i.dest), rname(i.src), i.im,
		)
	case "andi", "ori":
		return spf("%s %s, %s, 0x%x", i.op,
			rname(i.dest), rname(i.src), uint32(i.im),
		)
	default:
		panic("bug")
	}
}

func biStr(i *bi) string {
	switch i.op {
	case "bne", "beq":
		return spf("%s %s, %s, %d", i.op,
			rname(i.src1), rname(i.src2), i.off,
		)
	default:
		panic("bug")
	}
}

func jiStr(i *ji) string {
	switch i.op {
	case "j", "jal":
		return spf("%s %d", i.op, i.off)
	default:
		panic("bug")
	}
}

func iStr(i inst) string {
	switch i := i.(type) {
	case *ri:
		return riStr(i)
	case *mi:
		return miStr(i)
	case *ii:
		return iiStr(i)
	case *bi:
		return biStr(i)
	case *ji:
		return jiStr(i)
	default:
		panic("bug")
	}
}
