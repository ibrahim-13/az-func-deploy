package logger

import "fmt"

func (logger *Logger) BlackRedln(msg string) {
	logger.formattedLog("[black:red]%s[-:-:-]\n", msg)
}

func (logger *Logger) Redln(msg string) {
	logger.formattedLog("[red]%s[-:-:-]\n", msg)
}

func (logger *Logger) Highlightln(msg string) {
	logger.formattedLog("[#2dd1d6:#191bb0]%s[-:-:-]\n", msg)
}

func (logger *Logger) BlackYellowln(msg string) {
	logger.formattedLog("[black:yellow]%s[-:-:-]\n", msg)
}

func (logger *Logger) Greenln(msg string) {
	logger.formattedLog("[green]%s[-:-:-]\n", msg)
}

func (logger *Logger) WhiteGreenln(msg string) {
	logger.formattedLog("[white:green]%s[-:-:-]\n", msg)
}

func (logger *Logger) WhiteBlueln(msg string) {
	logger.formattedLog("[white:blue]%s[-:-:-]\n", msg)
}

func (logger *Logger) formattedLog(format string, msg string) {
	if logger.useColor {
		logger.writer.Write([]byte(fmt.Sprintf(format, msg)))
	} else {
		logger.writer.Write([]byte(msg))
	}
}
