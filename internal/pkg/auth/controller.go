package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/pkg/ginext"
)

type Controller struct {
	Context context.Context
}

func (c Controller) AuthenticateClient() api.AuthenticateClient {
	return c.Context.Value(api.AuthenticateClientInstance{}).(api.AuthenticateClient)
}

func (c Controller) ContextWithToken(ctx *gin.Context) context.Context {
	return AddTokenToContext(c.Context, ginext.GetAuthBearerToken(ctx))
}

func (c Controller) Authorize(ctx *gin.Context) (Credentials, error) {
	res, err := c.AuthenticateClient().Check(ctx, &api.AuthenticateCheckRequest{
		Token: ginext.GetAuthBearerToken(ctx),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
