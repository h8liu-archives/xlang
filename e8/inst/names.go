package inst

type codeMap struct {
	names map[uint8]string
	codes map[string]uint8
}

func reverseMap(m map[uint8]string) map[string]uint8 {
	ret := make(map[string]uint8)
	for k, v := range m {
		ret[v] = k
	}
	return ret
}

func newCodeMap(m map[uint8]string) *codeMap {
	ret := new(codeMap)
	ret.names = m
	ret.codes = reverseMap(m)
	return ret
}

func (m *codeMap) code(s string) uint8 { return m.codes[s] }
func (m *codeMap) name(c uint8) string { return m.names[c] }

var (
	functMap = newCodeMap(map[uint8]string{
		FnAdd:  "add",
		FnSub:  "sub",
		FnAnd:  "and",
		FnOr:   "or",
		FnXor:  "xor",
		FnNor:  "nor",
		FnSlt:  "slt",
		FnMul:  "mul",
		FnMulu: "mulu",
		FnDiv:  "div",
		FnDivu: "divu",
		FnMod:  "mod",
		FnModu: "modu",
		FnSll:  "sll",
		FnSrl:  "srl",
		FnSra:  "sra",
		FnSllv: "sllv",
		FnSrav: "srav",
	})

	opMap = newCodeMap(map[uint8]string{
		OpBeq:  "beq",
		OpBne:  "bne",
		OpAddi: "addi",
		OpSlti: "slti",
		OpAndi: "andi",
		OpOri:  "ori",
		OpLui:  "lui",
		OpLw:   "lw",
		OpLh:   "lh",
		OpLhu:  "lhu",
		OpLb:   "lb",
		OpLbu:  "lbu",
		OpSw:   "sw",
		OpSh:   "sh",
		OpSb:   "sb",
		OpJ:    "j",
		OpJal:  "jal",
	})
)

// FunctName returns the name of a funct code.
func FunctName(f uint8) string { return functMap.name(f) }

// OpName returns the name of an op code.
func OpName(op uint8) string { return opMap.name(op) }

// FunctCode returns the code of a funct name.
func FunctCode(f string) uint8 { return functMap.code(f) }

// OpCode returns the code of an op name.
func OpCode(op string) uint8 { return opMap.code(op) }
