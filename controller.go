package main

import (
	"github.com/gin-gonic/gin"
)

func init_controller() {

	health()

	callRemote()
}

/** 健康检查
 */
func health() {

	route.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "pong",
		})
	})
}

/**
* 测试 请求
 */

func callRemote() {
	route.GET("/test", func(c *gin.Context) {
		requestBaidu()
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "pong",
		})
	})
}
