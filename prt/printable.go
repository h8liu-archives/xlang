package prt

// Printable type can be printed via a printer interface
type Printable interface {
	PrintTo(p Iface)
}
