package genusecase

import (
	"codex/internal/pkg/domain"

	"strings"
)

type genresUsecase struct {
	genresRepo domain.GenresRepository
}

func InitColUsc(gr domain.GenresRepository) domain.GenresUsecase {
	return &genresUsecase{
		genresRepo: gr,
	}
}

func (gu genresUsecase) GetMovies(genre string) ([]domain.MovieBasic, error) {
	if len(strings.Trim(genre, " ")) > 50 {
		return []domain.MovieBasic{}, domain.Err.ErrObj.BadInput // 'cause genre is VARCHAR(50) in db abyway
	}

	movs, err := gu.genresRepo.GetMovies(strings.Trim(genre, " "))
	if err != nil {
		return []domain.MovieBasic{}, err
	}

	return movs, err
}
