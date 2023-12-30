package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mira-miracoli/brewess2/models"
	"gorm.io/gorm"
)
func GetResourceInterface(resource_type string) interface{} {
	switch resource_type {
	case "hop":
		return new(models.Hop)
	case "malt":
		return new(models.Malt)
	case "yeast":
		return new(models.Yeast)
	default:
		return fmt.Errorf("Resource %s  not found!", resource_type)
	}
	
}

func CheckResults (context *gin.Context, result *gorm.DB) bool {
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Oops something went wrong!"})
    }
	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not found!"})
    }
	return true
	
}