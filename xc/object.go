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

func (obj *Object) Sim() {
	obj.f.Sim()
}
