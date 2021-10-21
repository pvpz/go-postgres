package router

import (
	"go-postgres/pkg/middleware"

	"github.com/gorilla/mux"
)

func Router(handler middleware.Handler) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/{id}", handler.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", handler.GetAllUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user", handler.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", handler.UpdateUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/user/{id}", handler.DeleteUser).Methods("DELETE", "OPTIONS")

	return router
}
