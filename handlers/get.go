package handlers

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/mira-miracoli/brewess2/models"
	"github.com/mira-miracoli/brewess2/utils"
)

func GetAllHops(context *gin.Context) {
	hop := new([]models.Hop)
	result := models.DB.Find(&hop)
	if utils.CheckResults(context, result) {
		context.JSON(http.StatusOK, &hop)
	}
}

func DisplayResourceByID(context *gin.Context) {
	res := utils.GetResourceInterface(context.Params.ByName("resource"))
	if reflect.TypeOf(res) != reflect.TypeOf(fmt.Errorf("")) {
		result := models.DB.Find(&res, context.Params.ByName("id"))
		if utils.CheckResults(context, result) {
			context.JSON(http.StatusOK, &result)
		}

	} else {
		context.JSON(http.StatusNotFound, "The resource "+context.Params.ByName("resource")+" does not exist!")
	}
}
