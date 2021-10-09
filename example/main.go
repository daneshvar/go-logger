package main

import (
	"time"

	"github.com/daneshvar/go-logger"
)

func main() {
	core := logger.New()
	defer core.Close()

	core.RedirectStdLogAt("other", logger.ErrorLevel)

	core.Config(logger.ConsoleWriter(true))

	log := core.GetLogger("example")

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
	test(log.GetLogger("skip +1").Skip(1))

	// rtsp.GetPacketFunc(log)

	// log.Fatal("Fatal")
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
