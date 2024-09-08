package handler

import (
	"fmt"
	"net/http"

	"github.com/DiogoJorge1401/golang-api/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.First(&schemas.Opening{}, id).Error; err != nil {
		SendErrorJSONResponse(ctx, http.StatusNotFound, fmt.Sprintf("Opening not found with id: %v", id))
		return
	}

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

	SendJSONResponse(ctx, http.StatusCreated, gin.H{"data": opening, "message": "opening created successfully!"})
}
