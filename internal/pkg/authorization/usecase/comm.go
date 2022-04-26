package autusecase

import (
	"codex/internal/pkg/domain"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authRepo domain.AuthRepository
}

func InitAutUsc(ar domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		authRepo: ar,
	}
}

func (au authUsecase) Register(us domain.User) (domain.User, error) {
	trimCredentials(&us.Email, &us.Username, &us.Password, &us.RepeatPassword)

	if us.Email == "" || us.Username == "" || us.Password == "" || us.RepeatPassword == "" {
		return domain.User{}, domain.Err.ErrObj.EmptyField
	}

	if err := validateEmail(us.Email); err != nil {
		return domain.User{}, err
	}

	if err := validateUsername(us.Username); err != nil {
		return domain.User{}, err
	}

	if err := validatePassword(us.Password); err != nil {
		return domain.User{}, err
	}

	if us.Password != us.RepeatPassword {
		return domain.User{}, domain.Err.ErrObj.UnmatchedPasswords
	}

	if _, err := au.authRepo.GetByEmail(us.Email); err == nil {
		return domain.User{}, domain.Err.ErrObj.EmailExists
	}

	idupd, err := au.authRepo.AddUser(us)
	if err != nil {
		return domain.User{}, err
	}

	out, _ := au.authRepo.GetById(idupd)

	return out.ClearPasswords(), nil
}

func (au authUsecase) Login(ub domain.UserBasic) (domain.User, error) {
	if ub.Email == "" || ub.Password == "" {
		return domain.User{}, domain.Err.ErrObj.EmptyField
	}

	usr, err := au.authRepo.GetByEmail(ub.Email)
	if err != nil {
		return domain.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(ub.Password)); err != nil {
		return domain.User{}, domain.Err.ErrObj.BadPassword
	}

	return usr.ClearPasswords(), nil
}
