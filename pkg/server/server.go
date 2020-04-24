package server

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	log "github.com/mohsensamiei/golog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server interface {
	WithGrpc(host string, registering func(server *grpc.Server)) error
	WithHttp(host string, registering func(engine *gin.Engine)) error
	Run() error
}

func NewServer() Server {
	return &server{}
}

type server struct {
	GrpcServer   *grpc.Server
	GrpcListener net.Listener
	HttpServer   *gin.Engine
	HttpListener net.Listener
}

func (s *server) WithGrpc(host string, registering func(server *grpc.Server)) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return errors.WithStack(err)
	}
	s.GrpcListener = listener

	server := grpc.NewServer()
	registering(server)
	s.GrpcServer = server

	return nil
}

func (s *server) WithHttp(host string, registering func(engine *gin.Engine)) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return errors.WithStack(err)
	}
	s.HttpListener = listener

	engine := gin.Default()
	registering(engine)
	s.HttpServer = engine

	return nil
}

func (s server) Run() error {
	interrupt := make(chan error)
	if s.GrpcListener != nil && s.GrpcServer != nil {
		go func(interrupt chan error) {
			log.Data("serve", fmt.Sprint(s.GrpcListener.Addr())).Info("start listening grpc")
			if err := s.GrpcServer.Serve(s.GrpcListener); err != nil {
				interrupt <- errors.Wrap(err, "can not serve grpc")
			}
		}(interrupt)
	}
	if s.HttpListener != nil && s.HttpServer != nil {
		go func(interrupt chan error) {
			log.Data("serve", fmt.Sprint("http://", s.HttpListener.Addr())).Info("start listening http")
			if err := s.HttpServer.RunListener(s.HttpListener); err != nil {
				interrupt <- errors.Wrap(err, "can not serve http")
			}
		}(interrupt)
	}
	return <-interrupt
}
