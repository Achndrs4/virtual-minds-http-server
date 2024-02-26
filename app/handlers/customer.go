package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRequest struct {
	CustomerID uint   `json:"customerID" binding:"required"`
	RemoteIP   string `json:"remoteIP" binding:"required"`
	Timestamp  int64  `json:"timestamp" binding:"required"`
}

// @Summary		Create User Entry
// @Description	Get user by ID
// @ID				post-customer
// @Param			customerid	body	int		true	"Customer ID"
// @Param			useragent	header	string	true	"User Agent"
// @Param			remoteip	body	string	true	"Remote IP"
// @Param			timestamp	body	int		true	"Timestamp"
// @Success		200
// @Failure		400	"Bad Request"
// @Failure		500	"Server Error"
// @Router			/customer/ [post]
func CreateCustomerEntry(context *gin.Context) {
	// check if json can be bound to request with required fields
	var request models.CustomerRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.Status(http.StatusBadRequest)
		return
	}
	// check if the IP is banned
	isIPBanned, status := isIPBanned(request.RemoteIP)
	if isIPBanned {
		updateStatsCount(&request, false, &status)
		context.Status(status)
		return
	}

	// check if the user agent is banned
	isUserAgentBanned, status := isUserAgentBanned(context.Request.UserAgent())
	if isUserAgentBanned {
		updateStatsCount(&request, false, &status)
		context.Status(status)
		return
	}

	//check if the user is invalid
	isUserActive, status := isUserActive(request.CustomerID)
	if !isUserActive {
		updateStatsCount(&request, false, &status)
		context.Status(status)
		return
	} else {
		// our happy path
		updateStatsCount(&request, true, &status)
		context.Status(status)
	}
}

func isUserActive(customerID uint) (bool, int) {
	db := database.GetReaderDB()
	var customer models.Customer
	result := db.First(&customer, customerID)
	if result.Error == nil {
		if customer.Active {
			return true, http.StatusOK
		} else {
			return false, http.StatusForbidden
		}
	} else if result.Error == gorm.ErrRecordNotFound {
		return false, http.StatusNotFound
	} else {
		return false, http.StatusInternalServerError
	}
}

func isIPBanned(ip string) (bool, int) {
	// replace the instances of . in the ip address
	ip_int, err := strconv.Atoi(strings.ReplaceAll(ip, ".", ""))
	if err != nil {
		return true, http.StatusBadRequest
	}
	// check in the database to see if the IP exists in the block list
	db := database.GetReaderDB()
	var ipBlacklist models.IPBlacklist
	result := db.First(&ipBlacklist, ip_int)
	if result.Error == nil {
		return true, http.StatusForbidden
	} else if result.Error == gorm.ErrRecordNotFound {
		return false, http.StatusOK
	} else {
		return true, http.StatusInternalServerError
	}
}

func isUserAgentBanned(userAgent string) (bool, int) {
	// check in the database to see if the User Agent exists in the block list
	db := database.GetReaderDB()
	var uaBlacklist models.UABlacklist
	result := db.First(&uaBlacklist, userAgent)
	if result.Error == nil {
		return false, http.StatusForbidden
	} else if result.Error == gorm.ErrRecordNotFound {
		return true, http.StatusOK
	} else {
		return false, http.StatusInternalServerError
	}
}
func updateStatsCount(request *models.CustomerRequest, validated bool, status *int) {
	db := database.GetWriterDB()
	result := &models.HourlyStat{}

	// we want to save requests per hour, so we can round down the request to the hour in the timestamp
	hour := roundDownToHour(request.Timestamp)

	if err := db.FirstOrCreate(result, &models.HourlyStat{CustomerID: request.CustomerID, Time: hour}).Error; err != nil {
		*status = http.StatusInternalServerError
	}
	result.RequestCount += 1

	if !validated {
		result.InvalidCount += 1
	}

	if err := db.Save(result).Error; err != nil {
		*status = http.StatusInternalServerError
	}
}

func roundDownToHour(t int64) time.Time {
	// round requests down to the hour. 14:55:55 would be truncated as 14:00:00
	unix_time := time.Unix(t, 0)
	return time.Date(unix_time.Year(), unix_time.Month(), unix_time.Day(), unix_time.Hour(), 0, 0, 0, unix_time.Location())
}
