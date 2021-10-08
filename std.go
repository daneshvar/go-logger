package logger

import (
	"bytes"
	stdlog "log"
	"os"
)

const (
	_loggerWriterDepth  = 2
	_stdLogDefaultDepth = 1
)

func redirectStdLogAt(core *Core, scope string, level Level) func() {
	flags := stdlog.Flags()
	prefix := stdlog.Prefix()
	stdlog.SetFlags(0)
	stdlog.SetPrefix("")
	l := core.GetLogger(scope)
	l.Skip(_stdLogDefaultDepth + _loggerWriterDepth)
	logFunc := levelToFunc(l, level)

	stdlog.SetOutput(&loggerWriter{logFunc})
	return func() {
		stdlog.SetFlags(flags)
		stdlog.SetPrefix(prefix)
		stdlog.SetOutput(os.Stderr)
	}
}

func levelToFunc(logger *Logger, level Level) func(messages ...interface{}) {
	switch level {
	case TraceLevel:
		return logger.Trace
	case DebugLevel:
		return logger.Debug
	case InfoLevel:
		return logger.Info
	case WarnLevel:
		return logger.Warn
	case ErrorLevel:
		return logger.Error
	case FatalLevel:
		return logger.Fatal
	case PanicLevel:
		return logger.Panic
	}

	return logger.Info
}

type loggerWriter struct {
	logFunc func(messages ...interface{})
}

func (l *loggerWriter) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	l.logFunc(string(p))
	return len(p), nil
}
