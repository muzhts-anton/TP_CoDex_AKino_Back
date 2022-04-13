package domain

const BaseUserPicture = "/profile.svg"

type User struct {
	Id             uint64 `json:"ID"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email"`
	Imgsrc         string `json:"imgsrc"`
	RepeatPassword string `json:"repeatpassword,omitempty"`
}

func (us *User) ClearPasswords() User {
	us.Password = ""
	us.RepeatPassword = ""
	return *us
}

type UserBasic struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdUser struct {
	Username string `json:"name"`
	Imgsrc   string `json:"imgsrc"`
}

type UserRepository interface {
	GetById(id uint64) (User, error)
	GetByEmail(email string) (User, error)
	AddUser(user User) (uint64, error)
	GetBookmarks(id uint64) ([]Bookmark, error)
	UpdateUser(id uint64, upd UpdUser) (User, error)
	UpdateAvatar(id uint64, url string) (User, error)
}

type UserUsecase interface {
	Register(us User) (User, error)
	Login(ub UserBasic) (User, error)
	CheckAuth(id uint64) (User, error)
	GetBasicInfo(id uint64) (User, error)
	GetBookmarks(id uint64) ([]Bookmark, error)
	UpdateUser(id uint64, upd UpdUser) (User, error)
	UpdateAvatar(id uint64, url string) (User, error)
}
