package services

import (
	"context"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/app/shortener/models"
	"github.com/mohsensamiei/link-shortener/internal/pkg/auth"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/token"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func NewLink(ctx context.Context, configs env.Configs) api.LinkServer {
	return &link{
		Service: auth.Service{
			Context: ctx,
		},
		Configs: configs,
	}
}

type link struct {
	auth.Service
	Configs env.Configs
}

func (s link) LinkRepository() models.LinkRepository {
	return s.Context.Value(models.LinkRepositoryInstance{}).(models.LinkRepository)
}

func (s link) Create(ctx context.Context, req *api.LinkCreateRequest) (*api.LinkCreateResponse, error) {
	cred, err := s.Authorize(ctx)
	if err != nil {
		return nil, err
	}
	link := &models.Link{
		UserID: uuid.FromStringOrNil(cred.GetId()),
		Url:    req.Url,
	}
	for i := 0; i < token.StatesWithPrefix(s.Configs.ShortenerSlugLen, req.Slug); i++ {
		link.Slug = token.GenerateWithPrefix(s.Configs.ShortenerSlugLen, req.Slug)
		switch err := s.LinkRepository().Create(link); errors.Cause(err) {
		case nil:
			return &api.LinkCreateResponse{
				Slug: link.Slug,
			}, nil
		case errorsext.ErrConflict:
			break
		default:
			return nil, err
		}
	}
	return nil, errors.Wrapf(errorsext.ErrConflict, "can not create short link with slug '%v'", link.Slug)
}

func (s link) Return(_ context.Context, req *api.LinkReturnRequest) (*api.LinkReturnResponse, error) {
	link, err := s.LinkRepository().ReturnBySlug(req.Slug)
	if err != nil {
		return nil, err
	}
	return &api.LinkReturnResponse{
		Url: link.Url,
	}, err
}
