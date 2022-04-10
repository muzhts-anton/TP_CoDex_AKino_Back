package usrrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"

	"encoding/binary"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryGetByEmail = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE email = $1;
	`

	queryGetById = `
	SELECT id, username, email, imgsrc
	FROM users
	WHERE id = $1;
	`

	queryAddUser = `
	INSERT INTO
		users (username, email, password)
	VALUES
		($1, $2, $3)
	RETURNING id;
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

	row := resp[0]
	out := domain.User{
		Id:             binary.BigEndian.Uint64(row[0]),
		Username:       string(row[1]),
		Password:       "",
		Email:          string(row[2]),
		Imgsrc:         string(row[3]),
		RepeatPassword: "",
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

	row := resp[0]
	out := domain.User{
		Id:             binary.BigEndian.Uint64(row[0]),
		Username:       string(row[1]),
		Password:       "",
		Email:          string(row[2]),
		Imgsrc:         string(row[3]),
		RepeatPassword: "",
	}

	return out, nil
}

func (ur *dbUserRepository) AddUser(us domain.User) (uint64, error) {
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(us.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, domain.Err.ErrObj.InternalServer
	}

	resp, err := ur.dbm.Query(queryAddUser, us.Username, us.Email, passwordByte)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint64(resp[0][0]), nil
}
