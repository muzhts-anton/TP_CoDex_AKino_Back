package coldelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type CollectionsHandler struct {
	CollectionsUsecase domain.CollectionsUsecase
}

func SetColHandlers(router *mux.Router, uc domain.CollectionsUsecase) {
	handler := &CollectionsHandler{
		CollectionsUsecase: uc,
	}

	router.HandleFunc(getFeedUrl, handler.GetFeed).Methods("GET", "OPTIONS")
	router.HandleFunc(getCollectionUrl, handler.GetCollection).Methods("GET", "OPTIONS")
}
