package main

import (
	"github.com/mohsensamiei/golog"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
)

const (
	service = "account"
)

var (
	version = "NON"
	configs = env.Configs{}
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetConstant("service", service)
	log.SetConstant("version", version)

	if err := configs.LoadEnv("deploy/.env"); err != nil {
		log.With(err).Fatal("can not load environment values")
	}

	if configs.IsDebugMode() {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	log.Info("running service")
}
