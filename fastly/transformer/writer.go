package transformer

import (
	"bytes"
	"fmt"
)

type writer struct {
	buf bytes.Buffer
}

func newWriter() *writer {
	return &writer{}
}

func (w *writer) write(format string, args ...any) {
	w.buf.WriteString(fmt.Sprintf(format, args...))
}

func (w *writer) writeln(format string, args ...any) {
	w.buf.WriteString(fmt.Sprintf(format+"\n", args...))
}

func (w *writer) writeTo(to bytes.Buffer) {
	to.WriteString(w.buf.String())
}

func (w *writer) String() string {
	return w.buf.String()
}
