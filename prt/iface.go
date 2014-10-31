package prt

// Iface is the printer interface
type Iface interface {
	Print(a ...interface{}) (int, error)
	Println(a ...interface{}) (int, error)
	Printf(format string, a ...interface{}) (int, error)
	ShiftIn()
	ShiftOut(a ...interface{})
}
