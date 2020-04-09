package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(PingYe(), Where())
	r.GET("know", PingYe(), Where())
	r.GET("unknown", func(c *gin.Context) {
		c.String(200, "???")
	})
	r.Run()
}

func PingYe() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Next()
		c.String(200, "平也最帅")
		c.Abort()
	}
}

func Where() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "在全宇宙")
	}
}
