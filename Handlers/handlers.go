package handlers

import (
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yoyo"))
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
}

func ProfilePage(w http.ResponseWriter, r *http.Request) {
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
}
