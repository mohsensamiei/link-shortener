package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func AutoMigrate(db *gorm.DB) error {
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return errors.Wrap(err, "can not create extension uuid-ossp")
	}
	if err := db.AutoMigrate(Link{}).Error; err != nil {
		return errors.Wrap(err, "can not create table links")
	}
	return nil
}
