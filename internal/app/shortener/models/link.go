package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/gormext"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Link struct {
	gormext.Model

	UserID uuid.UUID `gorm:"not null;type:uuid;index"`
	Slug   string    `gorm:"not null;size:25;unique"`
	Url    string    `gorm:"not null;size:250"`
}

type LinkRepositoryInstance struct{}
type LinkRepository interface {
	Create(model *Link) error
	ReturnBySlug(slug string) (*Link, error)
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkRepository{
		db,
	}
}

type linkRepository struct {
	DB *gorm.DB
}

func (repo linkRepository) Create(model *Link) error {
	var count int
	if repo.DB.Model(Link{}).Where("slug = ?", model.Slug).Count(&count); count > 0 {
		return errors.Wrapf(errorsext.ErrConflict, "there is a link with slug '%s' is already exist", model.Slug)
	}
	if err := repo.DB.Create(model).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo linkRepository) ReturnBySlug(slug string) (*Link, error) {
	model := new(Link)
	if result := repo.DB.Where("slug = ?", slug).Find(model); result.RecordNotFound() {
		return nil, errors.Wrapf(errorsext.ErrNotFound, "there is no link with slug '%s' exists", slug)
	} else if result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}
	return model, nil
}
