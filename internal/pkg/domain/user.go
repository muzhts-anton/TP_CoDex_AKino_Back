package domain

const BaseUserPicture = "/profile.svg"

type User struct {
	Id             uint64 `json:"-"`
	Username       string `json:"username"`
	Password       string `json:"password,omitempty"`
	Email          string `json:"email"`
	RepeatPassword string `json:"repeatpassword,omitempty"`
}

type UserBasic struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetById(id uint64) (User, error)
	GetByEmail(email string) (User, error)
	AddUser(user User) (uint64, error)
}

type UserUsecase interface {
	GetBasicInfo(id uint64) (User, error)
	Register(us User) (User, error)
	Login(ub UserBasic) (User, error)
	CheckAuth(id uint64) (User, error)
}

func (us *User) ClearPasswords() User {
	us.Password = ""
	us.RepeatPassword = ""
	return *us
}
