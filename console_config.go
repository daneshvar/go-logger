package logger

type ConsoleConfig struct {
	Caller *bool
	Color  *bool
	Stack  *string
	Level  *string
	Scope  map[string]string
}
