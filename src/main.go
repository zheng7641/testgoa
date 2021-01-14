package main

import (
	c1 "./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", c1.Test1)
	_ = router.Run("localhost:9090")
}
