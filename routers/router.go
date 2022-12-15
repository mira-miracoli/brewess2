// Router for the app
package routers

import (
	"brewess2/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/malts", services.GetAllMalts).Methods("GET")
	router.HandleFunc("/malts", services.CreateMalt).Methods("POST")
	router.HandleFunc("/malts/{id}", services.GetMalt).Methods("GET")
	router.HandleFunc("/malts/{id}", services.UpdateMalt).Methods("PUT")
	router.HandleFunc("/malts/{id}", services.DeleteMalt).Methods("DELETE")
	router.HandleFunc("/hops", services.GetAllHops).Methods("GET")
	router.HandleFunc("/hops", services.CreateHop).Methods("POST")
	router.HandleFunc("/hops/{id}", services.GetHop).Methods("GET")
	router.HandleFunc("/hops/{id}", services.UpdateHop).Methods("PUT")
	router.HandleFunc("/hops/{id}", services.DeleteHop).Methods("DELETE")
	router.HandleFunc("/yeasts", services.GetAllYeasts).Methods("GET")
	router.HandleFunc("/yeasts", services.CreateYeast).Methods("POST")
	router.HandleFunc("/yeasts/{id}", services.GetYeast).Methods("GET")
	router.HandleFunc("/yeasts/{id}", services.UpdateYeast).Methods("PUT")
	router.HandleFunc("/yeasts/{id}", services.DeleteYeast).Methods("DELETE")
	return router
}
