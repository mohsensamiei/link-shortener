package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/pkg/auth"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"net/http"
)

type Visit interface {
	ReportByUrl(ctx *gin.Context)
	ReportByUser(ctx *gin.Context)
}

func NewVisit(ctx context.Context, configs env.Configs) Visit {
	return &visit{
		Controller: auth.Controller{
			Context: ctx,
		},
		Configs: configs,
	}
}

type visit struct {
	auth.Controller
	Configs env.Configs
}

func (c visit) VisitClient() api.VisitClient {
	return c.Context.Value(api.VisitClientInstance{}).(api.VisitClient)
}

func (c visit) ReportByUrl(ctx *gin.Context) {
	res, err := c.VisitClient().ReportByUrl(c.Context, &api.VoidRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c visit) ReportByUser(ctx *gin.Context) {
	res, err := c.VisitClient().ReportByUser(c.Context, &api.VoidRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
