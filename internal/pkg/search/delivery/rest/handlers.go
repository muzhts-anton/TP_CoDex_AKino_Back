package serdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type SearchHandler struct {
	SearchUsecase domain.SearchUsecase
}

func SetSerHandlers(router *mux.Router, au domain.SearchUsecase) {
	handler := &SearchHandler{
		SearchUsecase: au,
	}

	router.HandleFunc(searchUrl, handler.Search).Methods("GET", "OPTIONS")
}
