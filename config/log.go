package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	Level zapcore.Level
}

func InitLogConfig() *LogConfig {
	var logLevel zapcore.Level

	logLevelStr := viper.GetString("log.log_level")
	if err := logLevel.UnmarshalText([]byte(logLevelStr)); err != nil {
		panic(err)
	}

	return &LogConfig{
		Level: logLevel,
	}
}

func init() {
	viper.SetDefault("log.log_level", "DEBUG")
	InitError(viper.BindEnv("log.log_level", EnvPrefix+"_LOG_LEVEL"))
}
