package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Addr        string
	DBName      string
	Username    string
	Password    string
	StoragePath string
}

func InitDBConfig() *DBConfig {

	dbConfig := DBConfig{
		Addr:        viper.GetString("db.addr"),
		DBName:      viper.GetString("db.db_name"),
		Username:    viper.GetString("db.username"),
		Password:    viper.GetString("db.password"),
		StoragePath: viper.GetString("db.storage_path"),
	}
	return &dbConfig
}

func init() {
	viper.SetDefault("db.addr", "localhost:3306")
	viper.SetDefault("db.db_name", "thourus")
	viper.SetDefault("db.storage_path", "storage/")

	InitError(viper.BindEnv("db.addr", EnvPrefix+"_DB_ADDR"))
	InitError(viper.BindEnv("db.db_name", EnvPrefix+"_DB_NAME"))
	InitError(viper.BindEnv("db.username", EnvPrefix+"_DB_USERNAME"))
	InitError(viper.BindEnv("db.password", EnvPrefix+"_DB_PASSWORD"))
	InitError(viper.BindEnv("db.storage_path", EnvPrefix+"_STORAGE_PATH"))
}
