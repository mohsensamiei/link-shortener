package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/mohsensamiei/link-shortener/pkg/errorsext"
	"github.com/mohsensamiei/link-shortener/pkg/gormext"
	"github.com/pkg/errors"
)

type User struct {
	gormext.Model
	Email          string `gorm:"not null;size:100;unique" validate:"required,email,max=100"`
	Username       string `gorm:"not null;size:25;unique" validate:"required,min=3,max=25"`
	PasswordHashed string `gorm:"not null;size:256" validate:"required,max=256"`
}

type UserRepositoryInstance struct{}
type UserRepository interface {
	Create(model *User) error
	ReturnByEmailOrUsername(email, username string) (*User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB:       db,
		Validate: validator.New(),
	}
}

type userRepository struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func (repo userRepository) Create(model *User) error {
	if err := repo.Validate.Struct(model); err != nil {
		return errors.Wrap(errorsext.ErrValidation, err.Error())
	}

	var count int
	if repo.DB.Model(User{}).Where("Email = ? or Username = ?", model.Email, model.Username).Count(&count); count > 0 {
		return errors.Wrapf(errorsext.ErrConflict,
			"there is an account with email \"%v\" or username \"%v\" is already exist", model.Email, model.Username)
	}

	if err := repo.DB.Create(model).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo userRepository) ReturnByEmailOrUsername(email, username string) (*User, error) {
	user := new(User)
	if repo.DB.Where("Email = ? or Username = ?", email, username).Find(user).RecordNotFound() {
		return nil, errors.Wrapf(errorsext.ErrNotFound,
			"can not find user with email \"%v\" or username \"%v\"", email, username)
	}
	return user, nil
}
