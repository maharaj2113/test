package router

import (
	"github.com/gorilla/mux"
	"github.com/maharaj2113/test/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAll).Methods("GET")
	router.HandleFunc("api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleOne).Methods("DELETE")
	router.HandleFunc("/api/deleAll", controller.DeleAll).Methods("DELETE")
	return router
}
