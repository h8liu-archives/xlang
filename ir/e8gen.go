package ir

import (
	"fmt"
	"io"

	e8i "github.com/h8liu/xlang/e8/inst"
)

type e8Inst struct {
	i uint32

	lateBind         bool
	bindAddr         uint32
	bindMod, bindSym string
}

func newE8Inst(i uint32) *e8Inst {
	ret := new(e8Inst)
	ret.i = i
	return ret
}

type E8Gen struct {
	retAddr   *Var
	frameSize uint32

	insts []*e8Inst
}

func NewE8Gen() *E8Gen {
	ret := new(E8Gen)
	return ret
}

/*
E8 calling convention:

- $31 is program counter
- $30 is the return address
- $29 is the stack pointer
- $1-3 is the first three arguments and return values
- others are temps

when calling
[sp] = return address
[sp+4] = return value (if not void)
... 3-n other arguments

*/

const (
	e8AddrSize = 4

	regSP  = 29
	regT1  = 4
	regT2  = 5
	regRET = 30
	regPC  = 31
)

func (g *E8Gen) addInst(i uint32) {
	inst := newE8Inst(i)
	g.insts = append(g.insts, inst)
}

func (g *E8Gen) storeReg(r uint8, v *Var) {
	if !v.isConst {
		g.addInst(uint32(e8i.Iinst(e8i.OpSw, regSP, r, uint16(v.addr))))
	} else {
		panic("storing const")
	}
}

func (g *E8Gen) loadReg(r uint8, v *Var) {
	if !v.isConst {
		g.addInst(uint32(e8i.Iinst(e8i.OpLw, regSP, r, uint16(v.addr))))
	} else {
		i := v.value
		u := uint32(i) >> 16
		if u == 0 {
			g.addInst(uint32(e8i.Iinst(e8i.OpAddi, 0, r, uint16(i))))
		} else {
			g.addInst(uint32(e8i.Iinst(e8i.OpLui, 0, r, uint16(u))))
			g.addInst(uint32(e8i.Iinst(e8i.OpAddi, r, r, uint16(i))))
		}
	}
}

func (g *E8Gen) addRinst(s, t, d, funct uint8) {
	g.addInst(uint32(e8i.Rinst(s, t, d, funct)))
}

func (g *E8Gen) addSPChange(d int16) {
	g.addInst(uint32(e8i.Iinst(e8i.OpAddi, regSP, regSP, uint16(d))))
}

func (g *E8Gen) addJalSym(mod, sym string) {
	inst := newE8Inst(uint32(e8i.Jinst(e8i.OpJal, 0)))

	inst.lateBind = true
	inst.bindMod = mod
	inst.bindSym = sym

	g.insts = append(g.insts, inst)
}

func (g *E8Gen) addSimpleOp(i *oper, funct uint8) {
	g.loadReg(regT1, i.a)
	g.loadReg(regT2, i.b)
	g.addRinst(regT1, regT2, regT1, funct)
	g.storeReg(regT1, i.dest)
}

// GenFunc generates the instructions for a function.
func (g *E8Gen) GenFunc(f *Func) {
	if len(f.blocks) == 0 {
		return
	}

	// $30 is stack counter
	g.frameSize = g.arrangeStack(f)
	if g.frameSize > 0x7fff {
		panic("frame too large")
	}

	g.genFuncPrologue(f)

	for _, b := range f.blocks {
		g.genBlock(b)
	}

	g.genFuncEpilogue(f)
}

func (g *E8Gen) arrangeStack(f *Func) uint32 {
	g.retAddr = f.stackAlloc(e8AddrSize) // allocate the return address

	offset := uint32(0)
	push := func(v *Var) {
		if v.size != e8AddrSize {
			panic("bug")
		}
		v.addr = offset
		offset += v.size
	}

	// parameters that were pushed on the stack
	// by the caller
	if len(f.rets) > 3 {
		for _, v := range f.rets[3:] {
			push(v)
		}
	}
	if len(f.args) > 3 {
		for _, v := range f.rets[3:] {
			push(v)
		}
	}

	// extra stack spaces for saving local variables
	// that were sent in via registers

	// return address
	push(g.retAddr)

	// return variables
	n := len(f.rets)
	if n > 3 {
		n = 3
	}
	if n > 0 {
		for _, v := range f.rets[:n] {
			push(v)
		}
	}

	// arguments
	n = len(f.args)
	if n > 3 {
		n = 3
	}
	if n > 0 {
		for _, v := range f.args[:n] {
			push(v)
		}
	}

	// local variables
	for _, v := range f.vars {
		push(v)
	}

	return offset
}

func (g *E8Gen) genFuncPrologue(f *Func) {
	// push the registers on statck
	g.storeReg(regRET, g.retAddr)

	// return values, zero them
	for _, v := range f.rets {
		if v.size != 4 {
			panic("todo")
		}
		g.storeReg(0, v)
	}

	n := len(f.args)
	if n > 3 {
		n = 3
	}
	if n > 0 {
		for i, v := range f.args[:n] {
			if v.size != 4 {
				panic("todo")
			}
			g.storeReg(uint8(i+1), v)
		}
	}

	for _, v := range f.vars {
		if v.size != 4 {
			panic("todo")
		}
		g.storeReg(0, v)
	}
}

func (g *E8Gen) genFuncEpilogue(f *Func) {
	for i, v := range f.rets {
		g.loadReg(uint8(i+1), v)
	}

	g.loadReg(regPC, g.retAddr)
}

func (g *E8Gen) genBlock(b *Block) {
	for _, i := range b.insts {
		switch i := i.(type) {
		case *oper:
			g.genOp(i)
		case *call:
			g.genCall(i)
		default:
			panic("bug")
		}
	}
}

func (g *E8Gen) genOp(i *oper) {
	if i.a == nil {
		switch i.op {
		case "", "+":
			g.loadReg(regT1, i.b)
			g.storeReg(regT1, i.dest)
		case "-":
			g.loadReg(regT1, i.b)
			g.addRinst(0, regT1, regT1, e8i.FnSub)
			g.storeReg(regT1, i.dest)
		default:
			panic("bug")
		}
	} else {
		switch i.op {
		case "+":
			g.addSimpleOp(i, e8i.FnAdd)
		case "-":
			g.addSimpleOp(i, e8i.FnSub)
		default:
			panic("bug")
		}
	}
}

func (g *E8Gen) genCall(i *call) {
	if i.f.isSymbol && i.f.modName == "<builtin>" {
		switch i.f.symName {
		case "print":
			if len(i.args) != 1 {
				panic("print not taking one")
			}

			g.loadReg(1, i.args[0])
			g.addSPChange(int16(g.frameSize))

			g.addJalSym(i.f.modName, i.f.symName)

			g.addSPChange(-int16(g.frameSize))
		default:
			panic("bug")
		}
	} else {
		panic("todo")
	}
}

func (g *E8Gen) PrintInst(out io.Writer) {
	for _, inst := range g.insts {
		fmt.Fprint(out, e8i.Inst(inst.i).String())
		if inst.lateBind {
			fmt.Fprintf(out, "  // %s.%s", inst.bindMod, inst.bindSym)
		}
		fmt.Fprintln(out)
	}
}
