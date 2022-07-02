package middlewares

import (
	"github.com/gorilla/csrf"
)

var CsrfMdlw = csrf.Protect(
	[]byte("32-byte-long-auth-key"),
	csrf.Path("/"),
	csrf.Secure(false),
)
