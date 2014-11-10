package ir

import (
	"fmt"
	"io"
)

// SimFunc simulates the function execution.
func SimFunc(f *Func, out io.Writer) {
	if len(f.blocks) == 0 {
		return
	}

	b := f.blocks[0]
	for {
		b = simBlock(b, out)
		if b == nil {
			break
		}
	}

	return
}

// simBlock simulates a basic block and returns the next
// basic block to run.
func simBlock(b *Block, out io.Writer) *Block {
	for _, i := range b.insts {
		switch i := i.(type) {
		case *oper:
			simOp(i, out)
		case *call:
			simCall(i, out)
		default:
			panic("bug")
		}
	}

	return nil // TODO return next block
}

func simOp(i *oper, out io.Writer) {
	if i.a == nil {
		switch i.op {
		case "", "+":
			i.dest.value = i.b.value
		case "-":
			i.dest.value = -i.b.value
		default:
			panic("bug")
		}
	} else {
		switch i.op {
		case "+":
			i.dest.value = i.a.value + i.b.value
		case "-":
			i.dest.value = i.a.value - i.b.value
		default:
			panic("bug")
		}
	}
}

func simCall(i *call, out io.Writer) {
	if i.f.isSymbol && i.f.modName == "<builtin>" {
		switch i.f.symName {
		case "print":
			for _, a := range i.args {
				fmt.Fprintln(out, a.value)
			}
		default:
			panic("bug")
		}
	} else {
		panic("todo")
	}
}
