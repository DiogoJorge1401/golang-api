package handler

import "github.com/gin-gonic/gin"

func SendErrorJSONResponse(ctx *gin.Context, code int, message string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{"error": message})
}

func SendJSONResponse(ctx *gin.Context, code int, message map[string]any) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, message)
}
