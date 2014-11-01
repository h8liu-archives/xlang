package ir

// Vars indexes a set of IR variables
type Vars struct {
	varMap map[string][]*Var
	vars   []*Var
}

// NewVars creates a new variable set.
func NewVars() *Vars {
	ret := new(Vars)
	ret.varMap = make(map[string][]*Var)
	return ret
}

// Temp declares an anonymous variable.
func (vs *Vars) Temp() *Var {
	return vs.Decl("_")
}

// Decl declares a variable. If the name is "_" it is an
// anonymous variable and can only be later referenced by its
// index.
func (vs *Vars) Decl(name string) *Var {
	ret := new(Var)
	ret.Name = name

	if name == "" {
		panic("var name cannot be empty")
	}

	if name == "_" {
		vers, found := vs.varMap[name]
		if !found {
			vers = make([]*Var, 0, 8)
			vs.varMap[name] = vers
		}

		ret.Version = len(vers)
		vers = append(vers, ret)
	}

	ret.Index = len(vs.vars)
	vs.vars = append(vs.vars, ret)

	return ret
}

// FindByName returns the latest version of the variable by name.
// If the variable has not been declared, it returns nil.
func (vs *Vars) FindByName(name string) *Var {
	vars, found := vs.varMap[name]
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
func (vs *Vars) FindByIndex(i int) *Var {
	return vs.vars[i]
}
