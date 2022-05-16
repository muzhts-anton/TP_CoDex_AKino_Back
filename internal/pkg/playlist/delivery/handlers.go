package pladelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type PlaylistHandler struct {
	PlaylistUsecase domain.PlaylistUsecase
}

func SetPlaHandlers(router *mux.Router, pu domain.PlaylistUsecase) {
	handler := &PlaylistHandler{
		PlaylistUsecase: pu,
	}
	router.HandleFunc(createPlaylistUrl, handler.CreatePlaylist).Methods("POST", "OPTIONS")
	router.HandleFunc(addMovieUrl, handler.AddMovie).Methods("POST", "OPTIONS")
	router.HandleFunc(deleteMovieUrl, handler.DeleteMovie).Methods("POST", "OPTIONS")
	router.HandleFunc(deletePlaylistUrl, handler.DeletePlaylist).Methods("POST", "OPTIONS")
	router.HandleFunc(alterPlaylistPublic, handler.AlterPlaylistPublic).Methods("POST", "OPTIONS")
}
