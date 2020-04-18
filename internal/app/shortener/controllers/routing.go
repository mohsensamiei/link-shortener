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
		AllowMethods: []string{"OPTIONS", "POST", "GET"},
	}))

	linkCtrl := NewLink(ctx, configs)
	engine.GET("/r/:slug", linkCtrl.Redirect)
	engine.POST("/links", linkCtrl.Create)
}
