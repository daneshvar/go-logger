// +build !release

package logger

func (l *Logger) Trace(messages ...interface{}) {
	log.print(TraceLevel, l.scope, l.skip, messages)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	log.printf(TraceLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Tracev(message string, keysValues ...interface{}) {
	log.printv(TraceLevel, l.scope, l.skip, message, keysValues)
}
