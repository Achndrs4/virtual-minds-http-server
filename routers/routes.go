package routers

import (
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "virtualminds/http-server/docs"

	"virtualminds/http-server/config"
	"virtualminds/http-server/database"
	"virtualminds/http-server/handlers"
	"virtualminds/http-server/services"

	"github.com/gin-gonic/gin"
)

// @title			Virtual Minds HTTP-Server
// @version		1.0
// @contact.name	Ani Chandrashekhar
// @contact.email	ani.chandrashekhar@gmail.com
// @description	A small API written in Go using Gin for routing and GORM as an ORM to a database
// @host			localhost:8080
// @BasePath		/api/v1
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
			handlers.DailyStats(context, statsService)
		})

	}
	docs := router.Group("/docs")
	{
		docs.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}
