package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mira-miracoli/brewess2/models"
)

func GetAllHops(context *gin.Context){
	hop := new([]models.Hop)
	results := models.DB.Find(&hop)
	context.JSON(http.StatusOK, &results)
}