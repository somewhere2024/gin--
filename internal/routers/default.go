package routers

import (
	"gin--/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

var R *gin.Engine

func DefaultRouter() {
	R = gin.New()
	R.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	// 默认路由
	R.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "test"})
	})

	R.Run(":8000")
}
