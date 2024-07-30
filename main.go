package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	server.Run(":8080")
}
