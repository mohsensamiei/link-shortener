package services

import (
	"context"
	"github.com/mohsensamiei/link-shortener/api"
	"github.com/mohsensamiei/link-shortener/internal/app/statistics/models"
	"github.com/mohsensamiei/link-shortener/internal/pkg/auth"
	"github.com/mohsensamiei/link-shortener/internal/pkg/env"
	uuid "github.com/satori/go.uuid"
	"time"
)

func NewVisit(ctx context.Context, configs env.Configs) api.VisitServer {
	return &visit{
		Service: auth.Service{
			Context: ctx,
		},
		Configs: configs,
	}
}

type visit struct {
	auth.Service
	Configs env.Configs
}

func (s visit) VisitRepository() models.VisitRepository {
	return s.Context.Value(models.VisitRepositoryInstance{}).(models.VisitRepository)
}

func (s visit) Register(_ context.Context, req *api.VisitRegisterRequest) (*api.VoidResponse, error) {
	model := &models.Visit{
		IsMobile: req.IsMobile,
		Browser:  req.Browser,
		Url:      req.Url,
	}
	if userID, err := uuid.FromString(req.UserId); err == nil {
		model.UserID = &userID
	}
	if err := s.VisitRepository().Create(model); err != nil {
		return nil, err
	}
	return &api.VoidResponse{}, nil
}

