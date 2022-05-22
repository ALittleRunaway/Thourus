package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	ServiceName = "thourus.api"
	EnvPrefix   = "THOURUS_API"
)

type Config struct {
	Log    *LogConfig
	Grpc   *GrpcConfig
	Crypto *CryptoConfig
	Nats   *NatsConfig
	DB     *DBConfig
}

func InitError(err error) {
	if err != nil {
		panic(err)
	}
}

// InitConfig initialise all the configurations
func InitConfig() *Config {
	return &Config{
		Log:    InitLogConfig(),
		Grpc:   InitGrpcConfig(),
		Crypto: InitCryptoConfig(),
		Nats:   InitNatsConfig(),
		DB:     InitDBConfig(),
	}
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("There is no .env file provided!")
	}
	viper.SetDefault("thourus.api.dedicated_env", "DEV")
	InitError(viper.BindEnv("thourus.api.dedicated_env", EnvPrefix+"_DEDICATED_ENV"))
}
