package color

import "io"

type Writer interface {
	io.Writer
	io.StringWriter
	Colorable() bool
	SetColorable(bool)
}

type writer struct {
	output  io.Writer
	colored bool
}

func NewWriter(out io.Writer) Writer {
	return &writer{
		output:  out,
		colored: Colorable(out),
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	if !w.colored {
		return w.output.Write(StripAnsi(p))
	}
	return w.output.Write(p)
}

func (w *writer) WriteString(s string) (n int, err error) {
	return w.Write([]byte(s))
}

func (w *writer) Colorable() bool {
	return w.colored
}

func (w *writer) SetColorable(b bool) {
	w.colored = b
}
