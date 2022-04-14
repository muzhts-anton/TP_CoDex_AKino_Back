package usrdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/log"
	"codex/internal/pkg/utils/filesaver"
	
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

const (
	root = "."
	path = "/static/avatars/"
)

func (handler *UserHandler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 * 1024 * 1024) // лимит 10МБ
	if err != nil {
		log.Warn(fmt.Sprintf("parse multipart form error: %s", err))
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	uploaded, header, err := r.FormFile("avatar")
	if err != nil {
		log.Warn(fmt.Sprintf("error while parsing file: %s", err))
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}
	defer uploaded.Close()

	filename, err := filesaver.UploadFile(uploaded, root, path, filepath.Ext(header.Filename))
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	us, err := handler.UserUsecase.UpdateAvatar(clientID, filename)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.AlreadyIn.Error(), http.StatusBadRequest)
		return
	}

	marshalledUs, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(marshalledUs)
}
