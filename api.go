package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		username := c.DefaultQuery("username", "flyingpot")
		info, error := fetch(username)
		if error != nil {
			c.AbortWithStatus(http.StatusBadGateway)
			return
		}
		c.String(200, render(info))
	})

	r.Run()
}
