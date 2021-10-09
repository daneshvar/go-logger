package logger

type ConsoleConfig struct {
	Caller *bool
	Color  *bool
	Level  *string
	Scope  map[string]string
}
