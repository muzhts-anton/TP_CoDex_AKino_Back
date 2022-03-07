package DB

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	// "strings"
	"sync"
)

type UserMockDatabase struct {
	sync.RWMutex
	users []User
}

type User struct {
	ID       uint64   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RepeatPassword string `json:"repeatpassword"`
	Email    string `json:"email"`
}


var basePicture = "/pic/1.jpg"

func (us *User) OmitPassword() {
	us.Password = ""
	// us.PasswordRepeat = ""
}

// func (us *User) OmitId() {
// 	us.ID = 0
// }

// func (us *User) OmitPic() {
// 	us.ProfilePic = ""
// }

var errorNoUser = errors.New("error: no user")

func (db *UserMockDatabase) AddUser(us *User) uint64 {
	db.RLock()
	us.ID = uint64(len(db.users) + 1)
	db.RUnlock()
	// us.Email = strings.ToLower(us.Email)
	// us.Surname = strings.Title(us.Surname)
	// us.FirstName = strings.Title(us.FirstName)
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0
	}
	us.Password = string(passwordByte)
	// us.ProfilePic = basePicture

	db.Lock()
	db.users = append(db.users, *us)
	db.Unlock()
	return us.ID
}

func (db *UserMockDatabase) FindEmail(email string) (User, error) {
	db.RLock()
	defer db.RUnlock()
	// loweredEmail := strings.ToLower(email)
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
	// loweredEmail := strings.ToLower(email)
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

// type CollectionPreview struct {
// 	Id         uint   `json:"id"`
// 	Title      string `json:"title"`
// 	PictureUrl string `json:"picture_url"`
// }

// var PreviewMock = []CollectionPreview{
// 	{Id: 1, Title: "Для ценителей Хогвартса", PictureUrl: "server/images/collections1.png"},
// 	{Id: 2, Title: "Про настоящую любовь", PictureUrl: "server/images/collections2.png"},
// 	{Id: 3, Title: "Аферы века", PictureUrl: "server/images/collections3.png"},
// 	{Id: 4, Title: "Про Вторую Мировую", PictureUrl: "server/images/collections4.jpg"},
// 	{Id: 5, Title: "Осеннее настроение", PictureUrl: "server/images/collections5.png"},
// 	{Id: 6, Title: "Летняя атмосфера", PictureUrl: "server/images/collections6.png"},
// 	{Id: 7, Title: "Про дружбу", PictureUrl: "server/images/collections7.png"},
// 	{Id: 8, Title: "Романтические фильмы", PictureUrl: "server/images/collections8.jpg"},
// 	{Id: 9, Title: "Джунгли зовут", PictureUrl: "server/images/collections9.jpg"},
// 	{Id: 10, Title: "Фантастические фильмы", PictureUrl: "server/images/collections10.jpg"},
// 	{Id: 11, Title: "Про петлю времени", PictureUrl: "server/images/collections11.png"},
// 	{Id: 12, Title: "Классика на века", PictureUrl: "server/images/collections12.jpg"},
// }

// type CollectionsMockDatabase struct {
// 	Previews []CollectionPreview
// 	sync.RWMutex
// }
