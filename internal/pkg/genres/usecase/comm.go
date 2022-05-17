package genusecase

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/addPreview"
	"strings"
)

type genresUsecase struct {
	genresRepo domain.GenresRepository
}

func InitGenUsc(gr domain.GenresRepository) domain.GenresUsecase {
	return &genresUsecase{
		genresRepo: gr,
	}
}

func (gu genresUsecase) GetGenre(genre string) (domain.GenreWithMovies, error) {
	if len(strings.Trim(genre, " ")) > 50 {
		return domain.GenreWithMovies{}, domain.Err.ErrObj.BadInput // 'cause genre is VARCHAR(50) in db abyway
	}

	movs, err := gu.genresRepo.GetGenre(strings.Trim(genre, " "))
	if err != nil {
		return domain.GenreWithMovies{}, err
	}

	for index := range movs.MovieList {
		movs.MovieList[index].Poster = addPreview.ToMiniCopy(movs.MovieList[index].Poster)
	}

	return movs, err
}

func (gu genresUsecase) GetGenres() ([]domain.Genre, error) {
	genres, err := gu.genresRepo.GetGenres()
	if err != nil {
		return []domain.Genre{}, err
	}

	return genres, err
}
