package annrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbAnnouncedRepository struct {
	dbm *database.DBManager
}

func InitAnnRep(manager *database.DBManager) domain.AnnouncedRepository {
	return &dbAnnouncedRepository{
		dbm: manager,
	}
}

func (cr *dbAnnouncedRepository) GetMovies() ([]domain.AnnouncedBasic, error) {
	resp, err := cr.dbm.Query(queryGetMovies)
	if err != nil {
		log.Warn("{GetMovies} in query: " + queryGetMovies)
		log.Error(err)
		return []domain.AnnouncedBasic{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.AnnouncedBasic{}, domain.Err.ErrObj.SmallDb
	}

	movies := make([]domain.AnnouncedBasic, 0)
	for i := range resp {
		movies = append(movies, domain.AnnouncedBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Info:        "Дата премьеры: " + cast.TimeToStr(cast.ToTime(resp[i][3]), false) + ". Осталось " +  cast.ToString(resp[i][4]) + " дня.",
			Description: cast.ToString(resp[i][5]),
		})
	}

	return movies, nil
}
