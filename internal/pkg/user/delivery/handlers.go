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

	router.HandleFunc("/user/signup", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login", handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/logout", handler.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/authcheck", handler.CheckAuth).Methods("GET", "OPTIONS")

	router.HandleFunc("/user/{id:[0-9]+}", handler.GetBasicInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/bookmarks/{id:[0-9]+}", handler.GetBookmarks).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/update/{id:[0-9]+}", handler.UpdateInfo).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/reviews/{id:[0-9]+}", handler.GetUserReviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/update/avatar/{id:[0-9]+}", handler.UploadAvatar).Methods("POST", "OPTIONS")

	router.HandleFunc("/csrf", handler.GetCsrf).Methods("GET", "OPTIONS")
}
