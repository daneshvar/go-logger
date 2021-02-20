# go-loggerger

```go
package main

import "github.com/daneshvar/go-logger"

func main() {
	defer logger.Close()
	logger.RedirectStdLog()

	consoleEnabler := func(l logger.Level, s string) bool { return true }
	stackEnabler := func(l logger.Level, s string) bool { return l == logger.ErrorLevel }
	logger.Config(logger.ConsoleWriter(true, stackEnabler, consoleEnabler))

	log := logger.GetLogger("example")

	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)

	log.Error("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
```
