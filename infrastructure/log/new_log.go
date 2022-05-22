package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"thourus-api/config"
)

func NewLogger(level zapcore.Level) (*zap.SugaredLogger, error) {
	var logConfig zap.Config

	if os.Getenv(config.EnvPrefix+"_DEDICATED_ENV") == "DEV" {
		logConfig = zap.NewDevelopmentConfig()
	} else {
		logConfig = zap.NewProductionConfig()
	}

	logConfig.EncoderConfig.LineEnding = "\n\n"
	logConfig.Level = zap.NewAtomicLevelAt(level)

	logger, _ := logConfig.Build()

	return logger.Sugar().Named(config.ServiceName), nil
}
