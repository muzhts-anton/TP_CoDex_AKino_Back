package sessions

import (
	"errors"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var ErrUserNotLoggedIn = errors.New("user not logged in")
var errUint64Cast = errors.New("id uint64 cast error")

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))
var sessionName = "session-name"

var hashKey = []byte("very-secret")
var blockKey = []byte("a-lot-secret")
var s = securecookie.New(hashKey, blockKey)

// func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
// 	value := map[string]string{
// 		"foo": "bar",
// 	}
// 	if encoded, err := s.Encode("cookie-name", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:  "cookie-name",
// 			Value: encoded,
// 			Path:  "/",
// 			Secure: true,
// 			HttpOnly: true,
// 		}
// 		http.SetCookie(w, cookie)
// 	}
// }

func StartSession(w http.ResponseWriter, r *http.Request, id uint64) error {

	session, _ := store.Get(r, sessionName)
	session.Values["id"] = id
	session.Options = &sessions.Options{
		MaxAge:   100000,
		Secure:   true,
		HttpOnly: true,
		// SameSite: http.SameSiteNoneMode,
		Path:     "/",
	}
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("session_id", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session_id",
			Value: encoded,
			Path:  "/",
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
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
		session.Options.MaxAge = -1
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
