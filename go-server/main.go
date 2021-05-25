package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ip, _ := c.RemoteIP()
		c.String(200, fmt.Sprintf("request comes from: %v and RemoteIP is: %v\n", c.Request.RemoteAddr, ip))
	})

	r.Run("0.0.0.0:443")
}