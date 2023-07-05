package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-json-web-token/controllers"
	"github.com/mhdianrush/go-json-web-token/middleware"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	// will be check the user
	router.Use(middleware.Auth)
	
	router.HandleFunc("/me", controllers.Me).Methods(http.MethodGet)
}
