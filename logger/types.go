package logger

import "io"

const (
	__printReset string = "[-:-:-]"
)

type Logger struct {
	writer   io.Writer
	scope    string
	useColor bool
}

func NewLogger(writer io.Writer) *Logger {
	logger := Logger{writer: writer}
	return &logger
}

func (logger *Logger) SetScope(scope string) {
	logger.scope = scope
}

func (logger *Logger) SetColor(val bool) {
	logger.useColor = val
}
