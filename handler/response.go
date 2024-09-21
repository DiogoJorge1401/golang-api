package handler

import (
	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func SendErrorJSONResponse(ctx *gin.Context, code int, message string) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{"error": message})
}

func SendJSONResponse(ctx *gin.Context, code int, content any) {
	ctx.Header("content-type", "application/json")
	ctx.JSON(code, gin.H{"data": content})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type OpeningJSONReponse struct {
	Data schemas.OpeningResponse `json:"data"`
}
type AllOpeningsJSONReponse struct {
	Data []schemas.OpeningResponse `json:"data"`
}
