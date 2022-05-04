package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/config"
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
	resp, err := cr.dbm.Query(queryGetCollectionBasic, id)
	if err != nil {
		log.Warn("{GetCollection} in query: " + queryGetCollectionBasic)
		log.Error(err)
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetCollection}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Collection{}, domain.Err.ErrObj.SmallDb
	}
	out := domain.Collection{
		Title:       cast.ToString(resp[0][0]),
		Description: cast.ToString(resp[0][1]),
		Public:      cast.ToBool(resp[0][2]),
	}

	resp, err = cr.dbm.Query(queryGetCollectionMovies, id)
	if err != nil {
		log.Warn("{GetCollection} in query: " + queryGetCollectionMovies)
		log.Error(err)
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
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
	out.MovieList = movies

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.FeedResponse, error) {
	resp, err := cr.dbm.Query(queryGetFeed, config.ProdConfigStore.Feed)
	if err != nil {
		log.Warn("{GetFeed} in query: " + queryGetFeed)
		log.Error(err)
		return domain.FeedResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovies}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.FeedResponse{}, domain.Err.ErrObj.SmallDb
	}

	movies := make([]domain.Feed, 0)
	for i := range resp {
		movies = append(movies, domain.Feed{
			Description: cast.ToString(resp[i][0]),
			ImgSrc:      cast.ToString(resp[i][1]),
			Page:        "collections",
			Num:         cast.IntToStr(cast.ToUint64(resp[i][2])),
		})
	}

	out := domain.FeedResponse{
		CollectionList: movies,
	}

	return out, nil
}
