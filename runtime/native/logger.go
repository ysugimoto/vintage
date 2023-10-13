package native

import (
	"io"
	"log"
	"os"
)

var loggerStack = map[string]io.Writer{}

type logger struct {
	l *log.Logger
}

func (l *logger) Write(p []byte) (int, error) {
	l.l.Println(string(p))
	return len(p), nil
}

func LoggerInitiator(name string) (io.Writer, error) {
	if v, ok := loggerStack[name]; ok {
		return v, nil
	}
	w := &logger{
		l: log.New(os.Stderr, "["+name+"] ", log.LstdFlags),
	}
	loggerStack[name] = w
	return w, nil
}
