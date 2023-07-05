package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mhdianrush/go-json-web-token/configs"
	"github.com/mhdianrush/go-json-web-token/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(file)

	logger.Println("Server Running On Port 8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
