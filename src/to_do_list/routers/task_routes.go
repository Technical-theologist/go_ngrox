package routers

import (
	"to-do-list/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/createtask", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/changestatus", controllers.ChangeTaskStatus).Methods("PUT")
	router.HandleFunc("/api/deletetask", controllers.DeleteTask).Methods("DELETE")
	return router
}
