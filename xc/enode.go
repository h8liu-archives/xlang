package xc

type enode struct {
}

func (n *enode) Type() *xtype {
	return new(xtype)
}

func (n *enode) addressable() bool {
	return false
}

type xtype struct {
	isBasic bool
	name    string
}

func newBasicType(name string) *xtype {
	ret := new(xtype)
	ret.isBasic = true
	ret.name = name
	return ret
}

func (t *xtype) String() string {
	return "int"
}

func (t *xtype) canAssignTo(d *xtype) bool {
	return true
}
