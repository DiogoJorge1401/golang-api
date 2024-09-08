package handler

import (
	"errors"
	"fmt"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	opening := &schemas.Opening{}

	id := ctx.Param("id")

	if result := db.First(opening, id); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		SendErrorJSONResponse(ctx, 404, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

	db.Delete(opening)

	SendJSONResponse(ctx, 204, gin.H{})

}
