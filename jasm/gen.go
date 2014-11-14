package jasm

const header = `
function jasm(stdlib, foreign, heap) {
    "use asm";

    var pc = 0;     // pseudo program counter
    var sp = 0;     // stack pointer
    var ret = 0;    // return address, for jal
    var r0 = 0, r1 = 0, r2 = 0, r3 = 0; // general purpose 32-bit registers
    var f0 = 0.0, f1 = 0.0, f2 = 0.0, f3 = 0.0; // temp floating point registers
    var err = 0;

    var memI32 = new stdlib.Int32Array(heap);
    var memU32 = new stdlib.Uint32Array(heap);
    var memI8 = new stdlib.Int8Array(heap);
    var memU8 = new stdlib.Uint8Array(heap);
    var memF64 = new stdlib.Float64Array(heap);

    function setpc(newpc) { newpc = newpc|0; pc = newpc|0; }
    function setsp(newsp) { newsp = newsp|0; sp = newsp|0; }
    function seterr(newerr) { newerr = newerr|0; err = newerr|0; }
    function setret(newret) { newret = newret|0; ret = newret|0; }
    function getpc() { return pc|0; }
    function getsp() { return sp|0; }
    function getret() { return ret|0; }
    function geterr() { return err|0; }
    function getr1() { return r1|0; }
    function getr2() { return r2|0; }
    function getr3() { return r3|0; }
    function getf0() { return +f0; }
    function getf1() { return +f1; }
    function getf2() { return +f2; }
    function getf3() { return +f3; }
    function clearRegs() {
        pc = 0|0;
        sp = 0|0;
        ret = 0|0;
        err = 0|0;
        r0 = 0|0;
        r1 = 0|0;
        r2 = 0|0;
        r3 = 0|0;
        f0 = 0.0;
        f1 = 0.0;
        f2 = 0.0;
        f3 = 0.0;
    }

    function step() {
		var pc = 0;
		pc_ = pc|0;
        pc = (pc + 4) | 0;

        switch (pc_|0) {
`

const footer = `
        default: err = 1|0;
        }
    }

    function run(ncycle) {
        ncycle = ncycle|0;

        while (ncycle|0 > 0) {
            step();
            r0 = 0|0;
            ncycle = ((ncycle|0) + -1)|0;

            if ((err|0) != (0|0)) {
                break;
            }
        }
    }

    return {
        setpc: setpc,
        setsp: setsp,
        seterr: seterr,
        setret: setret,

        getpc: getpc,
        getsp: getsp,
        geterr: geterr,
        getret: getret,

        getr1: getr1,
        getr2: getr2,
        getr3: getr3,

        getf0: getf0,
        getf1: getf1,
        getf2: getf2,
        getf3: getf3,
        
        clearRegs: clearRegs,

        run: run,
    };
}
`
