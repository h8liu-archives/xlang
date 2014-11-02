package xc

import (
	"bytes"
	"fmt"
)

type enode struct {
	name string // this is just for debugging
	t    *xtype

	isVoid bool

	isConst bool
	value   int32

	onHeap bool
	addr   int32
}

func (n *enode) typ() *xtype {
	return n.t
}

func (n *enode) addressable() bool {
	if n.isConst || n.isVoid {
		return false
	}

	return true
}

func (n *enode) String() string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, "<")

	if n.isVoid {
		fmt.Fprintf(ret, "void")
	} else if n.isConst {
		fmt.Fprintf(ret, "C#%d", n.value)
	} else if n.onHeap {
		fmt.Fprintf(ret, "H#%d", n.addr)
	} else {
		// on stack
		fmt.Fprintf(ret, "S#%d", n.addr)
	}

	if n.name != "" {
		fmt.Fprintf(ret, " [%s]", n.name)
	}

	fmt.Fprintf(ret, ">")
	return ret.String()
}
