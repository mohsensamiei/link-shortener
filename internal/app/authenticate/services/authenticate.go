package services

import (
	"context"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/app/authenticate/models"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/hashext"
	"github.com/mohsensamiei/link-shortener/pkg/jwtext"
	"github.com/pkg/errors"
)

const (
	invalidCredentials = "invalid email/username or password"
	invalidToken       = "invalid authentication token"
)

func NewAuthenticate(ctx context.Context, configs env.Configs) api.AuthenticateServer {
	return &authenticate{
		Context: ctx,
		Configs: configs,
	}
}

type authenticate struct {
	Context context.Context
	Configs env.Configs
}

func (s authenticate) UserRepository() models.UserRepository {
	return s.Context.Value(models.UserRepositoryInstance{}).(models.UserRepository)
}

func (s authenticate) Login(_ context.Context, req *api.AuthenticateLoginRequest) (*api.AuthenticateLoginResponse, error) {
	user, err := s.UserRepository().ReturnByEmailOrUsername(req.Email, req.Username)
	switch errors.Cause(err) {
	case nil:
	case errorsext.ErrNotFound:
		return nil, errors.Wrap(errorsext.ErrUnauthorized, invalidCredentials)
	default:
		return nil, err
	}

	if user.PasswordHashed != hashext.Encode(req.Password) {
		return nil, errors.Wrap(errorsext.ErrUnauthorized, invalidCredentials)
	}

	claims := jwtext.NewClaims(user.ID, user.Email, user.Username, s.Configs.AuthenticateTokenDuration)
	token, err := claims.Token(s.Configs.AuthenticateTokenKey)
	if err != nil {
		return nil, err
	}

	return &api.AuthenticateLoginResponse{
		Token:    token,
		Duration: int32(s.Configs.AuthenticateTokenDuration.Seconds()),
	}, nil
}

func (s authenticate) Register(_ context.Context, req *api.AuthenticateRegisterRequest) (*api.VoidResponse, error) {
	user := models.User{
		Email:          req.Email,
		Username:       req.Username,
		PasswordHashed: hashext.Encode(req.Password),
	}
	if err := s.UserRepository().Create(&user); err != nil {
		return nil, err
	}
	return &api.VoidResponse{}, nil
}

func (s authenticate) Check(_ context.Context, req *api.AuthenticateCheckRequest) (*api.AuthenticateCheckResponse, error) {
	claims, err := jwtext.LoadClaims(s.Configs.AuthenticateTokenKey, req.Token)
	if err != nil {
		return nil, errors.Wrap(errorsext.ErrUnauthorized, err.Error())
	}
	return &api.AuthenticateCheckResponse{
		Email:    claims.Email,
		Username: claims.Username,
	}, nil
}
