package services

import (
	"strconv"
	"time"
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"
	"virtualminds/http-server/utils"
)

type StatsServiceInterface interface {
	GetDailyStats(customer string, dateString string) (*models.DailyStats, error)
	ValidateRequest(customer string, dateString string) (int, *time.Time, error)
}

type StatsService struct {
	DB database.DatabaseInterface
}

func (s *StatsService) GetDailyStats(customer int, date *time.Time) (*models.DailyStats, error) {

	dailyTotal, err := s.DB.GetDailyTotal(*date)
	if err != nil {
		return nil, err
	}

	customerStat, err := s.DB.GetCustomerTotal(*date, customer)
	if err != nil {
		return nil, err
	}

	return &models.DailyStats{
		CustomerValidRequests:   customerStat.TotalCount,
		CustomerInvalidRequests: customerStat.TotalInvalidCount,
		TotalDailyRequests:      dailyTotal,
	}, nil
}

func (s *StatsService) ValidateRequest(customer string, dateString string) (int, *time.Time, error) {
	customer_int, err := strconv.Atoi(customer)
	if err != nil {
		return -1, nil, err
	}
	date, err := utils.ParseDateString(dateString)
	if err != nil {
		return -1, nil, err
	}
	return customer_int, date, nil
}
