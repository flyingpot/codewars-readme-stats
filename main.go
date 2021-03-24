package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		username := c.DefaultQuery("username", "flyingpot")
		info := fetch(username)
		c.JSON(200, info)
	})

	r.Run()
}
