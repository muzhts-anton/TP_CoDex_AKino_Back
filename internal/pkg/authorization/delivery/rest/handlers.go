package autdelivery

import (
	"codex/internal/pkg/authorization/delivery/grpc"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthClient grpc.AutherClient
}

func SetAutHandlers(router *mux.Router, uc grpc.AutherClient) {
	handler := &AuthHandler{
		uc,
	}

	router.HandleFunc(signupUrl, handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc(loginUrl, handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc(logoutUrl, handler.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc(authcheckUrl, handler.CheckAuth).Methods("GET", "OPTIONS")
}
