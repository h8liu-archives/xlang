package inst

// Core defines the interface that a CPU requires for
// executing instructions.
type Core interface {
	// Integer register operations
	WriteReg(a uint8, v uint32)
	ReadReg(a uint8) uint32

	// Floating point register operations
	WriteFloatReg(a uint8, v float64)
	ReadFloatReg(a uint8) float64

	// Memory operations
	WriteU8(addr uint32, v uint8)
	WriteU16(addr uint32, v uint16)
	WriteU32(addr uint32, v uint32)
	WriteF64(addr uint32, v float64)

	ReadU8(addr uint32) uint8
	ReadU16(addr uint32) uint16
	ReadU32(addr uint32) uint32
	ReadF64(addr uint32) float64
}

const (
	// Nreg is the number of registers
	Nreg = 32

	// RegPC is the index of the program counter register
	RegPC = Nreg - 1

	// RegRet is the index of the function return register
	RegRet = Nreg - 2
)
