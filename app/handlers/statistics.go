package handlers

import (
	"net/http"
	"time"
	"virtualminds/http-server/database"
	"virtualminds/http-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary		Get statistics based on day and customer
// @Description	Get statistics based on Day and customer
// @ID				get-stats
// @Produce		json
// @Param			customer	query	int		true	"Customer ID"
// @Param			day			query	string	true	"Day"
// @Failure		404			"Customer data not found"
// @Failure		400			"Bad Request"
// @Failure		500			"Server Error"
// @Success		200			"Success"
// @Router			/statistics/ [post]
func GetDailyStats(context *gin.Context) {
	customer := context.Query("customer")
	dateString := context.Query("day")

	// check if the customer_id field exists before executing a query
	if customer == "" {
		context.JSON(http.StatusBadRequest, ("No CustomerID provided"))
		return
	}

	// next, make sure that the date is correct
	datetime, err := parseDateString(dateString)
	if err != nil {
		context.JSON(http.StatusBadRequest, ("Valid datetime should be passed in YYYYMMDD format"))
		return
	}
	db := database.GetReaderDB()

	// get the daily total overall (SUM)
	daily_total, err := getDailyTotal(db, datetime)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	// get the daily total overall for a user. Sending 2 seperate queries is (almost always) cheaper and faster than doing a partial query
	customer_total, customer_failed, err := getCustomerTotal(db, datetime, customer)
	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	if customer_total == 0 {
		context.Status(http.StatusNotFound)
		return
	}

	context.JSON(http.StatusOK, (gin.H{"Valid Requests": customer_total, "Invalid Requests": customer_failed, "Daily Requests": daily_total}))

}

func getDailyTotal(db *gorm.DB, datetime time.Time) (int, error) {
	var daily_total int
	daily_result := db.Model(models.HourlyStat{}).
		Select("SUM(RequestCount)").
		Where("time >= ? and time <= ? and customer_id = ?", datetime.Add(-24*time.Hour), datetime).
		Scan(&daily_total)
	if err := daily_result.Error; err != nil {
		return -1, err
	} else {
		return daily_total, nil
	}
}

func getCustomerTotal(db *gorm.DB, datetime time.Time, customer string) (int, int, error) {
	var customer_total, customer_failed int
	customer_result := db.Model(models.HourlyStat{}).
		Select("SUM(RequestCount) as total, SUM(InvalidCount) as failed").
		Where("time >= ? and time <= ? and customer_id = ?", datetime.Add(-24*time.Hour), datetime, customer).
		Group("key").
		Scan(&struct {
			Total  int
			Failed int
		}{customer_total, customer_failed})
	if err := customer_result.Error; err != nil {
		return -1, -1, err
	} else {
		return customer_total, customer_failed, nil
	}
}

func parseDateString(dateString string) (time.Time, error) {
	layout := "20060102"
	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}
