package inst

// Rinst makes an R instruction.
func Rinst(s, t, d, funct uint8) Inst {
	ret := uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(d) << 11
	ret |= uint32(funct)
	return Inst(ret)
}

// RinstShamt makes an R instruction with shift amount.
func RinstShamt(s, t, d, shamt, funct uint8) Inst {
	ret := uint32(Rinst(s, t, d, funct))
	ret |= uint32(shamt) << 6
	return Inst(ret)
}

// Iinst makes an I instruction.
func Iinst(op, s, t uint8, im uint16) Inst {
	ret := uint32(op) << 26
	ret |= uint32(s) << 21
	ret |= uint32(t) << 16
	ret |= uint32(im)
	return Inst(ret)
}

// Jinst makes an J instruction.
func Jinst(op uint8, off int32) Inst {
	ret := uint32(op) << 26
	ret |= uint32(off) & 0x3ffffff
	return Inst(ret)
}

// Noop makes an noop instruction.
func Noop() Inst {
	return Inst(0)
}
