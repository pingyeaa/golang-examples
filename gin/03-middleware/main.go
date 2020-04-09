package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.Use(PingYe(), Where())
	r.GET("/know", PingYe(), Where())
	r.GET("/unknown", func(c *gin.Context) {
		c.String(200, "???")
	})

	v1 := r.Group("v1")
	v1.Use(PingYe(), Where())
	{
		v1.GET("/know", func(c *gin.Context) {
			c.String(200, "know")
		})
		v1.GET("/unknown", func(c *gin.Context) {
			c.String(200, "unknown")
		})
	}
	r.Run()
}

func PingYe() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Next()
		c.String(200, "平也最帅")
		//c.Abort()
	}
}

func Where() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "在全宇宙")
	}
}
