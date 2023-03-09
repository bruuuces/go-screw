package utils

import "strings"

type LogBuffer struct {
	Text    *string
	builder strings.Builder
	onWrite func()
}

func (l *LogBuffer) Write(p []byte) (int, error) {
	write, err := l.builder.Write(p)
	*l.Text = l.builder.String()
	l.onWrite()
	return write, err
}

func (l *LogBuffer) WriteString(s string) (int, error) {
	write, err := l.builder.WriteString(s)
	*l.Text = l.builder.String()
	l.onWrite()
	return write, err
}

func (l *LogBuffer) Reset() {
	l.builder.Reset()
	*l.Text = ""
	l.onWrite()
}

func NewLogBuffer(onWrite func()) *LogBuffer {
	logText := ""
	return &LogBuffer{
		Text:    &logText,
		onWrite: onWrite,
	}
}
