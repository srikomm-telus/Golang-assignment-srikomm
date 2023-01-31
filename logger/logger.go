package logger

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	const url = "https://example.com"

	// In most circumstances, use the SugaredLogger. It's 4-10x faster than most
	// others structured logging packages and has a familiar, loosely-typed API.
	sugar := logger.Sugar()
	sugar.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	testString := "testString"
	sugar.Infof("Failed to fetch URL: %s", url)
	sugar.Warnf("test warning %s", testString)
	sugar.Debug(testString)

	// In the unusual situations where every microsecond matters, use the
	// Logger. It's even faster than the SugaredLogger, but only supports
	// structured logging.
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
