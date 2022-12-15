package main

import (
	"brewess2/routers"
	"brewess2/services"
	"brewess2/utils"
	"log"
	"net/http"
)

func main() {
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	services.InitDB()
	var appRouter = routers.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
}
