package route

import "github.com/gin-gonic/gin"

func GET(ctx *gin.Context) {
	ctx.String(200, "GET")
}
