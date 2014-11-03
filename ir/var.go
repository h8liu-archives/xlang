package ir

import (
	"bytes"
	"fmt"
)

// Var represents a variable (or a general memory area).
// Var does not have a type referring to it; typing and
// type checking is the compiler's issue.
type Var struct {
	Name string // just for debugging

	onHeap bool // if it is allocated on the heap or the stack
	size   uint32
	index  int
	addr   uint32

	isConst bool  // if it is a just a constant
	value   int32 // the constant value

	isVoid bool // if this variable is not a void spaceholder

	// TODO: think more about these optimization fields
	isVolatile bool // if the variable must be alloced in mem
}

func (v *Var) String() string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, "<")

	if v.isVoid {
		fmt.Fprintf(ret, "void")
	} else if v.isConst {
		fmt.Fprintf(ret, "C#%d", v.value)
	} else if v.onHeap {
		fmt.Fprintf(ret, "H#%d", v.index)
	} else {
		fmt.Fprintf(ret, "#%d", v.index)
	}

	if v.Name != "" {
		fmt.Fprintf(ret, " %q", v.Name)
	}

	fmt.Fprintf(ret, ">")
	return ret.String()
}

// IsConst returns true when the variable is a constant.
func (v *Var) IsConst() bool { return v.isConst }

// IsVoid returns true when the variable is the void placeholder.
func (v *Var) IsVoid() bool { return v.isVoid }

// IsAddressable return true when the variable is addressable.
func (v *Var) IsAddressable() bool { return !(v.isConst || v.isVoid) }

// NewConst returns a new constant variable.
func NewConst(v int32) *Var {
	ret := new(Var)
	ret.isConst = true
	ret.value = v
	return ret
}

// VoidVar is the variable that will be used as a placeholder for void
// return values.
var Void = &Var{isVoid: true}
