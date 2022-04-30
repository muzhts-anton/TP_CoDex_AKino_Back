package gendelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type GenresHandler struct {
	GenresUsecase domain.GenresUsecase
}

func SetGenHandlers(router *mux.Router, gu domain.GenresUsecase) {
	handler := &GenresHandler{
		GenresUsecase: gu,
	}

	router.HandleFunc(getMoviesUrl, handler.GetGenre).Methods("GET", "OPTIONS")
	router.HandleFunc(getGenresUrl, handler.GetGenres).Methods("GET", "OPTIONS")
}
