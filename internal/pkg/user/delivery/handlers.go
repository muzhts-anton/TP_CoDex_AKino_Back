package usrdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func SetUsrHandlers(router *mux.Router, uc domain.UserUsecase) {
	handler := &UserHandler{uc}

	router.HandleFunc("/user/{id:[0-9]+}", handler.GetBasicInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/signup", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login", handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/logout", handler.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/authcheck", handler.CheckAuth).Methods("GET", "OPTIONS")
}