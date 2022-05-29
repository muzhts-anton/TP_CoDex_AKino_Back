package pladelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/sessions"
	"codex/internal/pkg/utils/sanitizer"

	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/mailru/easyjson"
)

func (handler *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	playlistRequest := new(domain.PlaylistRequest)
	err = easyjson.Unmarshal(b, playlistRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
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

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	addPlaylistInfo := new(domain.MovieInPlaylist)
	err = easyjson.Unmarshal(b, addPlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
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

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieInPlaylist := new(domain.MovieInPlaylist)
	err = easyjson.Unmarshal(b, movieInPlaylist)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.PlaylistUsecase.DeleteMovie(*movieInPlaylist)
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

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	deletePlaylistInfo := new(domain.DeletePlaylistInfo)
	err = easyjson.Unmarshal(b, deletePlaylistInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
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

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	alterPlaylistPublicInfo := new(domain.AlterPlaylistPublicInfo)
	err = easyjson.Unmarshal(b, alterPlaylistPublicInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.PlaylistUsecase.AlterPlaylistPublic(*alterPlaylistPublicInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *PlaylistHandler) AlterPlaylistTitle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if _, err := sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
		return
	}

	alterPlaylistTitleInfo := new(domain.AlterPlaylistTitleInfo)
	err := json.NewDecoder(r.Body).Decode(&alterPlaylistTitleInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.PlaylistUsecase.AlterPlaylistTitle(*alterPlaylistTitleInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
