package usrrepository

import (
	"codex/internal/pkg/domain"
	
	"golang.org/x/crypto/bcrypt"
	"strings"
	"sync"
)

type dbUserRepository struct {
	sync.RWMutex
	users []domain.User
}

func InitUsrRep() domain.UserRepository {
	return &dbUserRepository{}
}

func (ur *dbUserRepository) GetByEmail(email string) (domain.User, error) {
	ur.RLock()
	defer ur.RUnlock()
	
	loweredEmail := strings.ToLower(email)
	for _, us := range ur.users {
		if us.Email == loweredEmail {
			return us, nil
		}
	}

	return domain.User{}, domain.Err.ErrObj.NoUser
}

func (ur *dbUserRepository) GetById(id uint64) (domain.User, error) {
	ur.RLock()
	defer ur.RUnlock()

	if int(id) <= len(ur.users) && id != 0 {
		return ur.users[id-1], nil
	}

	return domain.User{}, domain.Err.ErrObj.NoUser
}

func (ur *dbUserRepository) AddUser(us domain.User) (uint64, error) {
	ur.RLock()
	us.Id = uint64(len(ur.users) + 1)
	ur.RUnlock()

	us.Email = strings.ToLower(us.Email)
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, domain.Err.ErrObj.InternalServer
	}
	us.Password = string(passwordByte)

	ur.Lock()
	ur.users = append(ur.users, us)
	ur.Unlock()

	return us.Id, nil
}
