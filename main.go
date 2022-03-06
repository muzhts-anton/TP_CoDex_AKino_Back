package main

import (
	"codex/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://xenodochial-mayer-d916ec.netlify.app/") // url to deployed front
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token, Location")
		w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
		w.Header().Set("Access-Control-Max-Age", "600")
		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := handlers.NewMyHandler()
	router := mux.NewRouter()
	api := router.PathPrefix("").Subrouter()
	
	api.Use(CorsMiddleware)

	api.HandleFunc("/", handler.MainPage).Methods("GET", "OPTIONS")
	api.HandleFunc("/profile", handler.ProfilePage).Methods("GET", "OPTIONS")
	api.HandleFunc("/signup", handler.SignupPage).Methods("POST", "OPTIONS")
	api.HandleFunc("/login", handler.LoginPage).Methods("POST", "OPTIONS")
	api.HandleFunc("/logout", handler.LogoutPage).Methods("GET", "OPTIONS")

	port := os.Getenv("PORT") // to get port from Heroku
	if port == "" {
		port = "3000"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("connecting to port " + port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
