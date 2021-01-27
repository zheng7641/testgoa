package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type testjson struct {
	Name string `json:"name" binding:"required"`
}

func Test1(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func TestPost(c *gin.Context) {
	var t testjson
	err := c.ShouldBindJSON(&t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//panic("foo")
	c.String(http.StatusOK, t.Name)
}

func Call(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// Note that msg.Name becomes "user" in the JSON
	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	c.SecureJSON(http.StatusOK, msg)
}
