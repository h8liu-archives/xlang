package xc

import (
	"io"

	"github.com/h8liu/xlang/ir"
	"github.com/h8liu/xlang/prt"
)

// Object defines the compiling result of a source file.
type Object struct {
	header *Header
	f      *ir.Func
	b      *ir.Block

	e8 *ir.E8Gen
}

// Header returns the header of an object file.
func (obj *Object) Header() *Header {
	return obj.header
}

// PubHeader returns the public header of an object file.
func (obj *Object) PubHeader() *Header {
	return nil
}

// PrintIR prints the IR of the default block.
func (obj *Object) PrintIR(out io.Writer) {
	p := prt.New(out)
	obj.f.Print(p)
}

// Sim simulates the object and write the output result to out.
func (obj *Object) Sim(out io.Writer) {
	ir.SimFunc(obj.f, out)
}

// GenE8 generates e8 instructions.
func (obj *Object) GenE8() *ir.E8Gen {
	ret := ir.NewE8Gen()
	ret.GenFunc(obj.f)
	obj.e8 = ret

	return ret
}
