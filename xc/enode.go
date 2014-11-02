package xc

type enode struct {
	name string // this is just for debugging
	t    *xtype

	isConst bool
	value   int32

	onHeap bool
	addr   int32
}

func (n *enode) typ() *xtype {
	return n.t
}

func (n *enode) addressable() bool {
	if n.isConst {
		return false
	}

	return true
}
