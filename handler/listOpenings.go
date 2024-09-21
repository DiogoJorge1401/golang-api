package handler

import (
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Router /openings [get]
// @Summary All Job Opening
// @Description get all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} AllOpeningsJSONReponse
// @Failure 500 {object} ErrorResponse
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusInternalServerError, "error when fetching openings")
	}

	SendJSONResponse(ctx, http.StatusOK, openings)
}
