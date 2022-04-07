package app

import (
	"codex/Actors"
	"codex/Authorization"
	"codex/Collections"
	"codex/Movies"

	"codex/internal/pkg/middlewares"
	"codex/internal/pkg/utils/log"

	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func RunServer() {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(middlewares.Cors)

	api.HandleFunc("/signup", authorization.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", authorization.Login).Methods("POST", "OPTIONS")
	api.HandleFunc("/logout", authorization.Logout).Methods("POST", "OPTIONS")
	api.HandleFunc("/checkAuth", authorization.CheckAuth).Methods("GET", "OPTIONS")
	api.HandleFunc("/user/{id:[0-9]+}", authorization.GetUser).Methods("GET", "OPTIONS")

	api.HandleFunc("/collections/collection/{id:[0-9]+}", collections.GetCol).Methods("GET", "OPTIONS")
	api.HandleFunc("/collections", collections.GetCollections).Methods("GET", "OPTIONS")

	api.HandleFunc("/movies/{id:[0-9]+}", movies.GetMovie).Methods("GET", "OPTIONS")
	api.HandleFunc("/actors/{id:[0-9]+}", actors.GetActor).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT") // to get port from Heroku
	if port == "" {
		port = "3000"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}

	log.Info("connecting to port " + port)

	if err := server.ListenAndServe(); err != nil {
		log.Error(err)
	}
}
