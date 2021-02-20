// +build release

package logger

func (l *Logger) Trace(messages ...interface{}) {}

func (l *Logger) Tracef(format string, args ...interface{}) {}

func (l *Logger) Tracev(message string, keysValues ...interface{}) {}
