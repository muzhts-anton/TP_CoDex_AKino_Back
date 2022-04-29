package anndelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type AnnouncedHandler struct {
	AnnouncedUsecase domain.AnnouncedUsecase
}

func SetAnnHandlers(router *mux.Router, au domain.AnnouncedUsecase) {
	handler := &AnnouncedHandler{
		AnnouncedUsecase: au,
	}

	router.HandleFunc(getMoviesUrl, handler.GetMovies).Methods("GET", "OPTIONS")
}
