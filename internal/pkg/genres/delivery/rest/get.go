package gendelivery

import (
	"codex/internal/pkg/domain"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

//easyjson:json
type Genres []domain.Genre

func (handler *GenresHandler) GetGenre(w http.ResponseWriter, r *http.Request) {
	genre := mux.Vars(r)["genre"]
	if genre == "" {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	genreWithMovies, err := handler.GenresUsecase.GetGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(genreWithMovies)
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

	out, err := easyjson.Marshal(Genres(genres))
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
