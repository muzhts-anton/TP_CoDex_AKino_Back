package ratdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type RatingHandler struct {
	RatingUsecase domain.RatingUsecase
}

func SetRatHandlers(router *mux.Router, ru domain.RatingUsecase) {
	handler := &RatingHandler{
		RatingUsecase: ru,
	}

	router.HandleFunc(postRatingUrl, handler.PostRating).Methods("POST", "OPTIONS")
}