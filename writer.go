package logger

import "fmt"

type encoder interface {
	Print(l Level, s string, caller string, stack []string, messages []interface{})
	Prints(l Level, s string, caller string, stack []string, message string)
	Printf(l Level, s string, caller string, stack []string, format string, args []interface{})
	Printv(l Level, s string, caller string, stack []string, message string, keysValues []interface{})
	Close()
}

type Writer struct {
	encoder
	caller bool

	isEnable EnablerFunc
	isStack  Level
}

func NewWriter(caller bool, encoder encoder) *Writer {
	w := &Writer{
		encoder: encoder,
		caller:  caller,
		isStack: ErrorLevel,
	}

	w.Enabler(func(Level, string) bool { return true })

	return w
}

func (w *Writer) Config(scope map[string]string, level *string, stack *string) error {
	def := TraceLevel
	stackLevel := ErrorLevel

	if level != nil {
		var ok bool
		def, ok = ToLevel(*level)
		if !ok {
			return fmt.Errorf("config value is illegal: %s", *level)
		}
	}

	if stack != nil {
		var ok bool
		stackLevel, ok = ToLevel(*stack)
		if !ok {
			return fmt.Errorf("config value is illegal: %s", *stack)
		}
	}

	if scope != nil {
		r := make(map[string]Level)
		for key, value := range scope {
			l, ok := ToLevel(value)
			if !ok {
				return fmt.Errorf("config value is illegal: %s:%s", key, value)
			}
			r[key] = l
		}
		w.EnablerByScope(r, def)
	} else {
		w.EnablerByLevel(def)
	}

	w.StackByLevel(stackLevel)

	return nil
}

func (w *Writer) Enabler(fn EnablerFunc) *Writer {
	w.isEnable = fn
	return w
}

func (w *Writer) EnablerByLevel(l Level) *Writer {
	return w.Enabler(func(level Level, _ string) bool {
		return level >= l
	})
}

func (w *Writer) EnablerByScope(m map[string]Level, def Level) *Writer {
	return w.Enabler(func(level Level, scope string) bool {
		if l, ok := m[scope]; ok {
			return level >= l
		}
		return level >= def
	})
}

func (w *Writer) StackByLevel(l Level) *Writer {
	w.isStack = l
	return w
}
