package services

import (
	"encoding/json"
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"
	"virtualminds/http-server/utils"

	"gorm.io/gorm"
)

type CustomerServiceInterface interface {
	IsCustomerValid(customerID uint) (bool, error)
	IsIPValid(ip_str string) (bool, error)
	IsUserAgentBanned(userAgent string) (bool, error)
	GetRequestBody(body []byte) (*models.CustomerRequest, error)
}

type CustomerService struct {
	DB database.DatabaseInterface
}

func (s *CustomerService) IsCustomerValid(customerID uint) (bool, error) {
	// Check if the Customer is valid/exists
	customer, err := s.DB.GetCustomer(customerID)
	if err == gorm.ErrRecordNotFound || !customer.Active {
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

	return s.DB.IsIPBanned(ip_integer)
}

func (s *CustomerService) IsUserAgentBanned(userAgent string) (bool, error) {
	return s.DB.IsUserAgentBanned(userAgent)
}
func (s *CustomerService) GetRequestBody(body []byte) (*models.CustomerRequest, error) {
	// check if the json can be parsed, and if it has the required fields
	var request models.CustomerRequest
	if err := json.Unmarshal(body, &request); err != nil {
		return nil, err
	}
	return &request, nil
}
