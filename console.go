package logger

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	consoleLevelText  = []string{"  TRACE  ", "  DEBUG  ", "  INFO   ", "  WARN   ", "  ERROR  ", "  FATAL  ", "  PANIC  "}
	consoleLevelColor = []string{"96", "95", "92", "93", "91", "31", "31"}
)

const (
	defScopeAlign  = 12
	defCallerAlign = 30
)

// const messageAlign = 70

// Console write log to console stderr
type Console struct {
	pool        sync.Pool
	enableColor bool
	scopeAlign  int
	callerAlign int
}

func ConsoleWriter(caller bool, stack EnablerFunc, enabler EnablerFunc) *Writer {
	return newWriter(enabler, stack, caller, &Console{
		pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 150)) // buffer init with 150 size
			b.Reset()
			return b
		}},
		enableColor: true,
		scopeAlign:  defScopeAlign,
		callerAlign: defCallerAlign,
	})
}

// ConsoleWriterWithOptions use default value of scopeAlign & callerAlign with set they with -1
func ConsoleWriterWithOptions(caller bool, stack EnablerFunc, enabler EnablerFunc, scopeAlign int, callerAlign int) *Writer {
	if scopeAlign < 0 {
		scopeAlign = defScopeAlign
	}

	if callerAlign < 0 {
		callerAlign = defCallerAlign
	}

	return newWriter(enabler, stack, caller, &Console{
		pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 150)) // buffer init with 150 size
			b.Reset()
			return b
		}},
		enableColor: true,
		scopeAlign:  scopeAlign,
		callerAlign: callerAlign,
	})
}

func (c *Console) close() {}

func (c *Console) writeMessage(b *bytes.Buffer, l Level, scope string, caller string, m string) (n int) {
	b.WriteString(time.Now().Format("2006-01-02 15:04:05"))

	if c.enableColor {
		c.setColor(b, consoleLevelColor[l])
	}
	b.WriteString(consoleLevelText[l])

	c.writeScope(b, scope)
	c.writeCaller(b, caller)

	if c.enableColor {
		c.resetColor(b)
		b.WriteString(m)
	} else {
		b.WriteByte('"')
		b.WriteString(m)
		b.WriteByte('"')
	}

	n += len(m)
	return
}

func (c *Console) writeScope(b *bytes.Buffer, scope string) (n int) {
	if c.scopeAlign <= 0 {
		if scope != "" {
			b.WriteString("[" + scope + "]  ")
			n = len(scope) + 4
		}
	} else {
		if scope != "" {
			b.WriteString("[" + scope + "]")
			n = c.writeAlign(c.scopeAlign, len(scope)+2, b)
		} else {
			n = c.writeAlign(c.scopeAlign, 0, b)
		}
	}

	return
}

func (c *Console) writeCaller(b *bytes.Buffer, caller string) (n int) {
	if c.callerAlign <= 0 {
		if caller != "" {
			b.WriteString(caller + "  ")
			n = len(caller) + 2
		}
	} else {
		if caller != "" {
			b.WriteString(caller)
			n = c.writeAlign(c.callerAlign, len(caller), b)
		}
	}

	return
}

func (c *Console) writeAlign(align int, len int, b *bytes.Buffer) int {
	if len < align {
		for i := align - len; i > 0; i-- {
			b.WriteByte(32) // Space
		}
		return align
	} else {
		b.WriteByte(32) // Space
	}

	return len + 1
}

func (c *Console) writeEndValues(b *bytes.Buffer) {
	// b.WriteByte(32) // Space
}

func (c *Console) writeKey(b *bytes.Buffer, s string) {
	if c.enableColor {
		b.WriteByte(32) // Space
		c.setColor(b, "34")
		b.WriteString(s)
		c.resetColor(b)
		b.WriteByte('=')
	} else {
		b.WriteByte(32) // Space
		b.WriteString(s)
		b.WriteByte('=')
	}
}

func (c *Console) writeValue(b *bytes.Buffer, s string) {
	if c.enableColor {
		c.setColor(b, "36")
		b.WriteString(s)
		c.resetColor(b)
	} else {
		b.WriteByte('"')
		b.WriteString(s)
		b.WriteByte('"')
	}
}

func (c *Console) writeNewline(b *bytes.Buffer) {
	b.WriteByte('\r')
	b.WriteByte('\n')
}

func (c *Console) getBuffer() *bytes.Buffer {
	return c.pool.Get().(*bytes.Buffer)
}

func (c *Console) putBuffer(b *bytes.Buffer) {
	b.Reset()
	c.pool.Put(b)
}

func (c *Console) writeEnd(buf *bytes.Buffer, level Level, skipStack int) {
	c.writeNewline(buf)
	_, _ = buf.WriteTo(os.Stdout)
}

func (c *Console) Print(l Level, scope string, caller string, stack []string, message string) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)
	c.writeMessage(buf, l, scope, caller, message)
	if len(stack) > 0 {
		c.writeNewline(buf)
		for i := range stack {
			buf.WriteString("\t" + stack[i])
		}
	}
	c.writeEnd(buf, l, 3)
}

func (c *Console) Printv(l Level, scope string, caller string, stack []string, message string, keysValues []interface{}) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)

	c.writeMessage(buf, l, scope, caller, message)

	lenValues := len(keysValues)
	for i := 0; i < lenValues; i += 2 {
		if key, ok := keysValues[i].(string); ok {
			c.writeKey(buf, key)
		} else {
			c.writeKey(buf, "%KEY%")
		}
		c.writeValue(buf, fmt.Sprint(keysValues[i+1]))
	}
	c.writeEndValues(buf)
	if len(stack) > 0 {
		for i := range stack {
			buf.WriteString("\r\n    " + stack[i])
		}
		c.writeNewline(buf)
	}
	c.writeEnd(buf, l, 2)
}
