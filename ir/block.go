package ir

// Block is an IR basic block.
type Block struct {
	vars *Vars

	insts []*Inst
}

// Var is an variable.
type Var struct {
	Name    string
	Version int
	Index   int
	Birth   *Inst
}

// NewBlock creates an IR basic block.
func NewBlock(vars *Vars) *Block {
	ret := new(Block)
	ret.vars = vars
	return ret
}

// AddInst appends an instruction at the end of the block.
// It returns the index of the instruction in this block.
func (b *Block) AddInst(i *Inst) int {
	i.Index = len(b.insts)
	ret := i.Index
	b.insts = append(b.insts, i)
	return ret
}
