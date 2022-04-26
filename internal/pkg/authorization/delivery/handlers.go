package autdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func SetAutHandlers(router *mux.Router, uc domain.AuthUsecase) {
	handler := &AuthHandler{
		uc,
	}

	router.HandleFunc(signupUrl, handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc(loginUrl, handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc(logoutUrl, handler.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc(authcheckUrl, handler.CheckAuth).Methods("GET", "OPTIONS")
}
