package server

import (
	"net/http"

	"github.com/kyle-park-io/token-tracker/logger"
	"github.com/kyle-park-io/token-tracker/router"
	"github.com/spf13/viper"

	docs "github.com/kyle-park-io/token-tracker/docs/transfertracker"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Blockchain Transfer Tracker API
// @version 1.0
// @description API for tracking token transfers in blockchain
// @host localhost:9090
// @BasePath /api/v1
func StartTransferTrackerServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi! i'm token tracker.",
		})
	})

	docs.SwaggerInfo.BasePath = viper.GetString("server.api_prefix")
	// Group for API versioning
	v1 := r.Group(viper.GetString("server.api_prefix"))
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
		}

		track := v1.Group("/track")
		{
			// track Balance
			track.GET("/trackETH", router.TrackETH)     // Returns JSON structure
			track.GET("/trackERC20", router.TrackERC20) // Returns JSON structure
		}

		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", router.Helloworld)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Log.Infoln("Starting alternative server on :9090")
	r.Run(":9090") // listen and serve on 0.0.0.0:9090
}
