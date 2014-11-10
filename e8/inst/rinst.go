package inst

func opAdd(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s+t)
}

func opSub(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s-t)
}

func opMul(c Core, i Inst) {
	s := int32(c.ReadReg(i.Rs()))
	t := int32(c.ReadReg(i.Rt()))
	c.WriteReg(i.Rd(), uint32(s*t))
}

func opMulu(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s*t)
}

func opDiv(c Core, i Inst) {
	s := int32(c.ReadReg(i.Rs()))
	t := int32(c.ReadReg(i.Rt()))
	if t == 0 {
		c.WriteReg(i.Rd(), 0)
	} else {
		c.WriteReg(i.Rd(), uint32(s/t))
	}
}

func opMod(c Core, i Inst) {
	s := int32(c.ReadReg(i.Rs()))
	t := int32(c.ReadReg(i.Rt()))
	if t == 0 {
		c.WriteReg(i.Rd(), 0)
	} else {
		c.WriteReg(i.Rd(), uint32(s%t))
	}
}

func opDivu(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	if t == 0 {
		c.WriteReg(i.Rd(), 0)
	} else {
		c.WriteReg(i.Rd(), s/t)
	}
}

func opModu(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	if t == 0 {
		c.WriteReg(i.Rd(), 0)
	} else {
		c.WriteReg(i.Rd(), s%t)
	}
}

func opAnd(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s&t)
}

func opOr(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s|t)
}

func opXor(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), s^t)
}

func opNor(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), ^(s | t))
}

func opSlt(c Core, i Inst) {
	s := int32(c.ReadReg(i.Rs()))
	t := int32(c.ReadReg(i.Rt()))
	if s < t {
		c.WriteReg(i.Rd(), 1)
	} else {
		c.WriteReg(i.Rd(), 0)
	}
}

func opSll(c Core, i Inst) {
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), t<<i.Sh())
}

func opSrl(c Core, i Inst) {
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), t>>i.Sh())
}

func opSra(c Core, i Inst) {
	t := int32(c.ReadReg(i.Rt()))
	c.WriteReg(i.Rd(), uint32(t>>i.Sh()))
}

func opSllv(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), t<<s)
}

func opSrlv(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := c.ReadReg(i.Rt())
	c.WriteReg(i.Rd(), t>>s)
}

func opSrav(c Core, i Inst) {
	s := c.ReadReg(i.Rs())
	t := int32(c.ReadReg(i.Rt()))
	c.WriteReg(i.Rd(), uint32(t>>s))
}
