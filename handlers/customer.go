package handlers

import (
	"io"
	"net/http"
	"virtualminds/http-server/services"
	"virtualminds/http-server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary		Ingests and persists a customer record
// @Description	A POST endpoint that takes in a JSON and produces a record in a statistics table if successful
// @Accept			json
// @Produce		json
// @Param			user-agent	header	string					false	"User-Agent header for user identification"
// @Param			data		body	models.CustomerRequest	true	"Customer Request body"
// @Success		200
// @Failure		400
// @Failure		401
// @Failure		500
// @Router			/customer [post]
func PersistCustomerEntry(context *gin.Context, service *services.CustomerService) {
	// validate if the body data can be formatted as json
	body, err := io.ReadAll(context.Request.Body)

	if err != nil {
		handleCustomerErr(context, http.StatusBadRequest)
	}

	// validate if the fields are valid. If we cannot get a valid customer, we do not need to write stats because stats records rely on a foreign key customer to exist
	request, err := utils.GetRequestBody(body)
	if err != nil {
		handleCustomerErr(context, http.StatusBadRequest)
		return
	}

	// check if the Customer is valid
	isCustomerValid, err := service.IsCustomerValid(request.CustomerID)
	if err != nil || !isCustomerValid {
		service.WriteCustomerStatistic(request, false)
		handleCustomerErr(context, http.StatusNotFound)
		return
	}

	// check if the IP is valid
	if isIPValid := utils.IsValidIP(request.RemoteIP); !isIPValid {
		service.WriteCustomerStatistic(request, false)
		handleCustomerErr(context, http.StatusBadRequest)
		return
	}

	// check if the IP is banned
	isIPBanned, err := service.IsIPBanned(request.RemoteIP)
	if err != nil || isIPBanned {
		service.WriteCustomerStatistic(request, false)
		handleCustomerErr(context, http.StatusForbidden)
		return
	}

	// check if the User-Agent is banned
	isUseragentBanned, err := service.IsUserAgentBanned(context.Request.UserAgent())
	if err != nil || isUseragentBanned {
		service.WriteCustomerStatistic(request, false)
		handleCustomerErr(context, http.StatusForbidden)
		return
	}

	// we try to write our success case
	err = service.WriteCustomerStatistic(request, true)
	if err != nil {
		handleCustomerErr(context, http.StatusInternalServerError)
		return
	}
	context.Status(http.StatusOK)
}
