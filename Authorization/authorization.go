package authorization

import (
	"codex/DB"
	"codex/sessions"
	"encoding/json"
	"net/http"
	"strconv"

	collections "codex/Collections"
	"errors"
	"net/mail"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type userForLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db DB.UserMockDatabase
var invalidEmailError = errors.New("Invalid email")
var invalidUsernameError = errors.New("Invalid username")
var invalidPasswordError = errors.New("Invalid Password")

const (
	errorBadInput         = "error - bad input"
	errorAlreadyIn        = "error - already in"
	errorEmailNotFound    = "error - email not found"
	errorPasswordNotFound = "error - password not found"
	errorInternalServer   = "Internal server error"
	errorParseJSON        = "Error parse JSON"
	errorEmptyField       = "Empty field"
	unmatchedPasswords    = "Passwords are unmatched"
	invalidEmail          = "Invalid email"
	invalidUsername       = "Invalid username"
	invalidPassword       = "Invalid password"
	cantMarshal           = "cant marshal"
	userNotLoggedIn       = "User not logged in"
)

type authResponse struct {
	Status string `json:"status"`
	user   userWithoutPasswords
}

type userWithRepeatedPassword struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	RepeatPassword string `json:"repeatpassword"`
}

type userWithoutPasswords struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (us *userWithRepeatedPassword) OmitPassword() {
	us.Password = ""
	us.RepeatPassword = ""
}

func validEmail(address string) error {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return invalidEmailError
	}
	return nil
}

func validUsername(username string) error {
	for _, char := range username {
		if !(unicode.IsLetter(char) || unicode.Is(unicode.Cyrillic, char)) {
			return invalidUsernameError
		}
	}
	return nil
}
func validPassword(password string) error {
	if len(password) < 8 {
		return invalidPasswordError
	}

	return nil
}
func trimCredentials(email *string, username *string, password *string, repeatPassword *string) {
	*email = strings.Trim(*email, " ")
	*username = strings.Trim(*username, " ")
	*password = strings.Trim(*password, " ")
	*repeatPassword = strings.Trim(*repeatPassword, " ")
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

	trimCredentials(&userForm.Email, &userForm.Username, &userForm.Password, &userForm.RepeatPassword)

	if userForm.Email == "" || userForm.Username == "" || userForm.Password == "" || userForm.RepeatPassword == "" {
		http.Error(w, errorEmptyField, http.StatusBadRequest)
		return
	}

	if err = validEmail(userForm.Email); err != nil {
		http.Error(w, invalidEmail, http.StatusBadRequest)
		return
	}

	if err = validUsername(userForm.Username); err != nil {
		http.Error(w, invalidUsername, http.StatusBadRequest)
		return
	}

	if err = validPassword(userForm.Password); err != nil {
		http.Error(w, invalidPassword, http.StatusBadRequest)
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

	mockedResponse, _ := json.Marshal("")

	if err == sessions.ErrUserNotLoggedIn {
		http.Error(w, errorBadInput, http.StatusForbidden)
		w.Write(mockedResponse)
		return
	}
	err = sessions.FinishSession(w, r, id)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		w.Write(mockedResponse)
		return
	}

	w.Write(mockedResponse)
	w.WriteHeader(http.StatusOK)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(collections.DBFilms)
	if err != nil {
		http.Error(w, cantMarshal, http.StatusInternalServerError)
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
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
	userOutput := userWithoutPasswords{Email: userInfo.Email, Username: userInfo.Username}
	tmp := authResponse{Status: strconv.Itoa(http.StatusOK), user: userOutput}
	userInfoJson, err := json.Marshal(tmp)
	if err != nil {
		http.Error(w, errorInternalServer, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(userInfoJson)
}
