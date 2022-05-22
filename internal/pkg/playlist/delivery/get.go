package pladelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/sessions"
	"codex/internal/pkg/utils/sanitizer"

	"encoding/json"
	"net/http"

	"github.com/mailru/easyjson"
)

func (handler *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}

	playlistRequest := new(domain.PlaylistRequest)
	err := json.NewDecoder(r.Body).Decode(&playlistRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	sanitizer.SanitizePlaylistCreating(playlistRequest)

	us, err := handler.PlaylistUsecase.CreatePlaylist(*playlistRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *PlaylistHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}
	addPlaylistInfo := new(domain.MovieInPlaylist)
	err := json.NewDecoder(r.Body).Decode(&addPlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.AddMovie(*addPlaylistInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *PlaylistHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}
	MovieInPlaylist := new(domain.MovieInPlaylist)
	err := json.NewDecoder(r.Body).Decode(&MovieInPlaylist)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.DeleteMovie(*MovieInPlaylist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistHandler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}
	deletePlaylistInfo := new(domain.DeletePlaylistInfo)
	err := json.NewDecoder(r.Body).Decode(&deletePlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.DeletePlaylist(*deletePlaylistInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistHandler) AlterPlaylistPublic(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}

	alterPlaylistPublicInfo := new(domain.AlterPlaylistPublicInfo)
	err := json.NewDecoder(r.Body).Decode(&alterPlaylistPublicInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.AlterPlaylistPublic(*alterPlaylistPublicInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
