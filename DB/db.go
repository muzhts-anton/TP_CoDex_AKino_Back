package DB

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserMockDatabase struct {
	sync.RWMutex
	users []User
}

type User struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
}

const basePicture = "/pic/1.jpg"

func (us *User) OmitPassword() {
	us.Password = ""
}

var errorNoUser = errors.New("error: no user")

func (db *UserMockDatabase) AddUser(us *User) uint64 {
	db.RLock()
	us.ID = uint64(len(db.users) + 1)
	db.RUnlock()
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0
	}
	us.Password = string(passwordByte)

	db.Lock()
	db.users = append(db.users, *us)
	db.Unlock()
	return us.ID
}

func (db *UserMockDatabase) FindEmail(email string) (User, error) {
	db.RLock()
	defer db.RUnlock()
	for _, us := range db.users {
		if us.Email == email {
			return us, nil
		}
	}
	return User{}, errorNoUser
}

func (db *UserMockDatabase) FindUsername(username string) (User, error) {
	db.RLock()
	defer db.RUnlock()
	for _, us := range db.users {
		if us.Username == username {
			return us, nil
		}
	}
	return User{}, errorNoUser
}

func (db *UserMockDatabase) FindId(id uint64) (User, error) {
	db.RLock()
	defer db.RUnlock()
	if int(id) <= len(db.users) && id != 0 {
		return db.users[id-1], nil
	}
	return User{}, errorNoUser
}
