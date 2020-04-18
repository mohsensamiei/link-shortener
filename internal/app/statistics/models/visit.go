package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mohsensamiei/link-shortener/pkg/gormext"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	scriptReportByUrlAndTotal    = "select visits.url, visits.user_id, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id"
	scriptReportByUrlAndMobile   = "select visits.url, visits.user_id, visits.is_mobile, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id, visits.is_mobile"
	scriptReportByUrlAndBrowser  = "select visits.url, visits.user_id, visits.browser, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id, visits.browser"
	scriptReportByUserAndTotal   = "select visits.url, visits.user_id, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id"
	scriptReportByUserAndMobile  = "select visits.url, visits.user_id, visits.is_mobile, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id, visits.is_mobile"
	scriptReportByUserAndBrowser = "select visits.url, visits.user_id, visits.browser, count(*) from visits where visits.created_at >= ? and visits.created_at < ? group by visits.url, visits.user_id, visits.browser"
)

type Visit struct {
	gormext.Model

	Url      string     `gorm:"index"`
	UserID   *uuid.UUID `gorm:"type:uuid;index"`
	IsMobile bool       `gorm:"not null;index"`
	Browser  string     `gorm:"not null;size:100;index"`
}

type Report struct {
	Url      string
	Count    int64
	UserID   *uuid.UUID
	IsMobile *bool
	Browser  *string
}

type VisitRepositoryInstance struct{}
type VisitRepository interface {
	Create(model *Visit) error
	ReportByUrl(from, to time.Time) ([]*Report, error)
	ReportByUser(from, to time.Time) ([]*Report, error)
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{
		DB: db,
	}
}

type visitRepository struct {
	DB *gorm.DB
}

func (repo visitRepository) Create(model *Visit) error {
	if err := repo.DB.Create(model).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (repo visitRepository) ReportByUrl(from, to time.Time) ([]*Report, error) {
	var rows []*Report

	var total []*Report
	if err := repo.DB.Raw(scriptReportByUrlAndTotal, from, to).Scan(&total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var mobile []*Report
	if err := repo.DB.Raw(scriptReportByUrlAndMobile, from, to).Scan(&mobile).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var browser []*Report
	if err := repo.DB.Raw(scriptReportByUrlAndBrowser, from, to).Scan(&browser).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	rows = append(rows, total...)
	rows = append(rows, mobile...)
	rows = append(rows, browser...)
	return rows, nil
}
func (repo visitRepository) ReportByUser(from, to time.Time) ([]*Report, error) {
	var rows []*Report

	var total []*Report
	if err := repo.DB.Raw(scriptReportByUserAndTotal, from, to).Scan(&total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var mobile []*Report
	if err := repo.DB.Raw(scriptReportByUserAndMobile, from, to).Scan(&mobile).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var browser []*Report
	if err := repo.DB.Raw(scriptReportByUserAndBrowser, from, to).Scan(&browser).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	rows = append(rows, total...)
	rows = append(rows, mobile...)
	rows = append(rows, browser...)
	return rows, nil
}
