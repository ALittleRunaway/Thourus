package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"thourus-api/config"
	"thourus-api/domain/usecase"
	"thourus-api/entrypoint"
	"thourus-api/gateway"
)

func InitServer(appLogger *zap.SugaredLogger, grpcConn *grpc.ClientConn, dbConn *sql.DB, natsConn *nats.Conn, cfg *config.Config) (*gin.Engine, error) {

	server := gin.Default()

	server.Static("/templates", "./templates")
	server.LoadHTMLGlob("templates/*.html")

	companyGw := gateway.NewCompanyGateway(dbConn)
	spaceGw := gateway.NewSpaceGateway(dbConn)
	projectGw := gateway.NewProjectGateway(dbConn)
	documentGw := gateway.NewDocumentGateway(dbConn)
	userGw := gateway.NewUserGateway(dbConn)
	mailGw := gateway.NewMailGateway(natsConn)
	cryptoGw := gateway.NewCryptoGateway(cfg.Crypto.SecretString, cfg.Crypto.Rule, grpcConn)
	storageGw := gateway.NewStorageGateway(cfg.DB.StoragePath)

	fmt.Println(companyGw, storageGw, spaceGw, projectGw, documentGw, userGw, mailGw, cryptoGw)

	companyUc := usecase.NewCompanyUseCase(companyGw, appLogger)
	documentUc := usecase.NewDocumentUseCase(documentGw, storageGw, cryptoGw, appLogger)
	userUc := usecase.NewUserUseCase(userGw, appLogger)

	apiRoute := server.Group("/api")
	{

		apiRoute.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		apiRoute.POST("/document/upload", func(ctx *gin.Context) {
			entrypoint.UploadNewDocument(documentUc, ctx)
		})

		apiRoute.GET("/login/", func(ctx *gin.Context) {
			entrypoint.LoginUser(userUc, ctx)
		})

	}

	viewRoute := server.Group("/view")
	{
		viewRoute.GET("/login", func(ctx *gin.Context) {
			entrypoint.Login(ctx)
		})

		viewRoute.GET("/company/:uid", func(ctx *gin.Context) {
			entrypoint.GetSpacesInCompany(companyUc, ctx)
		})
	}

	go func() {
		server.Run(":9999")
	}()

	return server, nil
}
