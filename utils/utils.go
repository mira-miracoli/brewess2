package utils

import (
	"fmt"

	"github.com/mira-miracoli/brewess2/models"
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