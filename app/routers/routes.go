package routers

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"virtualminds/http-server/handlers"
)

//	@title			Virtual Minds Coding Challenge
//	@version		1.0
//	@description	A performant server with a master-slave databases running on gin

//	@contact.name	Ani Chandrashekhar
//	@contact.email	ani.chandrashekhar@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//
//	@BasePath	/api/v1
func SetupRoutes(logger *log.Logger) *gin.Engine {
	router := gin.Default()

	gin.DefaultWriter = logger.Writer()
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")
	{
		// our post endpoint for registering customer events
		v1.POST("/customer/", handlers.CreateCustomerEntry)

		// our get endpoint to get statistics
		v1.GET("/statistics/:customer/:day", handlers.GetDailyStats)

		// our swagger endpoint
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
