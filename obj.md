obj/asm file format

- use a human readable format first, linear parsing won't take that much time

```
package "builtin"

import {
 	p "package"
}

// you can just do it with some pattern matching
func printInt {
	sw $1, 16($0)     // print the integer

.tag:
	addi $31, $0, $30

.nonzero:

}

func printInt {
    x 32 24 32 32 //
    x 32 24 32 32 //
    x 32 24 32 32 // binary representation
    j package.Symbol
    jal package.Symbol
    lui $1, package.Symbol
    addi $1, $1, package.Symbol
}

var something 33/*size*/ / 4 /*alignment*/ {
    x 32 23 42 32 43 32 32 32
    c 'a' 'b' 'c'
    i8
    u8
    i16
    u16
    i32
    u32
    str "reader"
    u8 0
}

const something = 32 // must be i32
```
