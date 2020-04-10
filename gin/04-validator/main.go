package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" binding:"required,alphanum"`
	Password string `form:"password" binding:"required,min=6,max=12"`
}

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var login Login
		err := c.Bind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, login)
	})
	router.Run()
}
