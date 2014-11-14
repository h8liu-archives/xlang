package jasm

type inst interface{}

type ri struct {
	op         string
	dest       int
	src1, src2 int
}

type mi struct {
	op   string
	reg  int
	base int
	off  int32
}

type ii struct {
	op   string
	dest int
	src  int
	im   int32
}

type bi struct {
	op         string
	src1, src2 int
	off        int32
}

type ji struct {
	op  string
	off int32
}
