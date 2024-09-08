package handler

import (
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusInternalServerError, "error when fetching openings")
	}

	SendJSONResponse(ctx, http.StatusOK, gin.H{"data": openings})
}
