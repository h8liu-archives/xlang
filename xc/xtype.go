package xc

type xtype struct {
	isVoid bool

	isInt    bool
	unsigned bool
	width    uint32

	isFunc bool
	ret    *xtype
	args   []*xtype
}

var (
	typeVoid = &xtype{isVoid: true}

	typeInt  = &xtype{isInt: true, width: 32, unsigned: false}
	typeUint = &xtype{isInt: true, width: 32, unsigned: true}
	typeChar = &xtype{isInt: true, width: 8, unsigned: false}
	typeByte = &xtype{isInt: true, width: 8, unsigned: true}
)

func newFuncType(ret *xtype, args ...*xtype) *xtype {
	return &xtype{
		isFunc: true,
		ret:    ret,
		args:   args,
	}
}

func (t *xtype) String() string {
	if t.isVoid {
		return "void"
	}

	if t.isInt {
		switch {
		case t.width == 32 && !t.unsigned:
			return "int"
		case t.width == 32 && t.unsigned:
			return "uint"
		case t.width == 8 && !t.unsigned:
			return "char"
		case t.width == 8 && t.unsigned:
			return "byte"
		default:
			panic("bug")
		}
	}

	panic("bug")
}

func (t *xtype) canAssignTo(d *xtype) bool {
	if t.isVoid || d.isVoid {
		return false
	}

	if t.isInt {
		if t == d {
			return true
		}
		if !d.isInt {
			return false
		}
		if t.width != d.width {
			return false
		}
		return t.unsigned == d.unsigned
	}

	panic("bug")
}

func (t *xtype) size() uint32 {
	if t.isVoid {
		return 0
	}

	if t.isInt {
		return t.width / 8
	}

	panic("bug")
}

func (t *xtype) isNum() bool {
	if t.isInt {
		return true
	}
	return false
}

func (t *xtype) numEquals(other *xtype) bool {
	if !t.isInt {
		return false
	}
	if !other.isInt {
		return false
	}

	if t.width != other.width {
		return false
	}
	if t.unsigned != other.unsigned {
		return false
	}

	return true
}
