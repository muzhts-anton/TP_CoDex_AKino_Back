package sessions

import (
	"errors"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

var ErrUserNotLoggedIn = errors.New("user not logged in")
var errUint64Cast = errors.New("id uint64 cast error")

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
var sessionName = "session-name"

func StartSession(w http.ResponseWriter, r *http.Request, id uint64) error {
	session, _ := store.Get(r, sessionName)
	session.Values["id"] = id
	session.Options = &sessions.Options{
		MaxAge:   100000,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
	}
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func FinishSession(w http.ResponseWriter, r *http.Request, id uint64) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	sessionId, isIn := session.Values["id"]
	if isIn && id == sessionId {
		session.Options.MaxAge = 0
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
	}
	return nil
}

func CheckSession(r *http.Request) (uint64, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return 0, err
	}
	id, isIn := session.Values["id"]
	if !isIn || session.IsNew {
		return 0, ErrUserNotLoggedIn
	}
	idCasted, ok := id.(uint64)
	if !ok {
		return 0, errUint64Cast
	}
	return idCasted, nil
}
