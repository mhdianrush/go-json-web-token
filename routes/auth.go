package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-json-web-token/controllers"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
}
