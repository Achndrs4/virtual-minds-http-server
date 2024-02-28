package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var HTTPErrors = map[int]string{
	http.StatusBadRequest:          "Required data not received or formatted incorrectly. Check swagger documentation at /v1/api/swagger",
	http.StatusNotFound:            "The Customer could not be found or is inactive",
	http.StatusForbidden:           "The requested user has been blocked",
	http.StatusInternalServerError: "Server Error",
}

func handleError(context *gin.Context, statusCode int) {
	context.JSON(statusCode, gin.H{"Error": HTTPErrors[statusCode]})
}
