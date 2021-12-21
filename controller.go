package main

import (
	"github.com/gin-gonic/gin"
)

func init_controller(route *gin.Engine) {

	health(route)
}

/** 健康检查
 */
func health(route *gin.Engine) {

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "pong",
		})
	})
}
