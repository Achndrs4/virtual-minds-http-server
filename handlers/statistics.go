package handlers

import (
	"net/http"

	"virtualminds/http-server/services"

	"github.com/gin-gonic/gin"
)

func GetDailyStats(context *gin.Context, statsService *services.StatsService) {
	customer := context.Query("customer")
	dateString := context.Query("day")

	// check if the customer_id field exists before executing a query.
	if customer == "" {
		handleError(context, http.StatusBadRequest)
		return
	}

	// Get daily stats from the service
	stats, err := statsService.GetDailyStats(customer, dateString)
	if err != nil {
		handleError(context, http.StatusInternalServerError)
		return
	}

	// Respond with the daily stats
	context.JSON(http.StatusOK, gin.H{
		"Customer Valid Requests":   stats.CustomerValidRequests,
		"Customer Invalid Requests": stats.CustomerInvalidRequests,
		"Total Daily Requests":      stats.TotalDailyRequests,
	})
}
