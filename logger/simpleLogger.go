package logger

func (logger *Logger) BlackRedln(msg string) {
	logger.writer.Write([]byte("[black:red]" + msg + __printReset + "\n"))
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

func (logger *Logger) WhiteGreenln(msg string) {
	logger.writer.Write([]byte("[white:green]" + msg + __printReset + "\n"))
}

func (logger *Logger) WhiteBlueln(msg string) {
	logger.writer.Write([]byte("[white:blue]" + msg + __printReset + "\n"))
}
