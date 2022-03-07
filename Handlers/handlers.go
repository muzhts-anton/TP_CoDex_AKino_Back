package handlers

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var minPasswordLength = 6
var maxPasswordLength = 14
var minUsernameLength = 4
var maxUsernameLength = 14

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type MyHandler struct {
	sessions map[string]uint
	users    map[string]*User
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint, 10),
		users: map[string]*User{
			"rvasily": {2, "rvasily", "loveee", "123@gmail.com"},
		},
	}
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func isValidLogin(login string) error {
	login = strings.Trim(login, " ")

	if len(login) < minUsernameLength {
		return errors.New("Too short username " + login + strconv.Itoa(len(login)))
	}
	if len(login) > maxUsernameLength {
		return errors.New("Too long username ")
	}
	for _, symbol := range login {
		if !(unicode.IsLetter(symbol) || unicode.IsDigit(symbol) || symbol != []rune("_")[0]) {
			return errors.New("Contains illegal characters")
		}
	}
	return nil
}

func isValidPassword(password string) error {
	password = strings.Trim(password, " ")

	if len(password) < minPasswordLength {
		return errors.New("Too short password")
	}
	if len(password) > maxUsernameLength {
		return errors.New("Too long password")
	}

	for _, symbol := range password {
		if !(unicode.IsLetter(symbol) || unicode.IsDigit(symbol) || symbol != []rune("_")[0]) {
			return errors.New("Contains illegal characters")
		}
	}
	return nil
}

func isValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("Incorrect email")
	}
	// TODO: Check on uniq

	return nil
}

// {description: "Топ 256", imgSrc: "top.png", page: "movies", number: "1"}

type JSONCHIK struct {
	description string `json:"description"`
	imgSrc      string `json:"imgSrc"`
	page        string `json:"page"`
	number      string `json:"number"`
}

type Collection struct {
	coll []JSONCHIK
}

func (api *MyHandler) MainPage(w http.ResponseWriter, r *http.Request) {
	// authorized := false
	// session, err := r.Cookie("session_id")
	// if err == nil && session != nil {
	// 	_, authorized = api.sessions[session.Value]
	// }

	// if authorized {
	// 	w.Write([]byte("autrorized"))
	// } else {
	// 	w.Write([]byte("not autrorized"))
	// }

	jsonchik := make([]JSONCHIK, 1)
	jsonchik[0] = JSONCHIK{"Top 256", "top.png", "movies", "1"}
	coll := Collection{jsonchik}
	b, err := json.Marshal(coll)
	if err != nil {
		http.Error(w, "lolkek", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (api *MyHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	user, ok := api.users[r.FormValue("login")]
	if !ok {
		http.Error(w, `no user`, 404)
		return
	}

	if user.Password != r.FormValue("password") {
		http.Error(w, `bad pass`, 400)
		return
	}

	if err := isValidLogin(user.Username); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if err := isValidPassword(user.Password); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("Successful login"))
}

func (api *MyHandler) SignupPage(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	email := r.FormValue("email")

	if err := isValidLogin(login); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if _, ok := api.users[login]; ok {
		http.Error(w, "Already exist", 404)
		return
	}

	if err := isValidPassword(password); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	if err := isValidEmail(email); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	api.users[login] = &User{ID: 1, Username: login, Password: password, Email: email}
	user := api.users[login]

	SID := RandStringRunes(32)

	api.sessions[SID] = user.ID

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("Successful signup"))
}

func (api *MyHandler) ProfilePage(w http.ResponseWriter, r *http.Request) {
}

func (api *MyHandler) LogoutPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Error(w, `no sess on client`, 401)
		return
	}

	if _, ok := api.sessions[session.Value]; !ok {
		http.Error(w, `no sess on server`, 401)
		return
	}

	delete(api.sessions, session.Value)
	w.Write([]byte("In logout"))
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
}
