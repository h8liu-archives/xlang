package jasm

type inst interface{}
type symInst struct {
	inst
	sym *sym
}

type sym struct {
	labs []string
}

type ri struct {
	op         string
	dest       int
	src1, src2 int
}

type mi struct {
	op   string
	dest int
	base int
	off  int32
}

type ii struct {
	op   string
	dest int
	src  int
	im   int32
}

type ji struct {
	op  string
	off int32
}
