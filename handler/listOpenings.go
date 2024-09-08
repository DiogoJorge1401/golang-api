package handler

import (
	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	db.Find(&openings)

	SendJSONResponse(ctx, 200, gin.H{"data": openings})
}
