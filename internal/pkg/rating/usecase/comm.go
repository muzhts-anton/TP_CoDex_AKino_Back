package ratusecase

import (
	"codex/internal/pkg/domain"
)

type ratingUsecase struct {
	ratingRepo domain.RatingRepository
}

func InitRatUsc(rr domain.RatingRepository) domain.RatingUsecase {
	return &ratingUsecase{
		ratingRepo: rr,
	}
}

func (ru ratingUsecase) PostRating(movieId uint64, userId uint64, rating int) (float64, error) {
	if rating < 1 || rating > 10 {
		return 0.0, domain.Err.ErrObj.InvalidRating
	}

	newRating, err := ru.ratingRepo.PostRating(movieId, userId, rating)
	if err != nil {
		return 0.0, err
	}

	return newRating, nil
}
