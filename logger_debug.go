//go:build !release
// +build !release

package logger

func (l *Logger) Trace(messages ...interface{}) {
	l.core.print(TraceLevel, l.scope, l.skip, messages)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.core.printf(TraceLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Tracev(message string, keysValues ...interface{}) {
	l.core.printv(TraceLevel, l.scope, l.skip, message, keysValues)
}
