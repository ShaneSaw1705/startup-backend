package main

import (
	"startup-backend/packages/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	testRoute := server.Group("/test")
	{
		testRoute.GET("/ping", route.GET)
	}
	server.Run(":8080")
}
