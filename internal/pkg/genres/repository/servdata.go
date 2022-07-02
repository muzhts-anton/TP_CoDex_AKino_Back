package genrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/config"
	"codex/internal/pkg/utils/log"
)

type dbGenresRepository struct {
	dbm *database.DBManager
}

func InitGenRep(manager *database.DBManager) domain.GenresRepository {
	return &dbGenresRepository{
		dbm: manager,
	}
}

func (cr *dbGenresRepository) GetGenre(genre string) (domain.GenreWithMovies, error) {
	resp, err := cr.dbm.Query(queryGetGenreWithMovies, genre, config.ProdConfigStore.Genres)
	if err != nil {
		log.Warn("{GetMovies} in query: " + queryGetGenreWithMovies)
		log.Error(err)
		return domain.GenreWithMovies{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.GenreWithMovies{}, domain.Err.ErrObj.BadGenre
	}

	var genreWithMovies domain.GenreWithMovies
	genreWithMovies.Description = cast.ToString(resp[0][6])
	genreWithMovies.Title = cast.ToString(resp[0][7])
	for i := range resp {
		genreWithMovies.MovieList = append(genreWithMovies.MovieList, domain.MovieBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Rating:      cast.FlToStr(cast.ToFloat64(resp[i][3])),
			Info:        cast.ToString(resp[i][4]),
			Description: cast.ToString(resp[i][5]),
		})
	}

	return genreWithMovies, nil
}

func (cr *dbGenresRepository) GetGenres() ([]domain.Genre, error) {
	resp, err := cr.dbm.Query(queryGetGenres)
	if err != nil {
		log.Warn("{GetGenres} in query: " + queryGetGenres)
		log.Error(err)
		return []domain.Genre{}, domain.Err.ErrObj.InternalServer
	}

	genres := make([]domain.Genre, 0)
	for i := range resp {
		genres = append(genres, domain.Genre{
			Href:   "/genres/" + cast.ToString(resp[i][0]),
			Imgsrc: "genres" + cast.ToString(resp[i][1]),
		})
	}

	return genres, nil
}
