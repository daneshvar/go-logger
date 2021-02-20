package logger

// import (
// 	stdlog "log"
// 	"os"
// )

// func redirectStdLogAt(l *Logger, level Level) (func(), error) {
// 	flags := stdlog.Flags()
// 	prefix := stdlog.Prefix()
// 	stdlog.SetFlags(0)
// 	stdlog.SetPrefix("")
// 	logger := l.WithOptions(AddCallerSkip(_stdLogDefaultDepth + _loggerWriterDepth))
// 	logFunc, err := levelToFunc(logger, level)
// 	if err != nil {
// 		return nil, err
// 	}
// 	stdlog.SetOutput(&loggerWriter{logFunc})
// 	return func() {
// 		stdlog.SetFlags(flags)
// 		stdlog.SetPrefix(prefix)
// 		stdlog.SetOutput(os.Stderr)
// 	}, nil
// }
