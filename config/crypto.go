package config

import (
	"github.com/spf13/viper"
)

type CryptoConfig struct {
	SecretString string
	Rule         int64
}

func InitCryptoConfig() *CryptoConfig {

	cryptoConfig := CryptoConfig{
		SecretString: viper.GetString("cryptocore.secret_string"),
		Rule:         viper.GetInt64("cryptocore.rule"),
	}
	return &cryptoConfig
}

func init() {
	viper.SetDefault("cryptocore.rule", 54321)
	InitError(viper.BindEnv("cryptocore.secret_string", EnvPrefix+"_CRYPTO_SECRET_STRING"))
	InitError(viper.BindEnv("cryptocore.rule", EnvPrefix+"_CRYPTO_RULE"))
}
