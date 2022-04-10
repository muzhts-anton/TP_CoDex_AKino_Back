package actrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"

	_ "math"
	_ "time"
)

type dbActorRepository struct {
	dbm *database.DBManager
}

func InitActRep(manager *database.DBManager) domain.ActorRepository {
	return &dbActorRepository{
		dbm: manager,
	}
}

func (ar *dbActorRepository) GetActor(id uint64) (domain.Actor, error) {
	resp, err := ar.dbm.Query(queryGetActor, id)
	if err != nil {
		return domain.Actor{}, domain.Err.ErrObj.InternalServer
	}

	actor := domain.Actor{
		Id:           cast.IntToStr(cast.ToUint64(resp[0][0])),
		Imgsrc:       cast.ToString(resp[0][1]),
		Name:         cast.ToString(resp[0][2]),
		NameOriginal: cast.ToString(resp[0][3]),
		Career:       cast.ToString(resp[0][4]),
		Height:       cast.ToString(resp[0][5]),
		Birthday:     cast.ToString(resp[0][6]),
		Birthplace:   cast.ToString(resp[0][7]),
		Genres:       cast.ToString(resp[0][8]),
		Total:        cast.IntToStr(cast.ToUint64(resp[0][9])),
	}

	return actor, nil
}

func (ar *dbActorRepository) GetMovies(id uint64) ([]domain.MovieBasic, error) {
	resp, err := ar.dbm.Query(queryGetMovies, id)
	if err != nil {
		return nil, domain.Err.ErrObj.InternalServer
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

func (ar *dbActorRepository) GetRelated(id uint64) ([]domain.ActorBasic, error) {
	resp, err := ar.dbm.Query(queryGetRelated, id)
	if err != nil {
		return nil, domain.Err.ErrObj.InternalServer
	}

	actors := make([]domain.ActorBasic, 0)
	for i := range resp {
		actors = append(actors, domain.ActorBasic{
			Href:   "/actors/" + cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster: cast.ToString(resp[i][1]),
			Name:   cast.ToString(resp[i][2]),
		})
	}

	return actors, nil
}
