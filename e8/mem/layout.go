package mem

// SegSize is the size of a memory segment.
const SegSize = (1 << 32) / 4

// The entire memory is separated into 4 segments
// 		SegIO: 		the segment for memory mapped IO pages
//   	SegCode:	the segment for code instructions
//		SegHeap: 	the segment for the heap
//      SegStack:   the segment for the stack(s)
const (
	SegIO uint32 = SegSize * iota
	SegCode
	SegHeap
	SegStack
)
