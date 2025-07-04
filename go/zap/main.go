package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	url := "shivampatel.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("HHHHHHHFailed to fetch URL: %s", url)
}
