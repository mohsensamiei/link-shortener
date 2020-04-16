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
		AllowMethods: []string{"OPTIONS", "POST"},
	}))

	authenticateCtrl := NewAuthenticate(ctx, configs)
	engine.GET("/authenticates", authenticateCtrl.Check)
	engine.POST("/authenticates", authenticateCtrl.Login)
	engine.POST("/authenticates/register", authenticateCtrl.Register)
}
