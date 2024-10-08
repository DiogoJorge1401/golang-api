package handler

import (
	"fmt"
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Router /openings/{id} [put]
// @Summary Update Opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body UpdateOpening true "Reques body"
// @Param id path string true "Opening ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse'
// @Failure 500 {object} ErrorResponse'
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusNotFound, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

	request := new(UpdateOpening)

	ctx.BindJSON(request)

	if err := request.Validate(); err != nil {
		SendErrorJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Model(&opening).Updates(request).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusInternalServerError, "error when updating opening")
		return
	}

	SendJSONResponse(ctx, http.StatusNoContent, gin.H{})
}
