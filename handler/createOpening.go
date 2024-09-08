package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	SendJSONResponse(ctx, http.StatusCreated, gin.H{"data": opening, "message": "opening created successfully!"})
}
