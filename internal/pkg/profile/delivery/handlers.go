package profiledelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	ProfileUsecase domain.ProfileUsecase
}

func SetProfileHandlers(router *mux.Router, uc domain.ProfileUsecase) {
	handler := &ProfileHandler{
		ProfileUsecase: uc,
	}

	router.HandleFunc("/profile/{id:[0-9]+}", handler.getProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/profile/bookmarks/{id:[0-9]+}", handler.getBookmarks).Methods("GET", "OPTIONS")
	router.HandleFunc("/profile/reviews/{id:[0-9]+}", handler.getReviews).Methods("GET", "OPTIONS")
}