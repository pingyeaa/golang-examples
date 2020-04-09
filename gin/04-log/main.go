package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	gin.DebugPrintRouteFunc()

	gin.DefaultWriter = io.MultiWriter(f)
	r.GET("hello", func(c *gin.Context) {
		log.Println(123)
		c.JSON(200, "ok")
	})
	r.Run()
}
