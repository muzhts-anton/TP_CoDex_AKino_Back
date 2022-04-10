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

func (cu movieUsecase) GetMovie(id uint64) (domain.Movie, error) {
	movie, err := cu.movieRepo.GetMovie(id)
	if err != nil {
		return domain.Movie{}, err
	}

	return movie, err
}

func (cu movieUsecase) GetRelated(id uint64) ([]domain.MovieSummary, error) {
	related, err := cu.movieRepo.GetRelated(id)
	if err != nil {
		return nil, err
	}

	return related, err
}

func (cu movieUsecase) GetComments(id uint64) ([]domain.Comment, error) {
	comments, err := cu.movieRepo.GetComments(id)
	if err != nil {
		return nil, err
	}

	return comments, err
}

func (cu movieUsecase) PostRating(id uint64, authorId uint64, rating float64) (float64, error) {
	return 0.0, nil
}

func (cu movieUsecase) PostComment(id uint64, authorId uint64) (domain.Comment, error) {
	return domain.Comment{}, nil
}
