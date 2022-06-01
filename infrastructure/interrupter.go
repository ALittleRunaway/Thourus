package infrastructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

type InterruptParams struct {
	Logger   *zap.SugaredLogger
	Shutdown <-chan error
	GrpcConn *grpc.ClientConn
	DBConn   *sql.DB
	NatsConn *nats.Conn
	Server   *gin.Engine
}

func Interrupter(params InterruptParams) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-interrupt:
		params.Logger.Infow("Received", "signal", sig.String())
	case err := <-params.Shutdown:
		params.Logger.Errorw("Received error from functional unit", "err", err)
	}

	params.Logger.Info("Stopping the app...")

	// stopping the gRPC connection
	if params.GrpcConn != nil {
		params.GrpcConn.Close()
		params.Logger.Info("gRPC connection is closed.")
	}

	// stopping the DataBase connection
	if params.DBConn != nil {
		params.DBConn.Close()
		params.Logger.Info("DataBase connection is closed.")
	}

	// stopping the Nats connection
	if params.NatsConn != nil {
		if params.NatsConn.IsConnected() {
			params.NatsConn.Drain()
			params.NatsConn.Close()
		}
		params.Logger.Info("Nats connection is closed.")
	}

	// stopping the server
	if params.Server != nil {
		//params.Server.Close()
		params.Logger.Info("Server is stopped.")
	}

}
