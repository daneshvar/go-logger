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
	isStack  EnablerFunc
}

func NewWriter(caller bool, encoder encoder) *Writer {
	w := &Writer{
		encoder: encoder,
		caller:  caller,
	}

	w.Enabler(func(Level, string) bool { return true })
	w.Stack(func(l Level, _ string) bool { return l >= ErrorLevel })

	return w
}

func (w *Writer) Config(scope map[string]string, level *string) error {
	def := InfoLevel

	if level != nil {
		var ok bool
		def, ok = ToLevel(*level)
		if !ok {
			return fmt.Errorf("config value is illegal: %s", *level)
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

func (w *Writer) Stack(fn EnablerFunc) *Writer {
	w.isStack = fn
	return w
}

func (w *Writer) StackByLevel(l Level) *Writer {
	return w.Stack(func(level Level, _ string) bool {
		return level >= l
	})
}

func (w *Writer) StackByScope(m map[string]Level, def Level) *Writer {
	return w.Stack(func(level Level, scope string) bool {
		if l, ok := m[scope]; ok {
			return level >= l
		}
		return level >= def
	})
}
