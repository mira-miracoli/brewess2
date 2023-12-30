package main

import (
	"github.com/mira-miracoli/brewess2/models"
	"github.com/mira-miracoli/brewess2/routers"
)

func main() {
	models.ConnectDatabase()
	routers.CreateRouter().Run()
}
