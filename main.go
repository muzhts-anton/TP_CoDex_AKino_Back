package main

import (
	"codex/Authorization"
	"codex/Collections"
	"codex/Actors"
	"codex/Movies"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://park-akino.ru") // url to deployed front
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token, Location")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization, X-CSRF-Token")
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.Use(CorsMiddleware)
	api.HandleFunc("/signup", authorization.Register).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", authorization.Login).Methods("POST", "OPTIONS")
	api.HandleFunc("/logout", authorization.Logout).Methods("POST", "OPTIONS")
	api.HandleFunc("/checkAuth", authorization.CheckAuth).Methods("GET", "OPTIONS")

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
	fmt.Println("connecting to port ", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
