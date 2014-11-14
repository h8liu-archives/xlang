package jasm

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
				rname(i.reg), rname(i.base),
			)
		} else {
			return spf("%s %s, %d(%s)", i.op,
				rname(i.reg), i.off, rname(i.base),
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
	case "andi", "ori", "sll", "srl", "sra":
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
