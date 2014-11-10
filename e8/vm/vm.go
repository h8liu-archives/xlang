/*
Package vm maintains the basic CPU unit needed for instruction execution.
It defines the interface for a VM core, and also implements the registers in it.
*/
package vm

import (
	"fmt"
	"io"
	"os"

	"e8vm.net/e8/inst"
	"e8vm.net/e8/mem"
)

// VM has a core and also has the system page
type VM struct {
	Stdout io.Writer // standard output
	Log    io.Writer // debug logging

	core *Core
	sys  *SysPage
}

// New creates an empty virtual machine that has
// only the system page mounted.
func New() *VM {
	ret := new(VM)
	ret.Stdout = os.Stdout

	ret.core = NewCore()
	ret.sys = NewSysPage()

	ret.MapPage(0, ret.sys)

	return ret
}

// Step executes one instruction.
func (vm *VM) Step() {
	vm.sys.Reset()

	pc := vm.core.IncPC()
	u32 := vm.core.ReadU32(pc)
	in := inst.Inst(u32)
	if vm.Log != nil {
		fmt.Fprintf(vm.Log, "%08x: %08x   %v", pc, u32, in)
		if in.Op() != inst.OpJ {
			rs := in.Rs()
			rt := in.Rt()
			rsv := vm.core.ReadReg(rs)
			rtv := vm.core.ReadReg(rt)
			fmt.Fprintf(vm.Log, "  ; $%d=%d(%08x) $%d=%d(%08x)",
				rs, rsv, rsv, rt, rtv, rtv)
		}
		fmt.Fprintf(vm.Log, "\n")
		// vm.Registers.PrintTo(vm.Log)
	}
	inst.Exec(vm.core, in)

	vm.sys.FlushStdout(vm.Stdout)
}

// Run executes at most n instructions. Returns the number of instructions
// actually executed. A core may return early when the core halts.
func (vm *VM) Run(n int) int {
	i := 0
	for i < n {
		vm.Step()
		i++

		if vm.sys.Halted() {
			break
		}
	}

	return i
}

// SetPC sets the program counter. Note the last 2 bits are bind to 0, so the
// program counter will be automatically aligned.
func (vm *VM) SetPC(pc uint32) {
	vm.core.WriteReg(inst.RegPC, pc)
}

// Halted returns if the core halted.
// Currently, a core can halt gracefully by writing a byte to address 0x4.
// Or it will halt because of writing to address 0x0 to 0x7, which will
// cause the core halts because of an address error.
func (vm *VM) Halted() bool { return vm.sys.Halted() }

// AddrError returns if the core halted because of an address error.
// Address error currently only occurs when visiting the word at address 0.
func (vm *VM) AddrError() bool { return vm.sys.AddrError }

// HaltValue returns the value when the core halts. This the byte written to
// address 0x4.
func (vm *VM) HaltValue() uint8 { return vm.sys.HaltValue }

// RIP returns if the core rests in peace, which means it halt with a halt
// value of 0 (writing a byte 0 to 0x4).
func (vm *VM) RIP() bool {
	return vm.Halted() && vm.HaltValue() == 0 && !vm.AddrError()
}

// CheckPage checks if a page is valid
func (vm *VM) CheckPage(addr uint32) bool {
	return vm.core.Check(addr)
}

// MapPage maps a page at particular address
func (vm *VM) MapPage(addr uint32, p mem.Page) {
	vm.core.Map(addr, p)
}

// DumpRegs prints registers to out
func (vm *VM) DumpRegs(out io.Writer) {
	vm.core.PrintTo(out)
}
