package ir

// Block is an IR basic block.
type Block struct {
	varMap map[string][]*Var
	vars   []*Var

	insts []*Inst
}

// Var is an variable.
type Var struct {
	Name    string
	Version int
	Index   int
}

// NewBlock creates an IR basic block.
func NewBlock() *Block {
	ret := new(Block)
	ret.varMap = make(map[string][]*Var)
	return ret
}

// Temp declares an anonymous variable.
func (b *Block) Temp() *Var {
	return b.Decl("_")
}

// Decl declares a variable. If the name is "_" it is an
// anonymous variable and can only be later referenced by its
// index.
func (b *Block) Decl(name string) *Var {
	ret := new(Var)
	ret.Name = name

	if name == "" {
		panic("var name cannot be empty")
	}

	if name == "_" {
		vers, found := b.varMap[name]
		if !found {
			vers = make([]*Var, 0, 8)
			b.varMap[name] = vers
		}

		ret.Version = len(vers)
		vers = append(vers, ret)
	}

	ret.Index = len(b.vars)
	b.vars = append(b.vars, ret)

	return ret
}

// FindByName returns the latest version of the variable by name.
// If the variable has not been declared, it returns nil.
func (b *Block) FindByName(name string) *Var {
	vars, found := b.varMap[name]
	if !found {
		return nil
	}

	n := len(vars)
	if n == 0 {
		panic("bug")
	}
	return vars[n-1]
}

// FindByIndex returns the variable by index.
func (b *Block) FindByIndex(i int) *Var {
	return b.vars[i]
}

// AddInst appends an instruction at the end of the block.
// It returns the index of the instruction in this block.
func (b *Block) AddInst(i *Inst) int {
	i.Index = len(b.insts)
	ret := i.Index
	b.insts = append(b.insts, i)
	return ret
}
