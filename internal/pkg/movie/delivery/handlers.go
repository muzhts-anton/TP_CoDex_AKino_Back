package movdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	MovieUsecase domain.MovieUsecase
}

func SetMovHandlers(router *mux.Router, uc domain.MovieUsecase) {
	handler := &MovieHandler{
		MovieUsecase: uc,
	}

	router.HandleFunc("/movies/postrating", handler.PostRating).Methods("GET", "OPTIONS")
	router.HandleFunc("/movies/postcomment", handler.PostComment).Methods("POST", "OPTIONS")
	router.HandleFunc("/movies/{id:[0-9]+}", handler.GetMovie).Methods("GET", "OPTIONS")
}