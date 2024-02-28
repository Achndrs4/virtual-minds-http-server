package services

import (
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"
)

type CustomerServiceInterface interface {
	IsCustomerValid(customerID uint) (bool, error)
	IsUserAgentValid(userAgent string) (bool, error)
	IsUserAgentBanned(userAgent string) (bool, error)
	IsIPBanned(ip string) (bool, error)
}

type CustomerService struct {
	DB database.DatabaseInterface
}

func (s *CustomerService) IsCustomerValid(customerID uint) (bool, error) {
	// Check if the Customer is valid/exists
	customer, err := s.DB.GetCustomer(customerID)
	if err != nil {
		return false, err
	} else if !customer.Active {
		return false, nil
	}
	return true, nil
}

func (s *CustomerService) WriteCustomerStatistic(request *models.CustomerRequest, isValid bool) error {
	// Check if the Customer is valid/exists
	if err := s.DB.UpdateStatsCount(request, isValid); err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) IsUserAgentBanned(userAgent string) (bool, error) {
	// check if the user agent is banned
	return s.DB.IsUserAgentBanned(userAgent)
}

func (s *CustomerService) IsIPBanned(ip string) (bool, error) {
	// check if the user agent is banned
	return s.DB.IsIPBanned(ip)
}
