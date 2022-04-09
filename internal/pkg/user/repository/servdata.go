package usrrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"

	"encoding/binary"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryGetById = `
	SELECT * FROM Users
	WHERE id = $1
	`

	queryGetByEmail = `
	SELECT * FROM Users
	WHERE email = $1
	`

	queryAddUser = `
	INSERT INTO
		Users (username, email, password)
	VALUES
		($1, $2, $3)
	RETURNING id
	`
)

type dbUserRepository struct {
	dbm *database.DBManager
}

func InitUsrRep(manager *database.DBManager) domain.UserRepository {
	return &dbUserRepository{dbm: manager}
}

func (ur *dbUserRepository) GetByEmail(email string) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetByEmail, email)
	if len(resp) == 0 {
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	raw := resp[0]
	out := domain.User{
		Id:             binary.BigEndian.Uint64(raw[0]),
		Username:       string(raw[1]),
		Password:       string(raw[2]),
		Email:          string(raw[3]),
		RepeatPassword: string(raw[2]),
	}

	return out, nil
}

func (ur *dbUserRepository) GetById(id uint64) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetById, id)
	if len(resp) == 0 {
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	raw := resp[0]
	out := domain.User{
		Id:             binary.BigEndian.Uint64(raw[0]),
		Username:       string(raw[1]),
		Password:       string(raw[2]),
		Email:          string(raw[3]),
		RepeatPassword: string(raw[2]),
	}

	return out, nil
}

func (ur *dbUserRepository) AddUser(us domain.User) (uint64, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, domain.Err.ErrObj.InternalServer
	}

	us.Password = string(passwordByte)

	resp, err := ur.dbm.Query(queryAddUser, us.Username, us.Email, us.Password)

	us.Id = binary.BigEndian.Uint64(resp[0][0])

	return us.Id, nil
}
