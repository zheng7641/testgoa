package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"code":200,
			"success":true,
		})
	})
	_ = router.Run("localhost:9090")
}
