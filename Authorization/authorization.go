package authorization

import (
	"codex/DB"
	"codex/sessions"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"codex/Collections"
)

type userForLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const userNotLoggedIn = "User not logged in"
const cantMarshal = "cant marshal"

var db DB.UserMockDatabase

const (
	errorBadInput       = "error - bad input"
	errorAlreadyIn      = "error - already in"
	errorBadCredentials = "error - bad credentials"
	errorInternalServer = "Internal server error"
	errorParseJSON      = "Error parse JSON"
	errorEmptyField     = "Empty field"
)

type authResponse struct {
	Status string `json:"status"`
	user DB.User
}

type userWithRepeatedPassword struct{
	user DB.User
	RepeatPassword string `json:"repeatpassword"`
}

type userWithoutPasswords struct{
	Username       string `json:"username"`
	Email          string `json:"email"`
}

func (us *userWithRepeatedPassword) OmitPassword() {
	us.user.Password = ""
	us.RepeatPassword = ""
}

func GetBasicInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	u64, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	user, err := db.FindId(u64)
	if err != nil {
		http.Error(w, errorBadInput, http.StatusNotFound)
		return
	}
	user.OmitPassword()
	userInfoJson, err := json.Marshal(user)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userInfoJson)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(userWithRepeatedPassword)
	err := json.NewDecoder(r.Body).Decode(&userForm)

	if err != nil {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}

	if userForm.user.Email == "" || userForm.user.Username == "" || userForm.user.Password == "" || userForm.RepeatPassword == "" {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	if userForm.user.Password != userForm.RepeatPassword {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	_, err = db.FindEmail(userForm.user.Email)
	if err == nil {
		http.Error(w, errorAlreadyIn, http.StatusConflict)
		return
	}

	_, err = db.FindUsername(userForm.user.Username)
	if err == nil {
		http.Error(w, errorAlreadyIn, http.StatusConflict)
		return
	}

	idReg := db.AddUser(&DB.User{ID: userForm.user.ID, Username: userForm.user.Username, Password: "", Email: userForm.user.Email})

	userOut := userWithoutPasswords{Username: userForm.user.Username, Email: userForm.user.Email}

	err = sessions.StartSession(w, r, userForm.user.ID)
	if err != nil && idReg != 0 {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	userInfoJson, err := json.Marshal(userOut)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(userInfoJson)
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm := new(userForLogin)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	if userForm.Email == "" || userForm.Password == "" {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	user, err := db.FindEmail(userForm.Email)
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userForm.Password))
	if err != nil || errPassword != nil {
		http.Error(w, errorBadCredentials, http.StatusUnauthorized)
		return
	}
	_, err = sessions.CheckSession(r)
	if err != sessions.ErrUserNotLoggedIn {
		http.Error(w, errorAlreadyIn, http.StatusBadRequest)
		return
	}
	userOut := userWithoutPasswords{Username: user.Username, Email: user.Email}

	userOutMarshalled, err := json.Marshal(userOut)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}

	err = sessions.StartSession(w, r, user.ID)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userOutMarshalled)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	id, err := sessions.CheckSession(r)
	if err == sessions.ErrUserNotLoggedIn {
		http.Error(w, errorBadInput, http.StatusForbidden)
		return
	}
	err = sessions.FinishSession(w, r, id)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(collections.Alabdsel)
	if err != nil {
		http.Error(w, cantMarshal, http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	userID, err := sessions.CheckSession(r)
	if err == sessions.ErrUserNotLoggedIn {
		tmp, err := json.Marshal(authResponse{Status: strconv.Itoa(http.StatusBadRequest)})
		if err != nil {
			http.Error(w, cantMarshal, http.StatusInternalServerError)
			return
		}
		w.Write(tmp)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := DB.User{ID: userID}
	tmp := authResponse{Status: "", user:userInfo}
	userInfoJson, err := json.Marshal(tmp)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userInfoJson)
}
