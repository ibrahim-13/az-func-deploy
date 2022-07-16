package logger

func (logger *Logger) ScopedRedln(msg string) {
	logger.Redln(formatScopedMsg(logger.scope, msg))
}

func (logger *Logger) ScopedHighlightln(msg string) {
	logger.Highlightln(formatScopedMsg(logger.scope, msg))
}

func (logger *Logger) ScopedBlackYellowln(msg string) {
	logger.BlackYellowln(formatScopedMsg(logger.scope, msg))
}

func (logger *Logger) ScopedGreenln(msg string) {
	logger.Greenln(formatScopedMsg(logger.scope, msg))
}

func formatScopedMsg(scope string, msg string) string {
	return "[" + scope + "] " + msg
}
