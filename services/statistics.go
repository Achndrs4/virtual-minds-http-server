package services

import (
	"strconv"
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"
	"virtualminds/http-server/utils"
)

type StatsServiceInterface interface {
	GetDailyStats(customer string, dateString string) (*models.DailyStats, error)
}

type StatsService struct {
	DB database.DatabaseInterface
}

func (s *StatsService) GetDailyStats(customer string, dateString string) (*models.DailyStats, error) {
	datetime, err := utils.ParseDateString(dateString)
	if err != nil {
		return nil, err
	}
	customer_int, err := strconv.Atoi(customer)
	if err != nil {
		return nil, err
	}

	dailyTotal, err := s.DB.GetDailyTotal(datetime)
	if err != nil {
		return nil, err
	}
	print(dailyTotal)
	customerStat, err := s.DB.GetCustomerTotal(datetime, customer_int)
	if err != nil {
		return nil, err
	}

	return &models.DailyStats{
		CustomerValidRequests:   customerStat.TotalCount,
		CustomerInvalidRequests: customerStat.TotalInvalidCount,
		TotalDailyRequests:      dailyTotal,
	}, nil
}
