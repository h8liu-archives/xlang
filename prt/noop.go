package prt

type noopWriter struct{}

var noop = new(noopWriter)

func (w *noopWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}
