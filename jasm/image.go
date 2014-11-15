package jasm

type Image struct {
	code *cblock

	heapSize int
	data     []*dblock
}

type cblock struct {
	insts []inst
}

type dblock struct {
	off uint32
	dat []byte
}
