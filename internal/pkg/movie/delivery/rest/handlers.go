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

	router.HandleFunc(getMovieUrl, handler.GetMovie).Methods("GET", "OPTIONS")
}