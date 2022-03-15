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
	errorBadInput         = "error - bad input"
	errorAlreadyIn        = "error - already in"
	errorEmailNotFound    = "error - email not found"
	errorPasswordNotFound = "error - password not found"
	errorInternalServer   = "Internal server error"
	errorParseJSON        = "Error parse JSON"
	errorEmptyField       = "Empty field"
	unmatchedPasswords    = "Passwords are unmatched"
)

type authResponse struct {
	Status string `json:"status"`
	user DB.User
}

type userWithRepeatedPassword struct{
	// user DB.User
	// ID             uint64 `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	RepeatPassword string `json:"repeatpassword"`
}

type userWithoutPasswords struct{
	Username       string `json:"username"`
	Email          string `json:"email"`
}

func (us *userWithRepeatedPassword) OmitPassword() {
	us.Password = ""
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
		http.Error(w, errorParseJSON, http.StatusBadRequest)
		return
	}

	if userForm.Email == "" || userForm.Username == "" || userForm.Password == "" || userForm.RepeatPassword == "" {
		http.Error(w, errorEmptyField, http.StatusBadRequest)
		return
	}
	if userForm.Password != userForm.RepeatPassword {
		http.Error(w, unmatchedPasswords, http.StatusBadRequest)
		return
	}

	_, err = db.FindEmail(userForm.Email)
	if err == nil {
		http.Error(w, errorAlreadyIn, http.StatusConflict)
		return
	}

	_, err = db.FindUsername(userForm.Username)
	if err == nil {
		http.Error(w, errorAlreadyIn, http.StatusConflict)
		return
	}

	idReg := db.AddUser(&DB.User{Username: userForm.Username, Password: userForm.Password, Email: userForm.Email})

	userOut := userWithoutPasswords{Username: userForm.Username, Email: userForm.Email}

	err = sessions.StartSession(w, r, idReg)
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
	if err != nil {
		http.Error(w, errorEmailNotFound, http.StatusFailedDependency)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userForm.Password))
	if err != nil {
		http.Error(w, errorPasswordNotFound, http.StatusUnauthorized)
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
		w.WriteHeader(http.StatusOK)
		w.Write(tmp)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := DB.User{ID: userID}
	tmp := authResponse{Status: strconv.Itoa(http.StatusOK), user:userInfo}
	userInfoJson, err := json.Marshal(tmp)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userInfoJson)
}
