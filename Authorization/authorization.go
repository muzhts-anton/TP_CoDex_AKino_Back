package authorization

import (
	"codex/DB"
	"codex/sessions"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/securecookie"
	"time"
)

type userForLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	db                  DB.UserMockDatabase
	errorBadInput       = "error - bad input"
	errorAlreadyIn      = "error - already in"
	errorBadCredentials = "error - bad credentials"
	errorInternalServer = "Internal server error"
)

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
	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	var hashKey = []byte("very-secret")
	var blockKey = []byte("a-lot-secret")
	var s = securecookie.New(hashKey, blockKey)
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("session_id_register", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session_id_register",
			Value: encoded,
			Path:  "/",
			Secure: true,
			HttpOnly: true,
			Expires: time.Now().Add(10 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}

	defer r.Body.Close()
	userForm := new(DB.User)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}

	if userForm.Email == "" || userForm.Username == "" || userForm.Password == "" || userForm.RepeatPassword == "" {
		http.Error(w, errorBadInput, http.StatusBadRequest)
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

	idReg := db.AddUser(userForm)
	err = sessions.StartSession(w, r, userForm.ID)
	if err != nil && idReg != 0 {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	userForm.OmitPassword()
	b, err := json.Marshal(userForm)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
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
	user.OmitPassword()
	b, err := json.Marshal(user)
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
	_, err = w.Write(b)
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

func CheckAuth(w http.ResponseWriter, r *http.Request) {
	userID, err := sessions.CheckSession(r)
	if err == sessions.ErrUserNotLoggedIn {
		http.Error(w, errorBadInput, http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo := DB.User{ID: userID}
	b, err := json.Marshal(userInfo)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
}
