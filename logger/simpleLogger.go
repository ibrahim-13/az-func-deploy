package logger

func (logger *Logger) Redln(msg string) {
	logger.writer.Write([]byte("[red]" + msg + __printReset + "\n"))
}

func (logger *Logger) Highlightln(msg string) {
	logger.writer.Write([]byte("[#2dd1d6:#191bb0]" + msg + __printReset + "\n"))
}

func (logger *Logger) BlackYellowln(msg string) {
	logger.writer.Write([]byte("[black:yellow]" + msg + __printReset + "\n"))
}

func (logger *Logger) Greenln(msg string) {
	logger.writer.Write([]byte("[green]" + msg + __printReset + "\n"))
}
