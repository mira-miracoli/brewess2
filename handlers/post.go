package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mira-miracoli/brewess2/models"
)

func CreateHop(context *gin.Context){
	var input models.Hop
	if err := context.ShouldBindJSON(&input); err != nil {
 	   context.AbortWithError(http.StatusBadRequest, err)
  	  return
  	}
	  // Create book
	book := models.Hop{Amount: input.Amount, Iso: input.Iso, Name: input.Name}
	models.DB.Create(&book)

	  context.JSON(http.StatusOK, gin.H{"data": book})
}