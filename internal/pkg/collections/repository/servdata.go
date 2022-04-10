package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbCollectionsRepository struct {
	dbm *database.DBManager
}

func InitColRep(manager *database.DBManager) domain.CollectionsRepository {
	return &dbCollectionsRepository{
		dbm: manager,
	}
}

func (cr *dbCollectionsRepository) GetCollection(id uint64) (domain.Collection, error) {
	resp, err := cr.dbm.Query(queryCountCollections)
	if err != nil {
		log.Warn("{GetCollection} in query: " + queryCountCollections)
		log.Error(err)
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	if id > cast.ToUint64(resp[0][0]) {
		log.Warn("{GetCollection}")
		log.Error(domain.Err.ErrObj.SmallBd)
		return domain.Collection{}, domain.Err.ErrObj.SmallBd
	}

	resp, err = cr.dbm.Query(queryGetCollections, id)
	if err != nil {
		log.Warn("{GetCollection} in query: " + queryGetCollections)
		log.Error(err)
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.MovieBasic, 0)
	for i := range resp {
		movies = append(movies, domain.MovieBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][2])),
			Poster:      cast.ToString(resp[i][3]),
			Title:       cast.ToString(resp[i][4]),
			Rating:      cast.FlToStr(cast.ToFloat64(resp[i][5])),
			Info:        cast.ToString(resp[i][6]),
			Description: cast.ToString(resp[i][7]),
		})
	}

	out := domain.Collection{
		Title:       cast.ToString(resp[0][0]),
		Description: cast.ToString(resp[0][1]),
		MovieList:   movies,
	}

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.FeedResponse, error) {
	resp, err := cr.dbm.Query(queryGetFeed)
	if err != nil {
		log.Warn("{GetFeed} in query: " + queryGetFeed)
		log.Error(err)
		return domain.FeedResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovies}")
		log.Error(domain.Err.ErrObj.SmallBd)
		return domain.FeedResponse{}, domain.Err.ErrObj.SmallBd
	}

	movies := make([]domain.Feed, 0)
	for i := range resp {
		movies = append(movies, domain.Feed{
			Description: cast.ToString(resp[i][0]),
			ImgSrc:      cast.ToString(resp[i][1]),
			Page:        cast.ToString(resp[i][2]),
			Num:         cast.IntToStr(cast.ToUint64(resp[i][3])),
		})
	}

	out := domain.FeedResponse{
		CollectionList: movies,
	}

	return out, nil
}
