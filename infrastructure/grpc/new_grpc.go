package grpc

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"thourus-api/config"
)

const serviceName = "grpc"

func NewGrpcConnection(grpcCfg *config.GrpcConfig, logger *zap.SugaredLogger) (*grpc.ClientConn, error) {

	serviceLogger := logger.Named(serviceName)

	serviceLogger.Info("Establishing the gRPC connection...")

	conn, err := grpc.Dial(grpcCfg.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &grpc.ClientConn{}, err
	}

	serviceLogger.Info("Established the gRPC connection successfully.")
	return conn, nil
}
