package ratdelivery

import (
	"codex/internal/pkg/rating/delivery/grpc"

	"github.com/gorilla/mux"
)

type RatingHandler struct {
	RatingUsecase grpc.PosterClient
}

func SetRatHandlers(router *mux.Router, pc grpc.PosterClient) {
	handler := &RatingHandler{
		RatingUsecase: pc,
	}

	router.HandleFunc(postRatingUrl, handler.PostRating).Methods("POST", "OPTIONS")
}