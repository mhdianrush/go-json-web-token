package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mhdianrush/go-json-web-token/configs"
	"github.com/mhdianrush/go-json-web-token/routes"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.ConnectDB()

	route := mux.NewRouter()

	router := route.PathPrefix("/api").Subrouter()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Printf("failed create log file %s", err.Error())
	}
	logger.SetOutput(file)

	if err := godotenv.Load(); err != nil {
		logger.Printf("failed load env file %s", err.Error())
	}

	server := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: router,
	}
	if err = server.ListenAndServe(); err != nil {
		logger.Printf("failed connect to server")
	}

	logger.Printf("server running on port %s", os.Getenv("SERVER_PORT"))
}
