package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
