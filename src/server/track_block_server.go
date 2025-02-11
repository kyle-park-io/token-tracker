package server

import (
	"fmt"
	"net/http"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/router"
	"github.com/kyle-park-io/token-tracker/wss"

	docs "github.com/kyle-park-io/token-tracker/docs/transfertracker"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title Blockchain Transfer Tracker API
// @version 1.0
// @description API for tracking token transfers in blockchain
// @host localhost:8081
// @BasePath /api/v1
func StartTransferTrackerServer() {
	r := gin.Default()

	env := viper.GetString("ENV")
	if env == "" {
		logger.Log.Errorln("check env config")
	}
	root_path := viper.GetString("ROOT_PATH")
	url_prefix := viper.GetString(fmt.Sprintf("server.tracker.%s.url_prefix", env))
	api_url_prefix := viper.GetString(fmt.Sprintf("server.tracker.%s.api_url_prefix", env))

	// Group for Base URL
	base := r.Group(url_prefix)
	{
		base.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hi! i'm token tracker.",
			})
		})

		logs := base.Group("logs")
		{
			logs.GET("/tracker", func(c *gin.Context) {
				c.HTML(http.StatusOK, "logs_tracker.html", nil) // Render logs HTML
			})
		}

		// base.GET("/ws", ws.HandleWebSocket)
		base.GET("/ws", wss.HandleWebSocket)
	}

	docs.SwaggerInfo.BasePath = api_url_prefix
	// Group for API versioning
	v1 := r.Group(api_url_prefix)
	{
		get := v1.Group("/get")
		{
			// block number
			get.GET("/getLatestBlockNumber", router.GetLatestBlockNumber) // Returns JSON structure

			// block
			get.GET("/getRandomBlock", router.GetRandomBlock) // Returns JSON structure
			get.GET("/getBlock", router.GetBlock)             // Returns JSON structure

			// block timeStamp
			get.GET("/getBlockTimestamp", router.GetBlockTimestamp) // Returns JSON structure

			// block position
			get.GET("/getBlockPosition", router.GetBlockPosition) // Returns JSON structure

			// balance
			get.GET("/getETHBalance", router.GetETHBalance)     // Returns JSON structure
			get.GET("/getERC20Balance", router.GetERC20Balance) // Returns JSON structure
		}

		track := v1.Group("/track")
		{
			// track Balance
			// track.GET("/trackETH", router.TrackETH)     // Returns JSON structure
			track.GET("/trackETH", router.TrackETHBatch)   // Returns JSON structure
			track.GET("/trackETH2", router.TrackETHBatch2) // Returns JSON structure
			track.GET("/trackERC20", router.TrackERC20)    // Returns JSON structure
		}

		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", router.Helloworld)
		}
	}

	switch env {
	case "dev":
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	case "prod":
		r.GET("/tracker/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.LoadHTMLGlob(fmt.Sprintf("%s/html/%s/*.html", root_path, env))
	logger.Log.Infoln("Starting transfer-tracker server on :8081")
	r.Run(":" + viper.GetString(fmt.Sprintf("server.tracker.%s.port", env))) // listen and serve on 0.0.0.0:8081
}
