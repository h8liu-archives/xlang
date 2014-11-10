// Package align defines functions for aligning memory addresses.
package align

// A16 lower aligns an offset to 2-byte grids
func A16(offset uint32) uint32 { return offset >> 1 << 1 }

// A32 lower aligns an offset to 4-byte grids
func A32(offset uint32) uint32 { return offset >> 2 << 2 }

// A64 lower aligns an offset to 8-byte grids
func A64(offset uint32) uint32 { return offset >> 3 << 3 }
