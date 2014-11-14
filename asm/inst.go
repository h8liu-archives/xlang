package asm

type sym struct {
	labs []string
}

type rinst struct {
	funct uint32
	rs uint32
	rt uint32
	rd uint32
}

type iinst struct {
	op uint32
	rs uint32
	rt uint32
	im uint32
	imSym *sym
}

type jinst struct {
	op uint32
	off uint32
	sym *sym
}

