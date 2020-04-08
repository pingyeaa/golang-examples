package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("/article", func(c *gin.Context) {
		c.String(200, "article post")
	})
	router.DELETE("/article", func(c *gin.Context) {
		c.String(200, "article delete")
	})
	router.PUT("/article", func(c *gin.Context) {
		c.String(200, "article put")
	})
	router.GET("/article", func(c *gin.Context) {
		c.String(200, "article get")
	})
	router.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(200, id)
	})

	router.GET("/article/:id/*action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		c.String(200, id+" "+action)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "pingyeaa")
		lastname := c.Query("lastname")
		c.String(200, firstname+" "+lastname)
	})

	//表单提交
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//数组类型参数
	router.GET("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		c.String(200, ids["a"]+" "+ids["b"])
	})

	//文件上传
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		c.String(200, file.Filename)
		err := c.SaveUploadedFile(file, "/Users/enoch/Desktop/ab.png")
		if err != nil {
			c.String(500, err.Error())
		}
	})

	router.Run()
}
