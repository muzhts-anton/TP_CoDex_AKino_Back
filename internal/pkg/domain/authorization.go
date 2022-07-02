package domain

type AuthRepository interface {
	GetById(id uint64) (User, error)
	GetByEmail(email string) (User, error)
	AddUser(user User) (uint64, error)
}

type AuthUsecase interface {
	Register(us User) (User, error)
	Login(ub UserBasic) (User, error)
}
