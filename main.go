package main

import (
	"codex/Handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handler := handlers.NewMyHandler()
	router := mux.NewRouter()
	router.HandleFunc("/profile", handler.ProfilePage)
	router.HandleFunc("/signup", handler.SignupPage)
	router.HandleFunc("/login", handler.LoginPage)
	router.HandleFunc("/logout", handler.LogoutPage)
	router.HandleFunc("/", handler.MainPage)

	fmt.Println("connecting to port 8000")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
