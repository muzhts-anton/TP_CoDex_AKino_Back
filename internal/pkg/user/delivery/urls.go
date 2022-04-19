package usrdelivery

const (
	signupUrl    = "/user/signup"
	loginUrl     = "/user/login"
	logoutUrl    = "/user/logout"
	authcheckUrl = "/user/authcheck"

	getInfoUrl   = "/user/{id:[0-9]+}"
	bookmarksUrl = "/user/bookmarks/{id:[0-9]+}"
	updateUrl    = "/user/update/{id:[0-9]+}"
	reviewsUrl   = "/user/reviews/{id:[0-9]+}"
	avatarUrl    = "/user/update/avatar/{id:[0-9]+}"
)
