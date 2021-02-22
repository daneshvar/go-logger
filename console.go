package logger

import (
	"bytes"
	"fmt"
	"io"
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

type WriteSync interface {
	io.Writer
	Sync() error
}

type Console struct {
	pool        sync.Pool
	enableColor bool
	scopeAlign  int
	callerAlign int
	wr          WriteSync
	wrLock      sync.Mutex

	writePrefix func(b *bytes.Buffer, l Level, scope string, caller string)
	writeKey    func(b *bytes.Buffer, k interface{})
	writeValue  func(b *bytes.Buffer, v interface{})
	writeScope  func(b *bytes.Buffer, scope string)
	writeCaller func(b *bytes.Buffer, caller string)
}

func ConsoleWriter(caller bool, stack EnablerFunc, enabler EnablerFunc) *Writer {
	return ConsoleWriterWithOptions(caller, stack, enabler, true, defScopeAlign, defCallerAlign)
}

// use default value of scopeAlign & callerAlign with set they with -1 and 0 to disable
func ConsoleWriterWithOptions(caller bool, stack EnablerFunc, enabler EnablerFunc, enableColor bool, scopeAlign int, callerAlign int) *Writer {
	if scopeAlign < 0 {
		scopeAlign = defScopeAlign
	}

	if callerAlign < 0 {
		callerAlign = defCallerAlign
	}

	c := &Console{
		pool: sync.Pool{New: func() interface{} {
			b := bytes.NewBuffer(make([]byte, 150))
			b.Reset()
			return b
		}},
		enableColor: enableColor,
		scopeAlign:  scopeAlign,
		callerAlign: callerAlign,
		wr:          os.Stderr,
	}

	if c.enableColor {
		c.writePrefix = c.writePrefixColor
		c.writeKey = c.writeKeyColor
		c.writeValue = c.writeValueColor
	} else {
		c.writePrefix = c.writePrefixSimple
		c.writeKey = c.writeKeySimple
		c.writeValue = c.writeValueSimple
	}

	if c.scopeAlign > 0 {
		c.writeScope = c.writeScopeAlign
	} else {
		c.writeScope = c.writeScopeSimple
	}

	if c.callerAlign > 0 {
		c.writeCaller = c.writeCallerAlign
	} else {
		c.writeCaller = c.writeCallerSimple
	}

	return newWriter(enabler, stack, caller, c)
}

func (c *Console) Print(l Level, scope string, caller string, stack []string, messages []interface{}) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)

	c.writePrefix(buf, l, scope, caller)
	fmt.Fprint(buf, messages...)
	c.writeStack(buf, stack)
	c.writeEnd(buf, l, 2)
}

func (c *Console) Prints(l Level, scope string, caller string, stack []string, message string) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)

	c.writePrefix(buf, l, scope, caller)
	buf.WriteString(message)
	c.writeStack(buf, stack)
	c.writeEnd(buf, l, 2)
}

func (c *Console) Printf(l Level, scope string, caller string, stack []string, format string, args []interface{}) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)

	c.writePrefix(buf, l, scope, caller)
	fmt.Fprintf(buf, format, args...)
	c.writeStack(buf, stack)
	c.writeEnd(buf, l, 2)
}

func (c *Console) Printv(l Level, scope string, caller string, stack []string, message string, keysValues []interface{}) {
	buf := c.getBuffer()
	defer c.putBuffer(buf)

	c.writePrefix(buf, l, scope, caller)
	buf.WriteString(message)
	c.writeValues(buf, keysValues)
	c.writeStack(buf, stack)
	c.writeEnd(buf, l, 2)
}

func (c *Console) close() {
	_ = c.wr.Sync()
}

func (c *Console) writePrefixColor(b *bytes.Buffer, l Level, scope string, caller string) {
	b.WriteString(time.Now().Format("2006-01-02 15:04:05"))

	c.setColor(b, consoleLevelColor[l])
	b.WriteString(consoleLevelText[l])

	c.writeScope(b, scope)
	c.writeCaller(b, caller)

	c.resetColor(b)
}

func (c *Console) writeKeyColor(b *bytes.Buffer, k interface{}) {
	b.WriteByte(32) // Space
	c.setColor(b, "34")
	fmt.Fprint(b, k)
	c.resetColor(b)
	b.WriteByte('=')
}

func (c *Console) writeValueColor(b *bytes.Buffer, v interface{}) {
	c.setColor(b, "36")
	fmt.Fprint(b, v)
	c.resetColor(b)
}

func (c *Console) writePrefixSimple(b *bytes.Buffer, l Level, scope string, caller string) {
	b.WriteString(time.Now().Format("2006-01-02 15:04:05"))

	b.WriteString(consoleLevelText[l])

	c.writeScope(b, scope)
	c.writeCaller(b, caller)
}

func (c *Console) writeKeySimple(b *bytes.Buffer, k interface{}) {
	b.WriteByte(' ')
	fmt.Fprint(b, k)
	b.WriteByte('=')
}

func (c *Console) writeValueSimple(b *bytes.Buffer, v interface{}) {
	b.WriteByte('"')
	fmt.Fprint(b, v)
	b.WriteByte('"')
}

func (c *Console) writeScopeAlign(b *bytes.Buffer, scope string) {
	if scope != "" {
		b.WriteString("[" + scope + "]")
		c.writeAlign(c.scopeAlign, len(scope)+2, b)
	} else {
		c.writeAlign(c.scopeAlign, 0, b)
	}
}

func (c *Console) writeCallerAlign(b *bytes.Buffer, caller string) {
	if caller != "" {
		b.WriteString(caller)
		c.writeAlign(c.callerAlign, len(caller), b)
	}
}

func (c *Console) writeScopeSimple(b *bytes.Buffer, scope string) {
	if scope != "" {
		b.WriteString("[" + scope + "]  ")
	}
}

func (c *Console) writeCallerSimple(b *bytes.Buffer, caller string) {
	if caller != "" {
		b.WriteString(caller + "  ")
	}
}

func (c *Console) writeAlign(align int, len int, b *bytes.Buffer) {
	if len < align {
		for i := align - len; i > 0; i-- {
			b.WriteByte(32)
		}
	} else {
		b.WriteByte(32)
	}
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
	c.writeBuf(buf)
	if level >= ErrorLevel {
		_ = c.wr.Sync()
	}
}

func (c *Console) writeBuf(buf *bytes.Buffer) {
	c.wrLock.Lock()
	defer c.wrLock.Unlock()
	_, _ = c.wr.Write(buf.Bytes())
}

func (c *Console) writeValues(buf *bytes.Buffer, keysValues []interface{}) {
	lenValues := len(keysValues)
	for i := 0; i < lenValues; i++ {
		c.writeKey(buf, keysValues[i])

		i++
		if i < lenValues {
			c.writeValue(buf, keysValues[i])
		} else {
			c.writeValue(buf, "!VALUE")
		}
	}
}

func (c *Console) writeStack(buf *bytes.Buffer, stack []string) {
	if len(stack) > 0 {
		for i := range stack {
			c.writeNewline(buf)
			buf.WriteString("    " + stack[i])
		}
		c.writeNewline(buf)
	}
}
