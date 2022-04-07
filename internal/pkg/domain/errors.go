package domain

import (
	"errors"
)

type errMsg struct {
	tmp string
}

type errObj struct {
	UserNotLoggedIn error
	Uint64Cast      error
	FinishSession   error

	NoUser         error
	InternalServer error

	InvalidEmail    error
	InvalidUsername error
	InvalidPassword error

	EmptyField         error
	UnmatchedPasswords error
	BadPassword        error

	BadInput error

	AlreadyIn error
}

type err struct {
	ErrMsg errMsg
	ErrObj errObj
}

var Err = err{
	errMsg{
		tmp: "lolkek",
	},
	errObj{
		UserNotLoggedIn: errors.New("User not logged in"),
		Uint64Cast:      errors.New("Id uint64 cast error"),
		FinishSession:   errors.New("Passed through if on FinishSession"),

		NoUser:         errors.New("No user found"),
		InternalServer: errors.New("Internal server"),

		InvalidEmail:    errors.New("Invalid email"),
		InvalidUsername: errors.New("Invalid username"),
		InvalidPassword: errors.New("Invalid password"),

		EmptyField:         errors.New("Empty field"),
		UnmatchedPasswords: errors.New("Unmatched passwords"),
		BadPassword:        errors.New("Wrong password"),

		BadInput: errors.New("Bad input"),

		AlreadyIn: errors.New("User is already logged in"),
	},
}
