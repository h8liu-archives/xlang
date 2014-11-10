package ir

import (
	"bytes"
	"fmt"

	"github.com/h8liu/xlang/prt"
)

// Block is a basic block inside a function.
type Block struct {
	name  string // label name
	index int
	f     *Func
	insts []interface{}
}

type oper struct {
	dest *Var
	a    *Var
	b    *Var
	op   string
}

type call struct {
	dest *Var
	f    *Var
	args []*Var
}

func (b *Block) addInst(i interface{}) {
	b.insts = append(b.insts, i)
}

// AddCall appends a function call op into the basic block.
func (b *Block) AddCall(dest, f *Var, args ...*Var) {
	b.addInst(&call{
		dest: dest,
		f:    f,
		args: args,
	})
}

// AddAssign appends a simple assigning op into the basic block.
func (b *Block) AddAssign(dest, src *Var) {
	b.AddUnaryOp(dest, "", src)
}

// AddUnaryOp appends a unary op into the basic block.
func (b *Block) AddUnaryOp(dest *Var, op string, en *Var) {
	b.addInst(&oper{
		dest: dest,
		a:    nil,
		b:    en,
		op:   op,
	})
}

// AddBinaryOp appends a binary op into the basic block.
func (b *Block) AddBinaryOp(dest, x *Var, op string, y *Var) {
	b.addInst(&oper{
		dest: dest,
		a:    x,
		b:    y,
		op:   op,
	})
}

func (b *Block) instStr(i interface{}) string {
	switch i := i.(type) {
	case *oper:
		if i.a == nil {
			// unary op
			if i.op == "" {
				// just copying
				return fmt.Sprintf("%s = %s", i.dest, i.b)
			}
			return fmt.Sprintf("%s = %s %s", i.dest, i.op, i.b)
		}

		return fmt.Sprintf("%s = %s %s %s", i.dest, i.a, i.op, i.b)

	case *call:
		ret := new(bytes.Buffer)
		if i.dest != nil {
			fmt.Fprintf(ret, "%s = ", i.dest)
		}

		fmt.Fprintf(ret, "call %s (", i.f)

		for i, arg := range i.args {
			if i > 0 {
				fmt.Fprintf(ret, ",")
			}
			fmt.Fprintf(ret, "%s", arg)
		}

		fmt.Fprintf(ret, ")")

		return ret.String()
	default:
		panic("bug")
	}
}

// Print prints out the instructions in this basic block.
func (b *Block) Print(p *prt.Printer) {
	p.Printf("B%d", b.index)
	if b.name != "" {
		p.Printf(" %q", b.name)
	}
	p.Println(":")
	p.ShiftIn()

	for _, i := range b.insts {
		p.Println(b.instStr(i))
	}
	p.Println()
	p.ShiftOut()
}

// Func returns the function that this block belongs to.
func (b *Block) Func() *Func {
	return b.f
}
