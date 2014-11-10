package inst

import (
	"fmt"
)

func (i Inst) String() string {
	if uint32(i) == 0 {
		return "noop"
	}

	op := i.Op()

	if op == OpRinst {
		rs := i.Rs()
		rt := i.Rt()
		rd := i.Rd()
		shamt := i.Sh()
		funct := i.Fn()
		name := FunctName(funct)

		r3 := func() string {
			return fmt.Sprintf("%s $%d, $%d, $%d", name, rd, rs, rt)
		}
		r3r := func() string {
			return fmt.Sprintf("%s $%d, $%d, $%d", name, rd, rt, rs)
		}
		r3s := func() string {
			return fmt.Sprintf("%s $%d, $%d, $%d", name, rd, rt, shamt)
		}

		switch funct {
		case FnAdd, FnSub, FnAnd, FnOr, FnXor, FnNor, FnSlt,
			FnMul, FnMulu, FnDiv, FnDivu, FnMod, FnModu:
			return r3()
		case FnSll, FnSrl, FnSra:
			return r3s()
		case FnSllv, FnSrlv, FnSrav:
			return r3r()
		}

		return fmt.Sprintf("noop-r%d", funct)
	} else if op == OpJal {
		return fmt.Sprintf("jal %d", i.Off())
	} else if op == OpJ {
		return fmt.Sprintf("j %d", i.Off())
	}

	// Funct

	rs := i.Rs()
	rt := i.Rt()
	im := i.Imu()
	ims := i.Ims()
	name := OpName(op)

	i2 := func() string {
		return fmt.Sprintf("%s $%d, %d", name, rt, im)
	}
	i3sr := func() string {
		return fmt.Sprintf("%s $%d, $%d, %d", name, rs, rt, ims)
	}
	i3s := func() string {
		return fmt.Sprintf("%s $%d, $%d, %d", name, rt, rs, ims)
	}
	i3 := func() string {
		return fmt.Sprintf("%s $%d, $%d, %d", name, rt, rs, im)
	}
	i3a := func() string {
		return fmt.Sprintf("%s $%d, %d($%d)", name, rt, ims, rs)
	}

	switch op {
	case OpBeq, OpBne:
		return i3sr()
	case OpAndi, OpOri:
		return i3()
	case OpAddi, OpSlti:
		return i3s()
	case OpLui:
		return i2()
	case OpLw, OpLh, OpLhu, OpLb, OpLbu, OpSw, OpSh, OpSb:
		return i3a()
	}

	return fmt.Sprintf("noop-%d", op)
}
