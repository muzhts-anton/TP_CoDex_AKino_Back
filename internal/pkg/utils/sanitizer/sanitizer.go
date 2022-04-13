package sanitizer

import (
	"codex/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeUser(user *domain.User) {
	sanitizer := bluemonday.UGCPolicy()
	user.Email = sanitizer.Sanitize(user.Email)
	user.Username = sanitizer.Sanitize(user.Username)
	user.Imgsrc = sanitizer.Sanitize(user.Imgsrc)
}

func SanitizeUpdUser(user *domain.UpdUser) {
	sanitizer := bluemonday.UGCPolicy()
	user.Username = sanitizer.Sanitize(user.Username)
	user.Imgsrc = sanitizer.Sanitize(user.Imgsrc)
}

func SanitizeComment(comm *string) {
	sanitizer := bluemonday.UGCPolicy()
	*comm = sanitizer.Sanitize(*comm)
}

func SanitizeUserBasic(login *domain.UserBasic) {
	sanitizer := bluemonday.UGCPolicy()
	login.Email = sanitizer.Sanitize(login.Email)
}
