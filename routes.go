package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(config Config, engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "We got Gin")
	})

	engine.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})
}
