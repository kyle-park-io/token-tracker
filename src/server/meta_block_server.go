package server

import (
	"net/http"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/router"

	docs "github.com/kyle-park-io/token-tracker/docs/transfertracker"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Blockchain Timestamp API
// @version 1.0
// @description API for collecting and managing blockchain block timestamps
// @host localhost:8080
// @BasePath /api/v1
func StartBlockTimestampServer() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/api/v1"
	// Group for API versioning
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", router.Helloworld)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Log.Infoln("Starting web server on :8080")
	r.Run() // listen and serve on 0.0.0.0:8080
}