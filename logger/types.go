package logger

import "io"

const (
	__printReset string = "[-:-:-]"
)

type Logger struct {
	writer io.Writer
	scope  string
}

func NewLogger(writer io.Writer) *Logger {
	logger := Logger{writer: writer}
	return &logger
}

func (logger *Logger) SetScope(scope string) {
	logger.scope = scope
}
