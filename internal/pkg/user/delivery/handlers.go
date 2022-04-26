package usrdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func SetUsrHandlers(router *mux.Router, uc domain.UserUsecase) {
	handler := &UserHandler{
		uc,
	}

	router.HandleFunc(getInfoUrl, handler.GetBasicInfo).Methods("GET", "OPTIONS")
	router.HandleFunc(bookmarksUrl, handler.GetBookmarks).Methods("GET", "OPTIONS")
	router.HandleFunc(updateUrl, handler.UpdateInfo).Methods("POST", "OPTIONS")
	router.HandleFunc(reviewsUrl, handler.GetUserReviews).Methods("GET", "OPTIONS")
	router.HandleFunc(avatarUrl, handler.UploadAvatar).Methods("POST", "OPTIONS")
}
