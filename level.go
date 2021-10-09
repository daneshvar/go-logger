package logger

import "strings"

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

var (
	levelText = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC"}
	toLevel   = map[string]Level{
		"TRACE":   TraceLevel,
		"DEBUG":   DebugLevel,
		"INFO":    InfoLevel,
		"WARN":    WarnLevel,
		"WARNING": WarnLevel,
		"ERR":     ErrorLevel,
		"ERROR":   ErrorLevel,
		"FATAL":   FatalLevel,
		"PANIC":   PanicLevel,
	}
)

// ToLevel map text of level to log.Level
// Trace, Debug, Info, Warn or Warning, Error or Err, Fatal, Panic
func ToLevel(level string) (Level, bool) {
	if l, ok := toLevel[strings.ToUpper(level)]; ok {
		return l, true
	}

	return TraceLevel, false
}

func LevelText(level Level) string {
	return levelText[level]
}
