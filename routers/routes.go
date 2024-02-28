package routers

import (
	"log"

	"github.com/gin-gonic/gin"

	"virtualminds/http-server/config"
	"virtualminds/http-server/database"
	"virtualminds/http-server/handlers"
	"virtualminds/http-server/services"
)

func SetupRoutes(logger *log.Logger, database database.DatabaseInterface) *gin.Engine {
	// set gin settings
	router := gin.Default()
	router.SetTrustedProxies(config.GetTrustedProxies())
	gin.DefaultWriter = logger.Writer()
	router.Use(gin.Logger())

	customerService := &services.CustomerService{DB: database}
	statsService := &services.StatsService{DB: database}

	v1 := router.Group("/api/v1")
	{
		// our post endpoint for registering customer events
		v1.POST("/customer/", func(context *gin.Context) {
			handlers.PersistCustomerEntry(context, customerService)
		})

		// our get endpoint to get statistics
		v1.GET("/statistics/", func(context *gin.Context) {
			handlers.GetDailyStats(context, statsService)
		})
	}

	return router
}
