package logger

import (
	"fmt"
	"strconv"
)

type EnablerFunc func(level Level, scope string) bool

type core struct {
	writers []*Writer

	print  func(l Level, s string, skip int, messages []interface{})
	printf func(l Level, s string, skip int, format string, args []interface{})
	printv func(l Level, s string, skip int, message string, keysValues []interface{})
}

func (c *core) Close() {
	for _, w := range c.writers {
		w.close()
	}
}

func (c *core) Config(w []*Writer) {
	switch len(w) {
	case 0:
		c.print = func(l Level, s string, skip int, messages []interface{}) {}
		c.printf = func(l Level, s string, skip int, format string, args []interface{}) {}
		c.printv = func(l Level, s string, skip int, message string, keysValues []interface{}) {}
	case 1:
		c.print = c.print1
		c.printf = c.printf1
		c.printv = c.printv1
	// case 2:
	// 	c.Print = c.print2
	// 	c.Printf = c.printf2
	// 	c.Printv = c.printv2
	default:
		c.print = c.printAll
		c.printf = c.printfAll
		c.printv = c.printvAll
	}

	// close old writer
	c.Close()

	c.writers = w
}

func (c *core) print1(l Level, s string, skip int, messages []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stack []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stack = c.getStack(skip)
		}
		w.Print(l, s, caller, stack, messages)
	}
}

func (c *core) printf1(l Level, s string, skip int, format string, args []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stack []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stack = c.getStack(skip)
		}

		w.Printf(l, s, caller, stack, format, args)
	}
}

func (c *core) printv1(l Level, s string, skip int, message string, keysValues []interface{}) {
	w := c.writers[0]
	if w.isEnable(l, s) {
		caller := ""
		var stack []string
		if w.caller {
			caller = c.getCaller(skip)
		}
		if w.isStack(l, s) {
			stack = c.getStack(skip)
		}
		w.Printv(l, s, caller, stack, message, keysValues)
	}
}

func (c *core) printAll(l Level, s string, skip int, messages []interface{}) {
	caller := ""
	var stack []string

	callerS := ""
	var stackS []string

	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stack) == 0 {
					stack = c.getStack(skip)
				}
				stackS = stack
			}

			w.Print(l, s, callerS, stackS, messages)
		}
	}
}

func (c *core) printfAll(l Level, s string, skip int, format string, args []interface{}) {
	caller := ""
	var stack []string

	callerS := ""
	var stackS []string
	message := fmt.Sprintf(format, args...)
	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stack) == 0 {
					stack = c.getStack(skip)
				}
				stackS = stack
			}
			w.Prints(l, s, callerS, stackS, message)
		}
	}
}

func (c *core) printvAll(l Level, s string, skip int, message string, keysValues []interface{}) {
	caller := ""
	var stack []string

	callerS := ""
	var stackS []string
	for _, w := range c.writers {
		if w.isEnable(l, s) {
			if w.caller {
				if caller == "" {
					caller = c.getCaller(skip)
				}
				callerS = caller
			}

			if w.isStack(l, s) {
				if len(stack) == 0 {
					stack = c.getStack(skip)
				}
				stackS = stack
			}
			w.Printv(l, s, callerS, stackS, message, keysValues)
		}
	}
}

func (c *core) getCaller(skip int) string {
	frame, defined := getCallerFrame(skip) // log.callerSkip + callerSkipOffset
	if !defined {
		return ""
	}

	return getFolderFile(frame.File) + ":" + strconv.Itoa(frame.Line)
}

func (c *core) getStack(skip int) []string {
	return getStack(skip)
}
