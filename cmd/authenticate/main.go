package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mohsensamiei/golog"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/app/authenticate/controllers"
	"github.com/mohsensamiei/link-shortener/internal/app/authenticate/models"
	"github.com/mohsensamiei/link-shortener/internal/app/authenticate/services"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/mohsensamiei/link-shortener/pkg/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	service = "authenticate"
)

var (
	version = "NON"
	configs = env.Configs{}
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetConstant("service_name", service)
	log.SetConstant("service_version", version)

	if err := configs.LoadEnv("deploy/.env"); err != nil {
		log.With(err).Fatal("can not load environment values")
	}

	if configs.IsDebugMode() {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	log.Info("running service")

	ctx := context.Background()

	postgresDB, err := configs.ConnectPostgres(configs.AuthenticatePostgresDB)
	if err != nil {
		log.With(err).Fatal("can not connect to postgres")
	}
	defer func() {
		if err := postgresDB.Close(); err != nil {
			log.With(err).Fatal("can not disconnect from postgres")
		}
	}()
	if err := models.AutoMigrate(postgresDB); err != nil {
		log.With(err).Fatal("can not migrate database")
	}
	ctx = context.WithValue(ctx, models.UserRepositoryInstance{}, models.NewUserRepository(postgresDB))

	authenticateConn, err := grpc.DialContext(ctx, configs.AuthenticateGrpcHost, grpc.WithInsecure())
	if err != nil {
		log.With(errors.WithStack(err)).Fatal("can not connect to grpc client")
	}
	defer func() {
		if err := authenticateConn.Close(); err != nil {
			log.With(errors.WithStack(err)).Fatal("can not disconnect from grpc client")
		}
	}()
	ctx = context.WithValue(ctx, api.AuthenticateClientInstance{}, api.NewAuthenticateClient(authenticateConn))

	serving := server.NewServer()
	if err := serving.WithGrpc(configs.AuthenticateGrpcHost, func(server *grpc.Server) {
		api.RegisterAuthenticateServer(server, services.NewAuthenticate(ctx, configs))
	}); err != nil {
		log.With(err).Fatal("can not add grpc to server")
	}
	if err := serving.WithHttp(configs.AuthenticateHttpHost, func(engine *gin.Engine) {
		controllers.RegisterRoutes(engine, ctx, configs)
	}); err != nil {
		log.With(err).Fatal("can not add http to server")
	}
	if err := serving.Run(); err != nil {
		log.With(err).Fatal("service interrupted")
	}
}
