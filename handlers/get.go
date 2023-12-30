package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mira-miracoli/brewess2/models"
	"github.com/mira-miracoli/brewess2/utils"
)



func GetAllHops(context *gin.Context){
	hop := new([]models.Hop)
	result := models.DB.Find(&hop)
	if utils.CheckResults(context, result) {
		context.JSON(http.StatusOK, &hop)
	}
}