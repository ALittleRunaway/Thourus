package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"thourus-api/config"
)

const serviceName = "nats"

func NewNatsConnection(natsCfg *config.NatsConfig, logger *zap.SugaredLogger) (*nats.Conn, error) {

	serviceLogger := logger.Named(serviceName)

	serviceLogger.Info("Establishing the Nats connection...")

	natsConn, err := nats.Connect(natsCfg.Addr)
	if err != nil {
		return &nats.Conn{}, err
	}

	serviceLogger.Info("Established the Nats connection successfully.")
	return natsConn, nil
}
