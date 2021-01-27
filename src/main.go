package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import . "./router"

func main() {
	//gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	InitRouter(engine)

	s := &http.Server{
		Addr:           ":9090",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
