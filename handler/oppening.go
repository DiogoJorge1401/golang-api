package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "POST Oppening"})
}
func ListOppeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GET Oppening"})
}
func FindOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "GET Oppening id"})
}
func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "DELETE Oppening"})
}
func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "PUT Oppening"})
}
