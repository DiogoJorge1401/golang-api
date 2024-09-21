package handler

import (
	"fmt"
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Router /openings/{id} [get]
// @Summary Opening Details
// @Description get job opening details
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} OpeningJSONReponse
// @Failure 404 {object} ErrorResponse
func FindOpeningHandler(ctx *gin.Context) {
	opening := &schemas.Opening{}
	id := ctx.Param("id")

	if err := db.First(opening, id).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusNotFound, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

	SendJSONResponse(ctx, http.StatusOK, opening)
}
