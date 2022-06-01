//Copyright 2022 Maria Petrova marycool674@gmail.com
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package main

import (
	"thourus-api/config"
	"thourus-api/infrastructure"
	"thourus-api/infrastructure/database"
	"thourus-api/infrastructure/grpc"
	"thourus-api/infrastructure/log"
	"thourus-api/infrastructure/nats"
)

func main() {

	cfg := config.InitConfig()

	// New logger
	appLogger, err := log.NewLogger(cfg.Log.Level)
	if err != nil {
		panic(err)
	}

	appLogger.Info("The app is starting...")
	fatalErrCh := make(chan error, 2)

	// Establishing the gRPC connection
	grpcConn, err := grpc.NewGrpcConnection(cfg.Grpc, appLogger)
	if err != nil {
		appLogger.Error("The app could not establish the gRPC connection. Exiting")
		fatalErrCh <- err
	}

	// Establishing the DataBase connection
	DBConn, err := database.NewDBConnection(cfg.DB, appLogger)
	if err != nil {
		appLogger.Error("The app could not establish the DataBase connection. Exiting")
		fatalErrCh <- err
	}

	// Establishing the Nats connection
	NatsConn, err := nats.NewNatsConnection(cfg.Nats, appLogger)
	if err != nil {
		appLogger.Error("The app could not establish the Nats connection. Exiting")
		fatalErrCh <- err
	}

	// Starting the server
	server, err := InitServer(appLogger, grpcConn, DBConn, NatsConn, cfg)
	if err != nil {
		appLogger.Error("The app could not start the server. Exiting")
		fatalErrCh <- err
	}

	appLogger.Info("The app has started")

	params := infrastructure.InterruptParams{
		Logger:   appLogger,
		Shutdown: fatalErrCh,
		GrpcConn: grpcConn,
		DBConn:   DBConn,
		NatsConn: NatsConn,
		Server:   server,
	}
	infrastructure.Interrupter(params)

	appLogger.Info("The app is stopped.")
}
