package routers

import (
	"mydatabase/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/employee", controllers.CreateOneEmployee).Methods("POST")
	router.HandleFunc("/api/employee/{id}", controllers.GetOneEmployee).Methods("GET")
	router.HandleFunc("/api/employees", controllers.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/employee/{id}", controllers.DeleteOneEmployee).Methods("DELETE")
	router.HandleFunc("/api/employee/update/{id}", controllers.UpdateOneEmployee).Methods("PUT")
	return router
}
