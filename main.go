package main

import (
	"codex/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := handlers.NewMyHandler()
	router := mux.NewRouter()
	router.HandleFunc("/profile", handler.ProfilePage)
	router.HandleFunc("/signup", handler.SignupPage)
	router.HandleFunc("/login", handler.LoginPage)
	router.HandleFunc("/logout", handler.LogoutPage)
	router.HandleFunc("/", handler.MainPage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("connecting to port " + port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
