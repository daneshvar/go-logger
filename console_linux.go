package logger

import "bytes"

func (c *Console) resetColor(b *bytes.Buffer) {
	b.WriteString("\x1B[0m")
}

func (c *Console) setColor(b *bytes.Buffer, fg string) {
	b.WriteString("\x1B[" + fg + "m")
}

func (c *Console) writeNewline(b *bytes.Buffer) {
	b.WriteByte('\n')
}
