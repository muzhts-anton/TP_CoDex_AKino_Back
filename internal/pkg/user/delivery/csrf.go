package usrdelivery

import (
	"net/http"

	"github.com/gorilla/csrf"
)

func (handler UserHandler) GetCsrf(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-CSRF-Token", csrf.Token(r))
	w.WriteHeader(http.StatusNoContent)
}
