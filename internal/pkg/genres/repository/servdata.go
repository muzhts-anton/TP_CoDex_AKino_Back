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

func InitColRep(manager *database.DBManager) domain.GenresRepository {
	return &dbGenresRepository{
		dbm: manager,
	}
}

func (cr *dbGenresRepository) GetMovies(genre string) ([]domain.MovieBasic, error) {
	resp, err := cr.dbm.Query(queryGetMovies, genre, config.ProdConfigStore.Genres)
	if err != nil {
		log.Warn("{GetMovies} in query: " + queryGetMovies)
		log.Error(err)
		return []domain.MovieBasic{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.MovieBasic{}, domain.Err.ErrObj.BadGenre
	}

	movies := make([]domain.MovieBasic, 0)
	for i := range resp {
		movies = append(movies, domain.MovieBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Rating:      cast.FlToStr(cast.ToFloat64(resp[i][3])),
			Info:        cast.ToString(resp[i][4]),
			Description: cast.ToString(resp[i][5]),
		})
	}

	return movies, nil
}
