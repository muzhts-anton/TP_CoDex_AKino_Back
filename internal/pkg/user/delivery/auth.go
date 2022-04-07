package usrdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/sessions"

	"encoding/json"
	"net/http"
	"strconv"
)

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(domain.User)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	us, err := handler.UserUsecase.Register(*userForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, us.Id)
	if err != nil && us.Id != 0 {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(domain.UserBasic)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	us, err := handler.UserUsecase.Login(*userForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = sessions.CheckSession(r)
	if err != domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.AlreadyIn.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, us.Id)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	id, err := sessions.CheckSession(r)
	if err == domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusForbidden)
		return
	}

	err = sessions.FinishSession(w, r, id)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	type authResp struct {
		Status string `json:"status"`
		Id     string `json:"id,omitempty"`
	}

	userId, err := sessions.CheckSession(r)
	if err == domain.Err.ErrObj.UserNotLoggedIn {
		out, err := json.Marshal(authResp{Status: strconv.Itoa(http.StatusBadRequest)})
		if err != nil {
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	us, err := handler.UserUsecase.CheckAuth(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(authResp{Status: strconv.Itoa(http.StatusOK), Id: strconv.FormatUint(us.Id, 10)})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
