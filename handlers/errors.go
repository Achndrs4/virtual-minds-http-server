package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var CustomerHTTPErrors = map[int]string{
	http.StatusBadRequest:          "Required data not received or formatted incorrectly, reqires User-Agent(Header) STR, CustomerID uint (Body) , Timestamp int64 (Body), and RemoteIP string (Body) ",
	http.StatusNotFound:            "The Customer could not be found or is inactive",
	http.StatusForbidden:           "The requested User Agent or IP has been blocked",
	http.StatusInternalServerError: "Server Error",
}

var StatisticsHTTPErrors = map[int]string{
	http.StatusBadRequest:          "Required data not received or formatted incorrectly, requires query params: INT:customer and STR:date in the format YYYYMMDD",
	http.StatusNotFound:            "The Customer record could not be found or is inactive",
	http.StatusInternalServerError: "Server Error",
}

func handleCustomerErr(context *gin.Context, statusCode int) {
	context.JSON(statusCode, gin.H{"Error": CustomerHTTPErrors[statusCode]})
}

func handleStatisticsErr(context *gin.Context, statusCode int) {
	context.JSON(statusCode, gin.H{"Error": StatisticsHTTPErrors[statusCode]})
}
