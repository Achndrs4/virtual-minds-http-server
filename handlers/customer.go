package handlers

import (
	"io"
	"net/http"
	"virtualminds/http-server/services"
	"virtualminds/http-server/utils"

	"github.com/gin-gonic/gin"
)

func PersistCustomerEntry(context *gin.Context, service *services.CustomerService) {
	// validate if the body data can be formatted as json
	body, err := io.ReadAll(context.Request.Body)

	if err != nil {
		handleError(context, http.StatusBadRequest)
	}

	// validate if the fields are valid. If we cannot get a valid customer, we exit early because the stats record relies on a foreign key customer to exist
	request, err := utils.GetRequestBody(body)
	if err != nil {
		handleError(context, http.StatusBadRequest)
		return
	}

	// check if the Customer is valid
	isCustomerValid, err := service.IsCustomerValid(request.CustomerID)
	if err != nil || !isCustomerValid {
		service.WriteCustomerStatistic(request, false)
		handleError(context, http.StatusNotFound)
	}

	// check if the IP is valid
	isIPValid, err := service.IsIPValid(request.RemoteIP)
	if err != nil || !isIPValid {
		service.WriteCustomerStatistic(request, false)
		handleError(context, http.StatusForbidden)
	}

	// check if the User-Agent is valid
	isUserAgentValid, err := service.IsUserAgentValid(context.Request.UserAgent())
	if err != nil || !isUserAgentValid {
		service.WriteCustomerStatistic(request, false)
		handleError(context, http.StatusForbidden)
	}

	// we try to write our success case
	err = service.WriteCustomerStatistic(request, true)
	if err != nil {
		handleError(context, http.StatusInternalServerError)
	}
	context.Status(http.StatusOK)

}
