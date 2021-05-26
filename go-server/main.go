package main

import (
	"fmt"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, fmt.Sprintf("request comes from: %v and ClientIP is: %v\n", c.Request.RemoteAddr, c.ClientIP()))
	})

	r.GET("/debug", func(c *gin.Context) {
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			c.String(500, fmt.Sprintf("error: %v\n", err))
			return
		}
		c.String(200, fmt.Sprintf("%s\n", requestDump))
	})

	r.Run("0.0.0.0:443")
}