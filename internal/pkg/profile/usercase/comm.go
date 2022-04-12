package profileusecase

import (
	"codex/internal/pkg/domain"
)

type ProfileUsecase struct {
	profileRepo domain.ProfileRepository
}

func InitProfileUsc(pr domain.ProfileRepository) domain.ProfileUsecase {
	return &ProfileUsecase{
		ProfileRepo: pr,
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

func (mu movieUsecase) PostRating(movieId uint64, userId uint64, rating int) (float64, error) {
	if rating < 1 || rating > 10 {
		return 0.0, domain.Err.ErrObj.InvalidRating
	}

	newRating, err := mu.movieRepo.PostRating(movieId, userId, rating)
	if err != nil {
		return 0.0, err
	}

	return newRating, nil
}

func (mu movieUsecase) PostComment(movieId uint64, userId uint64, content string, commenttype int) (error) {
	var comtype string
	if commenttype == 1 {
		comtype = "good"
	} else if commenttype == 2 {
		comtype = "neutral"
	} else if commenttype == 3 {
		comtype = "bad"
	} else {
		return domain.Err.ErrObj.InvalidCommentType
	}

	return mu.movieRepo.PostComment(movieId, userId, content, comtype)
}
