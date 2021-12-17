// 官方 Demo
package main

import (
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func main() {

	init_controller(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
