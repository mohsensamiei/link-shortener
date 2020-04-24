package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/mohsensamiei/golog"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/pkg/auth"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/mohsensamiei/link-shortener/pkg/cash"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/ginext"
	"github.com/mssola/user_agent"
)

type Link interface {
	Create(ctx *gin.Context)
	Redirect(ctx *gin.Context)
}

type visitSeed struct {
	Url   string
	Token string
	Agent string
}

func NewLink(ctx context.Context, configs env.Configs) Link {
	ctrl := &link{
		Controller: auth.Controller{
			Context: ctx,
		},
		Configs: configs,
		Visits:  make(chan *visitSeed),
	}
	for i := 0; i < ctrl.Configs.AsyncCount; i++ {
		go func() {
			for visit := range ctrl.Visits {
				var userID string
				res, err := ctrl.AuthenticateClient().Check(ctrl.Context, &api.AuthenticateCheckRequest{
					Token: visit.Token,
				})
				if err == nil {
					userID = res.Id
				}

				ua := user_agent.New(visit.Agent)
				mobile := ua.Mobile()
				browser, _ := ua.Browser()

				if _, err := ctrl.VisitClient().Register(ctrl.Context, &api.VisitRegisterRequest{
					Url:      visit.Url,
					UserId:   userID,
					IsMobile: mobile,
					Browser:  browser,
				}); err != nil {
					log.With(err).Error("can not register visits")
				}
			}
		}()
	}
	return ctrl
}

type link struct {
	auth.Controller
	Configs env.Configs
	Visits  chan *visitSeed
}

func (c link) LinkClient() api.LinkClient {
	return c.Context.Value(api.LinkClientInstance{}).(api.LinkClient)
}
func (c link) CashManager() cash.Manager {
	return c.Context.Value(cash.ManagerInstance{}).(cash.Manager)
}
func (c link) VisitClient() api.VisitClient {
	return c.Context.Value(api.VisitClientInstance{}).(api.VisitClient)
}

func (c link) Create(ctx *gin.Context) {
	req := new(api.LinkCreateRequest)
	if err := ctx.Bind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, ctx.Error(err))
		return
	}
	res, err := c.LinkClient().Create(c.ContextWithToken(ctx), req)
	if err != nil {
		if errorsext.IsValidation(err) {
			ctx.JSON(http.StatusBadRequest, ctx.Error(err))
		} else if errorsext.IsUnauthorized(err) {
			ctx.JSON(http.StatusUnauthorized, ctx.Error(err))
		} else if errorsext.IsConflict(err) {
			ctx.JSON(http.StatusConflict, ctx.Error(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
		}
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
func (c link) Redirect(ctx *gin.Context) {
	slug := ctx.Param("slug")
	res := new(api.LinkReturnResponse)
	if err := c.CashManager().Get(slug, res); err != nil {
		if res, err = c.LinkClient().Return(c.Context, &api.LinkReturnRequest{
			Slug: slug,
		}); err != nil {
			if errorsext.IsNotFound(err) {
				ctx.JSON(http.StatusNotFound, ctx.Error(err))
			} else {
				ctx.JSON(http.StatusInternalServerError, ctx.Error(err))
			}
			return
		}
		if err := c.CashManager().Set(slug, res, c.Configs.ShortenerCashDuration); err != nil {
			log.With(err).Warning("can not set redirect cash")
		}
	}
	c.Visits <- &visitSeed{
		Url:   res.Url,
		Token: ginext.GetAuthBearerToken(ctx),
		Agent: ctx.Request.UserAgent(),
	}
	ctx.Redirect(http.StatusPermanentRedirect, res.Url)
}
