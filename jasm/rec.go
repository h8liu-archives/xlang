package jasm

type rec interface{}

type rx struct{ dat []byte }
type rc struct{ dat []byte }
type rstr struct{ dat string }
type ri32 struct{ dat []int32 }
type ru32 struct{ dat []uint32 }
