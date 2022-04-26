package autrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

	"golang.org/x/crypto/bcrypt"
)

type dbAuthRepository struct {
	dbm *database.DBManager
}

func InitAutRep(manager *database.DBManager) domain.AuthRepository {
	return &dbAuthRepository{
		dbm: manager,
	}
}

func (ur *dbAuthRepository) GetByEmail(email string) (domain.User, error) {
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
		Id:             cast.ToUint64(row[0]),
		Username:       cast.ToString(row[1]),
		Password:       cast.ToString(row[4]),
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		RepeatPassword: cast.ToString(row[4]),
	}

	return out, nil
}

func (ur *dbAuthRepository) GetById(id uint64) (domain.User, error) {
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
		Id:             cast.ToUint64(row[0]),
		Username:       cast.ToString(row[1]),
		Password:       "",
		Email:          cast.ToString(row[2]),
		Imgsrc:         cast.ToString(row[3]),
		RepeatPassword: "",
	}

	return out, nil
}

func (ur *dbAuthRepository) AddUser(us domain.User) (uint64, error) {
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

	return cast.ToUint64(resp[0][0]), nil
}
