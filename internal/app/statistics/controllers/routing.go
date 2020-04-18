package controllers

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
)

func RegisterRoutes(engine *gin.Engine, ctx context.Context, configs env.Configs) {
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET"},
	}))

	visitCtrl := NewVisit(ctx, configs)
	engine.GET("/visits/url", visitCtrl.ReportByUrl)
	engine.GET("/visits/user", visitCtrl.ReportByUser)
}
