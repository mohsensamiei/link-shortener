package test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	log "github.com/mohsensamiei/golog"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type tokenKey struct{}

var (
	configs = env.Configs{}
	ctx     = context.Background()
)

func TestMain(m *testing.M) {
	fmt.Println("initializing integration tests...")
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	if err := configs.LoadEnv("deploy/.env"); err != nil {
		log.With(err).Fatal("can not load environment values")
	}

	authenticateConn, err := grpc.DialContext(ctx, configs.AuthenticateGrpcHost, grpc.WithInsecure())
	if err != nil {
		log.With(errors.WithStack(err)).Fatal("can not connect to grpc client")
	}
	defer func() {
		if err := authenticateConn.Close(); err != nil {
			log.With(errors.WithStack(err)).Warning("can not disconnect from grpc client")
		}
	}()
	ctx = context.WithValue(ctx, api.AuthenticateClientInstance{}, api.NewAuthenticateClient(authenticateConn))

	token, err := authenticate()
	if err != nil {
		log.With(err).Fatal("can not authenticate test user")
	}
	ctx = context.WithValue(ctx, tokenKey{}, token)

	go func() {
		serve()
	}()
}

func shutdown() {

}

func authenticate() (string, error) {
	authenticateClient := ctx.Value(api.AuthenticateClientInstance{}).(api.AuthenticateClient)
	authenticateClient.Register(ctx, &api.AuthenticateRegisterRequest{
		Email:    "test@example.com",
		Username: "test",
		Password: "test",
	})
	res, err := authenticateClient.Login(ctx, &api.AuthenticateLoginRequest{
		Email:    "test@example.com",
		Username: "test",
		Password: "test",
	})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

func serve() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	engine := gin.Default()
	engine.Handle(http.MethodGet, "/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	http.ListenAndServe(":8080", engine)
}
