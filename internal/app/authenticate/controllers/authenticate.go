package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/ginext"
)

type Authenticate interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Check(ctx *gin.Context)
}

func NewAuthenticate(ctx context.Context, configs env.Configs) Authenticate {
	return &authenticate{
		Context: ctx,
		Configs: configs,
	}
}

type authenticate struct {
	Context context.Context
	Configs env.Configs
}

func (c authenticate) AuthenticateClient() api.AuthenticateClient {
	return c.Context.Value(api.AuthenticateClientInstance{}).(api.AuthenticateClient)
}

func (c authenticate) Register(ctx *gin.Context) {
	req := new(api.AuthenticateRegisterRequest)
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctx.Error(err))
		return
	}
	if _, err := c.AuthenticateClient().Register(c.Context, req); err != nil {
		if errorsext.IsValidation(err) {
			ctx.JSON(http.StatusBadRequest, ctx.Error(err))
		} else if errorsext.IsConflict(err) {
			ctx.JSON(http.StatusConflict, ctx.Error(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		}
		return
	}
	ctx.Status(http.StatusCreated)
}

func (c authenticate) Login(ctx *gin.Context) {
	req := new(api.AuthenticateLoginRequest)
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctx.Error(err))
		return
	}
	res, err := c.AuthenticateClient().Login(c.Context, req)
	if err != nil {
		if errorsext.IsUnauthorized(err) {
			ctx.JSON(http.StatusUnauthorized, ctx.Error(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c authenticate) Check(ctx *gin.Context) {
	req := &api.AuthenticateCheckRequest{
		Token: ginext.GetAuthBearerToken(ctx),
	}
	res, err := c.AuthenticateClient().Check(c.Context, req)
	if err != nil {
		if errorsext.IsUnauthorized(err) {
			ctx.JSON(http.StatusUnauthorized, ctx.Error(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		}
		return
	}
	ctx.JSON(http.StatusOK, res)
}
