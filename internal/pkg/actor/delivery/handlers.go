package actdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type ActorHandler struct {
	ActorUsecase domain.ActorUsecase
}

func SetActHandlers(router *mux.Router, uc domain.ActorUsecase) {
	handler := &ActorHandler{
		ActorUsecase: uc,
	}

	router.HandleFunc("/actors/{id:[0-9]+}", handler.GetActor).Methods("GET", "OPTIONS")
}