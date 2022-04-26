package movusecase

import (
	"codex/internal/pkg/domain"
)

type movieUsecase struct {
	movieRepo domain.MovieRepository
}

func InitMovUsc(mr domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieRepo: mr,
	}
}

func (mu movieUsecase) GetMovie(id uint64) (domain.Movie, error) {
	movie, err := mu.movieRepo.GetMovie(id)
	if err != nil {
		return domain.Movie{}, err
	}

	return movie, err
}

func (mu movieUsecase) GetRelated(id uint64) ([]domain.MovieSummary, error) {
	related, err := mu.movieRepo.GetRelated(id)
	if err != nil {
		return nil, err
	}

	return related, err
}

func (mu movieUsecase) GetComments(id uint64) ([]domain.Comment, error) {
	comments, err := mu.movieRepo.GetComments(id)
	if err != nil {
		return nil, err
	}

	return comments, err
}

func (mu movieUsecase) GetReviewRating(movieId, userId uint64) (string, string, error) {
	reviewExist, userRating, err := mu.movieRepo.GetReviewRating(movieId, userId)
	if err != nil {
		return "", "", err
	}

	return reviewExist, userRating, nil
}
