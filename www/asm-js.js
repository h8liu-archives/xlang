function asmjs(stdlib, foreign, heap) {
    "use asm";

    var pc = 0;     // pseudo program counter
    var sp = 0;     // stack pointer
    var ret = 0;
    var t0 = 0, t1 = 0, t2 = 0, t3 = 0; // temp 32-bit registers
    var f0 = 0.0, f1 = 0.0, f2 = 0.0, f3 = 0.0; // temp floating point registers

    var memI32 = new stdlib.Int32Array(heap);
    var memU32 = new stdlib.Uint32Array(heap);
    var memI8 = new stdlib.Int8Array(heap);
    var memU8 = new stdlib.Uint8Array(heap);
    var memF64 = new stdlib.Float64Array(heap);

    function setpc(newpc) { pc = newpc|0; }
    function setsp(newsp) { sp = newsp|0; }
    function getpc() { return pc|0; }
    function getsp() { return sp|0; }
    function getret() { return ret|0; }
    function gett1() { return t1|0; }
    function gett2() { return t2|0; }
    function gett3() { return t3|0; }
    function getf0() { return f0; }
    function getf1() { return f1; }
    function getf2() { return f2; }
    function getf3() { return f3; }
    function clearRegs() {
        pc = 0|0;
        sp = 0|0;
        ret = 0|0;
        t0 = 0|0;
        t1 = 0|0;
        t2 = 0|0;
        t3 = 0|0;
        f0 = 0.0;
        f1 = 0.0;
        f2 = 0.0;
        f3 = 0.0;
    }

    function step() {
        var nextpc = 0;
        nextpc = (pc + 4) | 0;
        switch (pc|0) {
        case 0: memU32[t0 >> 2] = t1|0; break; // memory access
        case 1: if ((t1|0) == (t2|0)) { nextpc = 3|0; } break; // jmp
        case 2: ret = nextpc; nextpc = ((nextpc|0) + 232) | 0; // call
        case 3: nextpc = ret|0; // return
        case 4: memU32[(sp + 4)>>2] = t1|0; break;
        case 5: memF64[(sp + 8)>>3] = f0; break;
        default: 
        }

        return nextpc|0;
    }

    function run(ncycle) {
        ncycle = ncycle|0;

        while (ncycle|0 > 0) {
            pc = step();
            t0 = 0|0;
            ncycle = ((ncycle|0) + -1)|0;
        }
    }

    return {};
}