func (s visit) ReportByUrl(context.Context, *api.VoidRequest) (*api.VisitReportByUrlResponse, error) {
	res := new(api.VisitReportByUrlResponse)
	res.Today = new(api.VisitReportByUrlResponse_Report)
	res.LastDay = new(api.VisitReportByUrlResponse_Report)
	res.LastWeek = new(api.VisitReportByUrlResponse_Report)
	res.LastMonth = new(api.VisitReportByUrlResponse_Report)

	now := time.Now()

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	today, err := s.VisitRepository().ReportByUrl(midnight, now)
	if err != nil {
		return nil, err
	}
	for _, item := range today {
		if item.Browser == nil && item.IsMobile == nil {
			res.Today.Total = append(res.Today.Total, &api.VisitReportByUrlResponse_Total{
				Url:   item.Url,
				Count: item.Count,
			})
		} else if item.Browser != nil {
			res.Today.ByBrowser = append(res.Today.ByBrowser, &api.VisitReportByUrlResponse_ByBrowser{
				Url:     item.Url,
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.Today.ByMobile = append(res.Today.ByMobile, &api.VisitReportByUrlResponse_ByMobile{
				Url:      item.Url,
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastDay, err := s.VisitRepository().ReportByUrl(now.Add(-1*time.Hour*24), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastDay {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastDay.Total = append(res.LastDay.Total, &api.VisitReportByUrlResponse_Total{
				Url:   item.Url,
				Count: item.Count,
			})
		} else if item.Browser != nil {
			res.LastDay.ByBrowser = append(res.LastDay.ByBrowser, &api.VisitReportByUrlResponse_ByBrowser{
				Url:     item.Url,
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastDay.ByMobile = append(res.LastDay.ByMobile, &api.VisitReportByUrlResponse_ByMobile{
				Url:      item.Url,
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastWeek, err := s.VisitRepository().ReportByUrl(now.Add(-1*time.Hour*24*7), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastWeek {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastWeek.Total = append(res.LastWeek.Total, &api.VisitReportByUrlResponse_Total{
				Url:   item.Url,
				Count: item.Count,
			})
		} else if item.Browser != nil {
			res.LastWeek.ByBrowser = append(res.LastWeek.ByBrowser, &api.VisitReportByUrlResponse_ByBrowser{
				Url:     item.Url,
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastWeek.ByMobile = append(res.LastWeek.ByMobile, &api.VisitReportByUrlResponse_ByMobile{
				Url:      item.Url,
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastMonth, err := s.VisitRepository().ReportByUrl(now.Add(-1*time.Hour*24*30), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastMonth {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastMonth.Total = append(res.LastMonth.Total, &api.VisitReportByUrlResponse_Total{
				Url:   item.Url,
				Count: item.Count,
			})
		} else if item.Browser != nil {
			res.LastMonth.ByBrowser = append(res.LastMonth.ByBrowser, &api.VisitReportByUrlResponse_ByBrowser{
				Url:     item.Url,
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastMonth.ByMobile = append(res.LastMonth.ByMobile, &api.VisitReportByUrlResponse_ByMobile{
				Url:      item.Url,
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	return res, nil
}

func (s visit) ReportByUser(context.Context, *api.VoidRequest) (*api.VisitReportByUserResponse, error) {
	res := new(api.VisitReportByUserResponse)
	res.Today = new(api.VisitReportByUserResponse_Report)
	res.LastDay = new(api.VisitReportByUserResponse_Report)
	res.LastWeek = new(api.VisitReportByUserResponse_Report)
	res.LastMonth = new(api.VisitReportByUserResponse_Report)

	uuidToString := func(id *uuid.UUID) string {
		if id == nil {
			return ""
		}
		return id.String()
	}
	now := time.Now()

	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	today, err := s.VisitRepository().ReportByUser(midnight, now)
	if err != nil {
		return nil, err
	}
	for _, item := range today {
		if item.Browser == nil && item.IsMobile == nil {
			res.Today.Total = append(res.Today.Total, &api.VisitReportByUserResponse_Total{
				Url:    item.Url,
				UserId: uuidToString(item.UserID),
				Count:  item.Count,
			})
		} else if item.Browser != nil {
			res.Today.ByBrowser = append(res.Today.ByBrowser, &api.VisitReportByUserResponse_ByBrowser{
				Url:     item.Url,
				UserId:  uuidToString(item.UserID),
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.Today.ByMobile = append(res.Today.ByMobile, &api.VisitReportByUserResponse_ByMobile{
				Url:      item.Url,
				UserId:   uuidToString(item.UserID),
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastDay, err := s.VisitRepository().ReportByUser(now.Add(-1*time.Hour*24), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastDay {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastDay.Total = append(res.LastDay.Total, &api.VisitReportByUserResponse_Total{
				Url:    item.Url,
				UserId: uuidToString(item.UserID),
				Count:  item.Count,
			})
		} else if item.Browser != nil {
			res.LastDay.ByBrowser = append(res.LastDay.ByBrowser, &api.VisitReportByUserResponse_ByBrowser{
				Url:     item.Url,
				UserId:  uuidToString(item.UserID),
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastDay.ByMobile = append(res.LastDay.ByMobile, &api.VisitReportByUserResponse_ByMobile{
				Url:      item.Url,
				UserId:   uuidToString(item.UserID),
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastWeek, err := s.VisitRepository().ReportByUser(now.Add(-1*time.Hour*24*7), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastWeek {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastWeek.Total = append(res.LastWeek.Total, &api.VisitReportByUserResponse_Total{
				Url:    item.Url,
				UserId: uuidToString(item.UserID),
				Count:  item.Count,
			})
		} else if item.Browser != nil {
			res.LastWeek.ByBrowser = append(res.LastWeek.ByBrowser, &api.VisitReportByUserResponse_ByBrowser{
				Url:     item.Url,
				UserId:  uuidToString(item.UserID),
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastWeek.ByMobile = append(res.LastWeek.ByMobile, &api.VisitReportByUserResponse_ByMobile{
				Url:      item.Url,
				UserId:   uuidToString(item.UserID),
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	lastMonth, err := s.VisitRepository().ReportByUser(now.Add(-1*time.Hour*24*30), now)
	if err != nil {
		return nil, err
	}
	for _, item := range lastMonth {
		if item.Browser == nil && item.IsMobile == nil {
			res.LastMonth.Total = append(res.LastMonth.Total, &api.VisitReportByUserResponse_Total{
				Url:    item.Url,
				UserId: uuidToString(item.UserID),
				Count:  item.Count,
			})
		} else if item.Browser != nil {
			res.LastMonth.ByBrowser = append(res.LastMonth.ByBrowser, &api.VisitReportByUserResponse_ByBrowser{
				Url:     item.Url,
				UserId:  uuidToString(item.UserID),
				Browser: *item.Browser,
				Count:   item.Count,
			})
		} else if item.IsMobile != nil {
			res.LastMonth.ByMobile = append(res.LastMonth.ByMobile, &api.VisitReportByUserResponse_ByMobile{
				Url:      item.Url,
				UserId:   uuidToString(item.UserID),
				IsMobile: *item.IsMobile,
				Count:    item.Count,
			})
		}
	}

	return res, nil
}
