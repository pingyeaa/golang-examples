package main

import "github.com/gin-gonic/gin"

func main() {
	//路由分组
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.String(200, "v1/login")
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.String(200, "v1/submit")
		})
	}

	v2 := r.Group("v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.String(200, "v2/login")
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.String(200, "v2/submit")
		})
	}
	r.Run()
}
