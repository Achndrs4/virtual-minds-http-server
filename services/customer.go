package services

import (
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"
	"virtualminds/http-server/utils"
)

type CustomerServiceInterface interface {
	IsCustomerValid(customerID uint) (bool, error)
	IsIPValid(ip_str string) (bool, error)
	IsUserAgentValid(userAgent string) (bool, error)
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

func (s *CustomerService) IsIPValid(ip_str string) (bool, error) {
	// Check if the IP is valid
	ip_integer, err := utils.GetValidIP(ip_str)
	if err != nil {
		return false, err
	}

	return s.DB.IsIPValid(ip_integer)
}

func (s *CustomerService) IsUserAgentValid(userAgent string) (bool, error) {
	// check if the user agent is banned
	return s.DB.IsUserAgentValid(userAgent)
}
