package routers

import (
	"my_mango/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateOneMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovies", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}
