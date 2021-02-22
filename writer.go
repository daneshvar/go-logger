package logger

type encoder interface {
	Print(l Level, s string, caller string, stack []string, messages []interface{})
	Prints(l Level, s string, caller string, stack []string, message string)
	Printf(l Level, s string, caller string, stack []string, format string, args []interface{})
	Printv(l Level, s string, caller string, stack []string, message string, keysValues []interface{})
	close()
}

type Writer struct {
	encoder
	enabler EnablerFunc
	stack   EnablerFunc
	caller  bool
}

func newWriter(enabler EnablerFunc, stack EnablerFunc, caller bool, encoder encoder) *Writer {
	return &Writer{
		encoder: encoder,
		enabler: enabler,
		caller:  caller,
		stack:   stack,
	}
}

func (w *Writer) isEnable(level Level, scope string) bool {
	return w.enabler(level, scope)
}

func (w *Writer) isStack(level Level, scope string) bool {
	return w.stack(level, scope)
}
