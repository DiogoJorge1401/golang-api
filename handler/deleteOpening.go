package handler

import (
	"fmt"
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Router /openings/{id} [delete]
// @Summary Delete Opening
// @Description Delete job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
func DeleteOpeningHandler(ctx *gin.Context) {
	opening := &schemas.Opening{}

	id := ctx.Param("id")

	if err := db.First(opening, id).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusNotFound, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

	if err := db.Delete(opening).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Error when deleting opening with id: %v", id))
		return
	}

	SendJSONResponse(ctx, http.StatusNoContent, gin.H{})

}
