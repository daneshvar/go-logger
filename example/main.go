package main

import (
	"time"

	"github.com/daneshvar/go-logger"
	"github.com/daneshvar/go-logger/example/rtsp"
)

func main() {
	defer logger.Close()
	logger.RedirectStdLog()

	consoleEnabler := func(l logger.Level, s string) bool { return true }
	stackEnabler := func(l logger.Level, s string) bool { return l == logger.ErrorLevel }
	logger.Config(logger.ConsoleWriter(true, stackEnabler, consoleEnabler))

	log := logger.GetLogger("example")

	log.Trace("Check Trace 1")

	log.Debug("Debug Code")
	log.Debugf("Debug Code %s", "Hello")

	log.Warn("Not Found config file")

	log.Infov("GET", "url", "http://example.com/data.json")
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	test(log)

	rtsp.GetPacketFunc()

	log.Fatal("Fatal")
}

func test(log *logger.Logger) {
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
	test2(log)
}

func test2(log *logger.Logger) {
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
