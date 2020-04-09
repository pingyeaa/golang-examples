package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.BasicAuth(gin.Accounts{
		"name": "pingye",
	}))
	r.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(200, user+"已登录成功")
	})
	r.Run()
}
