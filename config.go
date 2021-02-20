package logger

var log core = core{}

func init() {
	Config(ConsoleWriter(true, func(l Level, s string) bool { return l >= ErrorLevel }, func(l Level, s string) bool { return true }))
}

// Config logger enabler funcs is immutable response
func Config(writers ...*Writer) {
	log.Config(writers)
}

func Close() {
	log.Close()
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
func Sync() {}

// RedirectStdLog std log to this to Info Level
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLog() func() {
	return func() {}
}

// RedirectStdLogAt std log to this at log level
// It returns a function to restore the original prefix and flags and reset the
// standard library's output to os.Stderr.
func RedirectStdLogAt(level Level) func() {
	return func() {}
}
