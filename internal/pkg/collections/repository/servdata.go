package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"

	"encoding/binary"
	"fmt"
	"math"
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
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	dbsize := binary.BigEndian.Uint64(resp[0][0])
	if id > dbsize {
		return domain.Collection{}, domain.Err.ErrObj.SmallBd
	}

	resp, err = cr.dbm.Query(queryGetCollections, id)
	if err != nil {
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.MovieBasic, 0)
	for i := range resp {
		movies = append(movies, domain.MovieBasic{
			Id:          fmt.Sprint((binary.BigEndian.Uint64(resp[i][2]))),
			Poster:      string(resp[i][3]),
			Title:       string(resp[i][4]),
			Rating:      fmt.Sprint(math.Float64frombits(binary.BigEndian.Uint64(resp[i][5]))),
			Info:        string(resp[i][6]),
			Description: string(resp[i][7]),
		})
	}

	out := domain.Collection{
		Title:       string(resp[0][0]),
		Description: string(resp[0][1]),
		MovieList:   movies,
	}

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.FeedResponse, error) {
	resp, err := cr.dbm.Query(queryGetFeed)
	if err != nil {
		return domain.FeedResponse{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.Feed, 0)
	for i := range resp {
		movies = append(movies, domain.Feed{
			Description: string(resp[i][0]),
			ImgSrc:      string(resp[i][1]),
			Page:        string(resp[i][2]),
			Num:         fmt.Sprint((binary.BigEndian.Uint64(resp[i][3]))),
		})
	}

	out := domain.FeedResponse{
		CollectionList: movies,
	}

	return out, nil
}
