package router

import "github.com/gin-gonic/gin"
import . "../controller"

func InitRouter(engine *gin.Engine) *gin.Engine {

	engine.GET("/test", Test1)

	engine.POST("/testpost", TestPost)

	engine.POST("/call", Call)

	return engine
}
