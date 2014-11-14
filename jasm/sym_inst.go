package jasm

type symInst struct {
	inst
	sym *sym
}

type sym struct {
	labs []string
}
