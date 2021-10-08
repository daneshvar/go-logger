package rtsp

import (
	"time"

	"github.com/daneshvar/go-logger"
)

func GetPacketFunc(log *logger.Logger) {
	log = log.GetLogger("rtsp")

	//log.Info("Namitonam fetch konam")
	//
	//log.Infof("Namitonam fetch konam %s", "Hossein")
	log.Warn("Not Found config file")

	log.Infov("GET",
		"url", "http://example.com/data.json",
	)
	log.Errorv("Fetch",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	log.Infov("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)

	log.Errorv("Namitonam",
		"url", "http://example.com",
		"attempt", 3,
		"backoff", time.Second,
	)
}
