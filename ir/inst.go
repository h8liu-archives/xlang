package ir

// Inst is an IR instruction
type Inst struct {
	A     *Obj
	Op    string
	B     *Obj
	Extra []*Obj
	Index int
}

// Obj could be a variable or a int32 constant
type Obj struct {
	*Var
	Const int32
}
