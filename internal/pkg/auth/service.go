package auth

import (
	"context"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/pkg/errors"
)

type Service struct {
	Context context.Context
}

func (s Service) AuthenticateClient() api.AuthenticateClient {
	return s.Context.Value(api.AuthenticateClientInstance{}).(api.AuthenticateClient)
}

func (s Service) Authorize(ctx context.Context) (Credentials, error) {
	token, ok := GetTokenFromContext(ctx)
	if ok == false {
		return nil, errors.Wrap(errorsext.ErrUnauthorized, "there is not token set in context")
	}
	res, err := s.AuthenticateClient().Check(ctx, &api.AuthenticateCheckRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
