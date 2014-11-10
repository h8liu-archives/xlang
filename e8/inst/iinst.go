package inst

func memAddr(c Core, i Inst) uint32 {
	s := c.ReadReg(i.Rs())
	return s + signExt(i.Imu())
}

func signExt(im uint16) uint32 {
	return uint32(int32(int16(im)))
}

func unsignExt(im uint16) uint32 {
	return uint32(im)
}

func signExt8(a uint8) uint32 {
	return uint32(int32(int8(a)))
}

func unsignExt8(a uint8) uint32 {
	return uint32(a)
}

func opAddi(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	c.WriteReg(i.Rt(), s+signExt(i.Imu()))
}

func opLw(c Core, i Inst) {
	addr := memAddr(c, i)
	c.WriteReg(i.Rt(), c.ReadU32(addr))
}

func opLh(c Core, i Inst) {
	addr := memAddr(c, i)
	c.WriteReg(i.Rt(), signExt(c.ReadU16(addr)))
}

func opLhu(c Core, i Inst) {
	addr := memAddr(c, i)
	c.WriteReg(i.Rt(), unsignExt(c.ReadU16(addr)))
}

func opLb(c Core, i Inst) {
	addr := memAddr(c, i)
	c.WriteReg(i.Rt(), signExt8(c.ReadU8(addr)))
}

func opLbu(c Core, i Inst) {
	addr := memAddr(c, i)
	c.WriteReg(i.Rt(), unsignExt8(c.ReadU8(addr)))
}

func opSw(c Core, i Inst) {
	t := c.ReadReg(i.Rt())
	addr := memAddr(c, i)
	c.WriteU32(addr, t)
}

func opSh(c Core, i Inst) {
	t := uint16(c.ReadReg(i.Rt()))
	addr := memAddr(c, i)
	c.WriteU16(addr, t)
}

func opSb(c Core, i Inst) {
	t := uint8(c.ReadReg(i.Rt()))
	addr := memAddr(c, i)
	c.WriteU8(addr, t)
}

func opLui(c Core, i Inst) {
	c.WriteReg(i.Rt(), unsignExt(i.Imu())<<16)
}

func opAndi(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	c.WriteReg(i.Rt(), s&unsignExt(i.Imu()))
}

func opOri(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	c.WriteReg(i.Rt(), s|unsignExt(i.Imu()))
}

func opSlti(c Core, i Inst) {
	s := int32(c.ReadReg(i.Rs()))
	if s < int32(signExt(i.Imu())) {
		c.WriteReg(i.Rt(), 1)
	} else {
		c.WriteReg(i.Rt(), 0)
	}
}

func opBeq(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	if s == t {
		pc := c.ReadReg(RegPC)
		c.WriteReg(RegPC, pc+(signExt(i.Imu())<<2))
	}
}

func opBne(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	if s != t {
		pc := c.ReadReg(RegPC)
		c.WriteReg(RegPC, pc+(signExt(i.Imu())<<2))
	}
}
