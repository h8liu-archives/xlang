package vm

import (
	"bytes"
	"testing"

	"e8vm.net/e8/inst"
	"e8vm.net/e8/mem"

	// "os"
)

func TestHelloWorld(t *testing.T) {
	c := New()
	// c.Log = os.Stdout
	out := new(bytes.Buffer)
	str := "Hello, world.\n"
	c.Stdout = out

	dpage := mem.NewPage()
	copy(dpage.Bytes(), []byte(str+"\000"))

	ipage := mem.NewPage()
	c.MapPage(mem.PageStart(1), ipage)
	c.MapPage(mem.PageStart(2), dpage)

	a := &mem.Align{ipage}

	offset := uint32(0)
	w := func(i inst.Inst) uint32 {
		ret := offset
		a.WriteU32(offset, i.U32())
		offset += 4
		return ret
	}

	/*
			add $1, $0, $0		; init counter
		loop:
			lbu $2, $1[0x2000] 	; load byte
			beq $2, $0, end 	; +5
		wait:
			lbu $3, $0[0x9]    	; is output ready?
			bne $3, $0, wait 	; -2
			sb $2, $0[0x9]		; output byte
			addi $1, $1, 1		; increase counter
			j loop 				; -7
		end:
			sb $0, [0x8]
	*/

	Rinst := inst.Rinst
	Iinst := inst.Iinst
	Jinst := inst.Jinst

	w(Rinst(0, 0, 1, inst.FnAdd))       // 000
	w(Iinst(inst.OpLbu, 1, 2, 0x2000))  // 004
	w(Iinst(inst.OpBeq, 2, 0, 0x0005))  // 008
	w(Iinst(inst.OpLbu, 0, 3, 0x0009))  // 00c
	w(Iinst(inst.OpBne, 3, 0, 0xfffe))  // 010
	w(Iinst(inst.OpSb, 0, 2, 0x0009))   // 014
	w(Iinst(inst.OpAddi, 1, 1, 0x0001)) // 018
	w(Jinst(inst.OpJ, -7))              // 01c
	w(Iinst(inst.OpSb, 0, 0, 0x0008))   // 020

	c.SetPC(mem.PageStart(1))
	used := c.Run(1000)

	if used > 150 {
		t.Fail()
	}

	if !c.RIP() {
		t.Fail()
	}

	if out.String() != str {
		t.Fail()
	}
}
