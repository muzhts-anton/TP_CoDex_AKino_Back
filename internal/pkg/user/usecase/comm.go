package usrusecase

import (
	"codex/internal/pkg/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func InitUsrUsc(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (uc userUsecase) GetBasicInfo(id uint64) (domain.User, error) {
	us, err := uc.userRepo.GetById(id)
	if err != nil {
		return domain.User{}, domain.Err.ErrObj.InternalServer
	}

	return us.ClearPasswords(), nil
}

func (uc userUsecase) GetBookmarks(id uint64) ([]domain.Bookmark, error) {
	bookmarks, err := uc.userRepo.GetBookmarks(id)
	if err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (uc userUsecase) UpdateUser(id uint64, upd domain.UpdUser) (domain.User, error) {
	if validateUsername(upd.Username) != nil {
		return domain.User{}, domain.Err.ErrObj.InvalidUsername
	}

	usr, err := uc.userRepo.UpdateUser(id, upd)
	if err != nil {
		return domain.User{}, err
	}

	return usr, nil
}

func (uc userUsecase) GetUserReviews(id uint64) ([]domain.UserReview, error) {
	reviews, err := uc.userRepo.GetUserReviews(id)
	if err != nil {
		return []domain.UserReview{}, err
	}

	return reviews, nil
}

func (uc userUsecase) UpdateAvatar(clientID uint64, url string) (domain.User, error) {
	us, err := uc.userRepo.UpdateAvatar(clientID, url)
	if err != nil {
		return domain.User{}, err
	}
	return us, nil
}
