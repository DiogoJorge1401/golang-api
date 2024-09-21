package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Router /openings [post]
// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpening true "Reques body"
// @Success 201 {object} OpeningJSONReponse
// @Failure 400 {object} ErrorResponse
func CreateOpeningHandler(ctx *gin.Context) {
	request := new(CreateOpening)

	ctx.BindJSON(request)

	if err := request.Validate(); err != nil {
		SendErrorJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := request.GetOpening()

	if err := db.Create(opening).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	SendJSONResponse(ctx, http.StatusCreated, opening)
}
