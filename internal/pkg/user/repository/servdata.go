package usrrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	_"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

	"encoding/binary"
	"golang.org/x/crypto/bcrypt"
)

type dbUserRepository struct {
	dbm *database.DBManager
}

func InitUsrRep(manager *database.DBManager) domain.UserRepository {
	return &dbUserRepository{
		dbm: manager,
	}
}

func (ur *dbUserRepository) GetByEmail(email string) (domain.User, error) {
	resp, err := ur.dbm.Query(queryGetByEmail, email)
	if len(resp) == 0 {
		log.Warn("{GetByEmail}")
		log.Error(domain.Err.ErrObj.NoUser)
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		log.Warn("{GetByEmail} in query: " + queryGetByEmail)
		log.Error(err)
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
		log.Warn("{GetById}")
		log.Error(domain.Err.ErrObj.NoUser)
		return domain.User{}, domain.Err.ErrObj.NoUser
	}
	if err != nil {
		log.Warn("{GetById} in query: " + queryGetById)
		log.Error(err)
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
		log.Warn("{AddUser}")
		log.Error(err)
		return 0, domain.Err.ErrObj.InternalServer
	}

	resp, err := ur.dbm.Query(queryAddUser, us.Username, us.Email, passwordByte)
	if err != nil {
		log.Warn("{AddUser} in query: " + queryAddUser)
		log.Error(err)
		return 0, err
	}

	return binary.BigEndian.Uint64(resp[0][0]), nil
}
