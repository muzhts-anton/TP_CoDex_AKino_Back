package main

import (
	"codex/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profile", handlers.ProfilePage)
	router.HandleFunc("/signup", handlers.SignupPage)
	router.HandleFunc("/login", handlers.LoginPage)
	router.HandleFunc("/logout", handlers.LogoutPage)
	router.HandleFunc("/", handlers.MainPage)

	fmt.Println("connecting to port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
