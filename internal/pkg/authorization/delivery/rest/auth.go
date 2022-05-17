package autdelivery

import (
	"codex/internal/pkg/authorization/delivery/grpc"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/sessions"
	"codex/internal/pkg/utils/sanitizer"

	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id             uint64 `json:"ID"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email"`
	Imgsrc         string `json:"imgsrc"`
	RepeatPassword string `json:"repeatpassword,omitempty"`
}

func (handler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(domain.User)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	sanitizer.SanitizeUser(userForm)

	us, err := handler.AuthClient.Register(context.Background(), &grpc.User{
		ID:             userForm.Id,
		Username:       userForm.Username,
		Password:       userForm.Password,
		Email:          userForm.Email,
		Imgsrc:         userForm.Imgsrc,
		RepeatPassword: userForm.RepeatPassword,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, us.ID)
	if err != nil && us.ID != 0 {
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

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(domain.UserBasic)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	sanitizer.SanitizeUserBasic(userForm)

	us, err := handler.AuthClient.Login(context.Background(), &grpc.UserBasic{
		Email:    userForm.Email,
		Password: userForm.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = sessions.CheckSession(r); err != domain.Err.ErrObj.UserNotLoggedIn {
		http.Error(w, domain.Err.ErrObj.AlreadyIn.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, us.ID)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
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

func (handler *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	type authResp struct {
		Status string `json:"status"`
		Id     string `json:"ID,omitempty"`
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

	out, err := json.Marshal(authResp{
		Status: strconv.Itoa(http.StatusOK),
		Id:     strconv.FormatUint(userId, 10),
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
