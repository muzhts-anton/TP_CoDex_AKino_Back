package usrdelivery

import (
	"codex/internal/pkg/domain"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (handler *UserHandler) GetBasicInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	us, err := handler.UserUsecase.GetBasicInfo(userId)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusNotFound)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *UserHandler) GetBookmarks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	bookmarks, err := handler.UserUsecase.GetBookmarks(userId)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusNotFound)
		return
	}

	out, err := json.Marshal(domain.BookmarkResp{
		UserId:    userId,
		Bookmarks: bookmarks,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *UserHandler) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	newUsrInfo := new(domain.UpdUser)
	err := json.NewDecoder(r.Body).Decode(&newUsrInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	usr, err := handler.UserUsecase.UpdateUser(userId, *newUsrInfo)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(usr)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
