package jasm

type Image struct {
	code     *cblock
	heapSize int
	data     []*dblock
}

type cblock struct {
	insts []inst
}

type dblock struct {
	start uint32
	dat   []byte
}

type fn struct {
	blocks []*bb          // function basic blocks
	labels map[string]int // function basic block labels

	loc uint32
}

// basic block in function
type bb struct {
	insts  []*symInst
	offset uint32
}

type data struct {
	size  int
	align int
	recs  []rec
}
