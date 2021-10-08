package logger

import (
	"fmt"
)

const (
	defaultSkip = 3
)

type Logger struct {
	core   *Core
	scope  string
	caller bool
	skip   int
}

func (l *Logger) GetLogger(scope string) *Logger {
	return l.core.GetLogger(scope)
}

func (l *Logger) Skip(skip int) *Logger {
	l.skip = skip + defaultSkip
	return l
}

func (l *Logger) Print(messages ...interface{}) {
	l.core.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Println(messages ...interface{}) {
	l.core.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.core.printf(InfoLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Debug(messages ...interface{}) {
	l.core.print(DebugLevel, l.scope, l.skip, messages)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.core.printf(DebugLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Debugv(message string, keysValues ...interface{}) {
	l.core.printv(DebugLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Info(messages ...interface{}) {
	l.core.print(InfoLevel, l.scope, l.skip, messages)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.core.printf(InfoLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Infov(message string, keysValues ...interface{}) {
	l.core.printv(InfoLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Warn(messages ...interface{}) {
	l.core.print(WarnLevel, l.scope, l.skip, messages)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.core.printf(WarnLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Warnv(message string, keysValues ...interface{}) {
	l.core.printv(WarnLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Error(messages ...interface{}) {
	l.core.print(ErrorLevel, l.scope, l.skip, messages)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.core.printf(ErrorLevel, l.scope, l.skip, format, args)
}

func (l *Logger) Errorv(message string, keysValues ...interface{}) {
	l.core.printv(ErrorLevel, l.scope, l.skip, message, keysValues)
}

func (l *Logger) Fatal(messages ...interface{}) {
	l.core.print(FatalLevel, l.scope, l.skip, messages)
	l.core.exit()
}

func (l *Logger) Fatalln(messages ...interface{}) {
	l.core.print(FatalLevel, l.scope, l.skip, messages)
	l.core.exit()
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.core.printf(FatalLevel, l.scope, l.skip, format, args)
	l.core.exit()
}

func (l *Logger) Fatalv(message string, keysValues ...interface{}) {
	l.core.printv(FatalLevel, l.scope, l.skip, message, keysValues)
	l.core.exit()
}

func (l *Logger) Panic(messages ...interface{}) {
	l.core.print(PanicLevel, l.scope, l.skip, messages)
	panic(messages)
}

func (l *Logger) Panicln(messages ...interface{}) {
	l.core.print(PanicLevel, l.scope, l.skip, messages)
	panic(messages)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.core.printf(PanicLevel, l.scope, l.skip, format, args)
	panic(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicv(message string, keysValues ...interface{}) {
	l.core.printv(PanicLevel, l.scope, l.skip, message, keysValues)
	panic(message)
}
