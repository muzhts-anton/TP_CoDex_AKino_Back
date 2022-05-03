package autusecase

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/authorization/delivery/grpc"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	grpc.UnimplementedAutherServer
	authRepo domain.AuthRepository
}

func InitAutUsc(ar domain.AuthRepository) grpc.AutherServer {
	return &authUsecase{
		authRepo: ar,
	}
}

func (au authUsecase) Register(ctx context.Context, us *grpc.User) (*grpc.User, error) {
	trimCredentials(&us.Email, &us.Username, &us.Password, &us.RepeatPassword)

	if us.Email == "" || us.Username == "" || us.Password == "" || us.RepeatPassword == "" {
		return nil, domain.Err.ErrObj.EmptyField
	}

	if err := validateEmail(us.Email); err != nil {
		return nil, err
	}

	if err := ValidateUsername(us.Username); err != nil {
		return nil, err
	}

	if err := validatePassword(us.Password); err != nil {
		return nil, err
	}

	if us.Password != us.RepeatPassword {
		return nil, domain.Err.ErrObj.UnmatchedPasswords
	}

	if _, err := au.authRepo.GetByEmail(us.Email); err == nil {
		return nil, domain.Err.ErrObj.EmailExists
	}

	idupd, err := au.authRepo.AddUser(domain.User{
		Id:             us.GetID(),
		Username:       us.GetUsername(),
		Password:       us.GetPassword(),
		Email:          us.GetEmail(),
		Imgsrc:         us.GetImgsrc(),
		RepeatPassword: us.GetRepeatPassword(),
	})
	if err != nil {
		return nil, err
	}

	out, _ := au.authRepo.GetById(idupd)
	out = out.ClearPasswords()

	return &grpc.User{
		ID:             out.Id,
		Username:       out.Username,
		Password:       out.Password,
		Email:          out.Email,
		Imgsrc:         out.Imgsrc,
		RepeatPassword: out.RepeatPassword,
	}, nil
}

func (au authUsecase) Login(ctx context.Context, ub *grpc.UserBasic) (*grpc.User, error) {
	if ub.Email == "" || ub.Password == "" {
		return nil, domain.Err.ErrObj.EmptyField
	}

	usr, err := au.authRepo.GetByEmail(ub.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(ub.Password)); err != nil {
		return nil, domain.Err.ErrObj.BadPassword
	}

	usr = usr.ClearPasswords()

	return &grpc.User{
		ID:             usr.Id,
		Username:       usr.Username,
		Password:       usr.Password,
		Email:          usr.Email,
		Imgsrc:         usr.Imgsrc,
		RepeatPassword: usr.RepeatPassword,
	}, nil
}
