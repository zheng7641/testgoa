package controller

import "github.com/gin-gonic/gin"

func Test1(c *gin.Context) {
	c.JSON(200, "a")
}
