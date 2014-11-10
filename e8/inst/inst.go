// Package inst defines all the E8 instructions.
package inst

// Inst represents an E8 instruction.
type Inst uint32

// The bit positions and masks
const (
	OpShift    = 26
	RsShift    = 21
	RtShift    = 16
	RdShift    = 11
	ShamtShift = 6

	OpMask    = 0x3f << OpShift
	RsMask    = 0x1f << RsShift
	RtMask    = 0x1f << RtShift
	RdMask    = 0x1f << RdShift
	ShamtMask = 0x1f << ShamtShift
	FunctMask = 0x3f
	ImMask    = 0xffff
)

// The total possible number of functs and ops.
const (
	Nfunct = 64
	Nop    = 64
)

// U32 returns the uint32 representation
func (i Inst) U32() uint32 { return uint32(i) }

// Op returns the op field
func (i Inst) Op() uint8 { return uint8(i >> 26) }

// Rs returns the rs field
func (i Inst) Rs() uint8 { return uint8(i>>21) & 0x1f }

// Rt returns the rt field
func (i Inst) Rt() uint8 { return uint8(i>>16) & 0x1f }

// Rd returns the rd field
func (i Inst) Rd() uint8 { return uint8(i>>11) & 0x1f }

// Sh returns the shamt field
func (i Inst) Sh() uint8 { return uint8(i>>6) & 0x1f }

// Fn returns the funct field
func (i Inst) Fn() uint8 { return uint8(i) & 0x3f }

// Imu returns the immediate (16-bit) field as an unsigned int
func (i Inst) Imu() uint16 { return uint16(i) }

// Ims returns the immediate (16-bit) field as an signed int
func (i Inst) Ims() int16 { return int16(uint16(i)) }

// Off returns the address field
func (i Inst) Off() int32 { return int32(i) << 6 >> 6 }

type instFunc func(c Core, i Inst)

func makeInstList(m map[uint8]instFunc, n uint8) []instFunc {
	ret := make([]instFunc, n)
	for i := range ret {
		ret[i] = opNoop
	}
	for i, inst := range m {
		ret[i] = inst
	}
	return ret
}

// Instruction op codes.
const (
	OpRinst = 0
	OpJ     = 0x02
	OpJal   = 0x03
	OpBeq   = 0x04
	OpBne   = 0x05

	OpAddi = 0x08
	OpSlti = 0x0A
	OpAndi = 0x0C
	OpOri  = 0x0D
	OpLui  = 0x0F

	OpLw  = 0x23
	OpLh  = 0x21
	OpLhu = 0x25
	OpLb  = 0x20
	OpLbu = 0x24
	OpSw  = 0x2B
	OpSh  = 0x29
	OpSb  = 0x28
)

var instList = makeInstList(
	map[uint8]instFunc{
		OpRinst: opRinst,
		OpJ:     opJ,
		OpJal:   opJal,
		OpBeq:   opBeq,
		OpBne:   opBne,

		OpAddi: opAddi,
		OpSlti: opSlti,
		OpAndi: opAndi,
		OpOri:  opOri,
		OpLui:  opLui,

		OpLw:  opLw,
		OpLh:  opLh,
		OpLhu: opLhu,
		OpLb:  opLb,
		OpLbu: opLbu,
		OpSw:  opSw,
		OpSh:  opSh,
		OpSb:  opSb,
	}, Nop,
)

// Instruction funct codes.
const (
	FnAdd = 0x20
	FnSub = 0x22
	FnAnd = 0x24
	FnOr  = 0x25
	FnXor = 0x26
	FnNor = 0x27
	FnSlt = 0x2A

	FnMul  = 0x18
	FnMulu = 0x19
	FnDiv  = 0x1A
	FnDivu = 0x1B
	FnMod  = 0x1C
	FnModu = 0x1D

	FnSll  = 0x00
	FnSrl  = 0x02
	FnSra  = 0x03
	FnSllv = 0x04
	FnSrlv = 0x06
	FnSrav = 0x07
)

var rInstList = makeInstList(
	map[uint8]instFunc{
		FnAdd: opAdd,
		FnSub: opSub,
		FnAnd: opAnd,
		FnOr:  opOr,
		FnXor: opXor,
		FnNor: opNor,
		FnSlt: opSlt,

		FnMul:  opMul,
		FnMulu: opMulu,
		FnDiv:  opDiv,
		FnDivu: opDivu,
		FnMod:  opMod,
		FnModu: opModu,

		FnSll:  opSll,
		FnSrl:  opSrl,
		FnSra:  opSra,
		FnSllv: opSllv,
		FnSrlv: opSrlv,
		FnSrav: opSrav,
	}, Nfunct,
)

// Exec executes an instruction.
func Exec(c Core, i Inst) { instList[i.Op()](c, i) }

func opRinst(c Core, i Inst) { rInstList[i.Fn()](c, i) }

func opJ(c Core, i Inst) {
	pc := c.ReadReg(RegPC)
	c.WriteReg(RegPC, pc+uint32(int32(i<<6)>>4))
}

func opJal(c Core, i Inst) {
	pc := c.ReadReg(RegPC)
	c.WriteReg(RegPC, pc+uint32(int32(i<<6)>>4))
	c.WriteReg(RegRet, pc)
}

func opNoop(c Core, i Inst) {}
