package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"thourus-api/config"
)

const serviceName = "db"

func NewDBConnection(dbCfg *config.DBConfig, logger *zap.SugaredLogger) (*sql.DB, error) {

	serviceLogger := logger.Named(serviceName)

	serviceLogger.Info("Establishing connection with the DataBase...")

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbCfg.Username, dbCfg.Password, dbCfg.Addr, dbCfg.DBName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return &sql.DB{}, err
	}

	serviceLogger.Info("Established the DataBase connection successfully.")
	return db, nil
}
