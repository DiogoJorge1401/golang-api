package handler

import (
	"errors"
	"fmt"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindOpeningHandler(ctx *gin.Context) {
	opening := &schemas.Opening{}
	id := ctx.Param("id")

	if results := db.First(opening, id); errors.Is(results.Error, gorm.ErrRecordNotFound) {
		SendErrorJSONResponse(ctx, 404, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

	SendJSONResponse(ctx, 200, gin.H{"data": opening})
}
