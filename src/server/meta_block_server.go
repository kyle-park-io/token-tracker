package server

import (
	"fmt"
	"net/http"

	"github.com/kyle-park-io/token-tracker/executor"
	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/router"
	"github.com/kyle-park-io/token-tracker/ws"

	docs "github.com/kyle-park-io/token-tracker/docs/transfertracker"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title Blockchain Timestamp API
// @version 1.0
// @description API for collecting and managing blockchain block timestamps
// @host localhost:8080
// @BasePath /api/v1
func StartBlockTimestampServer() {
	r := gin.Default()

	env := viper.GetString("ENV")
	if env == "" {
		logger.Log.Errorln("check env config")
	}
	root_path := viper.GetString("ROOT_PATH")
	url_prefix := viper.GetString(fmt.Sprintf("server.recorder.%s.url_prefix", env))
	api_url_prefix := viper.GetString(fmt.Sprintf("server.recorder.%s.api_url_prefix", env))

	// Group for Base URL
	base := r.Group(url_prefix)
	{
		base.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi! i'm block timestamp recoder.",
			})
		})
		base.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		logs := base.Group("logs")
		{
			logs.GET("/timestamp", func(c *gin.Context) {
				c.HTML(http.StatusOK, "logs_timestamp.html", nil) // Render logs HTML
			})
		}

		base.GET("/ws", ws.HandleWebSocket)
	}

	docs.SwaggerInfo.BasePath = api_url_prefix
	// Group for API versioning
	v1 := r.Group(api_url_prefix)
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", router.Helloworld)
		}
	}

	switch env {
	case "dev":
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	case "prod":
		r.GET("/recorder/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.LoadHTMLGlob(fmt.Sprintf("%s/html/%s/*.html", root_path, env))
	logger.Log.Infoln("Starting block-timestamp server on :8080")

	go executor.EnhancedBlockTimestampRecorder()
	r.Run() // listen and serve on 0.0.0.0:8080
}
