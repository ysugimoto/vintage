package fastly

import (
	"io"

	"github.com/fastly/compute-sdk-go/rtlog"
)

var loggerStack = map[string]io.Writer{}

func LoggerInitiator(name string) (io.Writer, error) {
	if v, ok := loggerStack[name]; ok {
		return v, nil
	}
	endpoint := rtlog.Open(name)
	loggerStack[name] = endpoint
	return endpoint, nil
}
