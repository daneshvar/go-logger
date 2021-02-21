package logger

import "bytes"

func (c *Console) resetColor(b *bytes.Buffer) {
}

func (c *Console) setColor(b *bytes.Buffer, fg string) {
}

func (c *Console) writeNewline(b *bytes.Buffer) {
	b.WriteByte('\r')
	b.WriteByte('\n')
}
