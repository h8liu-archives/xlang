// Package regs defines the registers for a E8 virtual machine CPU.
package regs

import (
	"fmt"
	"io"

	"e8vm.net/e8/align"
	"e8vm.net/e8/inst"
)

// Registers container for e8 ALU. Based on the instruction design, it has 32
// 32-bit integer registers arnd 32 64-bit floating point registers.
// For integer register, $0 is bind to 0. $31 is program counter, and its
// last 2 bits are bind to 0. All other ones are general purpose registers.
type Registers struct {
	ints   []uint32
	floats []float64
}

// New creates a new register container.
func New(nint, nfloat int) *Registers {
	ret := new(Registers)

	ret.ints = make([]uint32, nint)
	ret.floats = make([]float64, nfloat)

	return ret
}

// ReadReg reads an integer register.
func (rs *Registers) ReadReg(a uint8) uint32 { return rs.ints[a] }

// ReadFloatReg reads a floating-point register.
func (rs *Registers) ReadFloatReg(a uint8) float64 { return rs.floats[a] }

// WriteReg writes an integer register with value v.
// Writing to $0 will have no effect,
// writing to $31 will be automatically aligned.
func (rs *Registers) WriteReg(a uint8, v uint32) {
	if a == 0 {
		// do nothing
	} else if a == inst.RegPC {
		rs.ints[inst.RegPC] = align.A32(v)
	} else {
		rs.ints[a] = v
	}
}

// WriteFloatReg writes to a floating-point register with value v.
// Writing to $0 will have no effect.
func (rs *Registers) WriteFloatReg(a uint8, v float64) {
	if a == 0 {
		return
	}
	rs.floats[a] = v
}

// IncPC increases the program counter, $31 by 4.
func (rs *Registers) IncPC() uint32 {
	ret := rs.ints[inst.RegPC]
	rs.ints[inst.RegPC] += 4
	return ret
}

// PrintTo prints the register values to an output stream.
// FIXME: currently it only prints the integer registers.
func (rs *Registers) PrintTo(w io.Writer) {
	for i := uint8(0); i < inst.Nreg; i++ {
		fmt.Fprintf(w, "$%02d:%08x", i, rs.ReadReg(i))
		if (i+1)%4 == 0 {
			fmt.Fprintln(w)
		} else {
			fmt.Fprint(w, " ")
		}
	}
}
