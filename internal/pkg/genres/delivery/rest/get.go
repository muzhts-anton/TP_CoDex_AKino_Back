package gendelivery

import (
	"codex/internal/pkg/domain"

	"net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

func (handler *GenresHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	genre := mux.Vars(r)["genre"]
	if genre == "" {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	movs, err := handler.GenresUsecase.GetMovies(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(movs)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *GenresHandler) GetGenres(w http.ResponseWriter, r *http.Request) {

	genres, err := handler.GenresUsecase.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
