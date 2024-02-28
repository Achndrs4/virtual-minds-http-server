package handlers

import (
	"net/http"

	"virtualminds/http-server/services"

	"github.com/gin-gonic/gin"
)

// @Summary		Serves statistics about a certain endpoint
// @Description	An endpoint that takes in a customerID and a date in YYYYMMDD format as query parameters and produces aggregations over the day in the table
// @Param			customer	query	string	false	"string customer"
// @Param			date		query	string	false	"string date"
// @Success		200
// @Failure		400
// @Failure		401
// @Produce		json
// @Failure		500
// @Router			/statistics [get]
func DailyStats(context *gin.Context, statsService *services.StatsService) {
	// Get query parameters from the request
	customer := context.Query("customer")
	dateString := context.Query("date")

	// Validate that those query parameters are valid
	customer_int, date, err := statsService.ValidateRequest(customer, dateString)
	if err != nil {
		handleStatisticsErr(context, http.StatusBadRequest)
		return
	}

	// Get daily statistics as described in models.DailyStats
	stats, err := statsService.GetDailyStats(customer_int, date)
	if err != nil {
		handleStatisticsErr(context, http.StatusInternalServerError)
		return
	}

	// If no error, return them as a JSON body
	context.JSON(http.StatusOK, gin.H{
		"Customer Valid Requests":   stats.CustomerValidRequests,
		"Customer Invalid Requests": stats.CustomerInvalidRequests,
		"Total Daily Requests":      stats.TotalDailyRequests,
	})
}
