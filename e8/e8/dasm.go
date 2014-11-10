package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/h8liu/xlang/e8/img"
	"github.com/h8liu/xlang/e8/inst"
	"github.com/h8liu/xlang/e8/mem"
)

func isCode(s uint32) bool {
	return s >= mem.SegCode && s-mem.SegCode < mem.SegSize
}

func mainDasm(args []string) {
	for _, f := range args {
		fin, e := os.Open(f)
		if e != nil {
			printError(e)
			continue
		}

		for {
			header, bytes, e := img.Read(fin)
			if e == io.EOF {
				break
			}
			if e != nil {
				fmt.Fprintf(os.Stderr, "error: invalid image: %s\n", e)
				break
			}

			start := header.Start()

			if isCode(start) {
				dumpCodeSeg(start, bytes)
			} else {
				dumpDataSeg(start, bytes)
			}
		}

		fin.Close()
	}
}

func makeInst(buf []byte) inst.Inst {
	ret := uint32(buf[0])
	ret |= uint32(buf[1]) << 8
	ret |= uint32(buf[2]) << 16
	ret |= uint32(buf[3]) << 24

	return inst.Inst(ret)
}

func dumpCodeSeg(start uint32, b []byte) {
	n := len(b)
	fmt.Printf("[code] // %08x - %08x, %d bytes\n",
		start, start+uint32(n), n,
	)

	for i := 0; i < n; i += 4 {
		fmt.Printf("%04x:%04x :", uint16(i>>16), uint16(i))

		j := i + 4
		if j >= n {
			j = n
		}

		b := b[i:j]
		for i := 3; i >= 0; i-- {
			if i >= len(b) {
				fmt.Printf("   ")
			} else {
				fmt.Printf(" %02x", b[i])
			}
		}

		if len(b) == 4 {
			line := makeInst(b)
			fmt.Printf("     %s", line.String())
		}

		fmt.Println()
	}
	fmt.Println()
}

func dumpDataSeg(start uint32, b []byte) {
	n := len(b)
	fmt.Printf("[data] // %08x - %08x, %d bytes\n",
		start, start+uint32(n), n,
	)

	dumper := hex.Dumper(os.Stdout)
	dumper.Write(b)
	dumper.Close()
	fmt.Println()
}
