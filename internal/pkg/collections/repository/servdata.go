package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"

	"encoding/binary"
	"fmt"
	"math"
)

const (
	queryCountCollections = `
	SELECT COUNT(*) FROM collections;
	`

	queryGetCollections = `
	SELECT collections.title, collections.description, movies.id, movies.poster, movies.title, movies.rating, movies.info, movies.description
	FROM collections
	JOIN movies on collections.id = movies.incollection
	WHERE collections.id = $1;
	`

	queryGetFeed = `
	SELECT description, poster, page, num
	FROM feed;
	`
)

type dbCollectionsRepository struct {
	dbm *database.DBManager
}

func InitColRep(manager *database.DBManager) domain.CollectionsRepository {
	return &dbCollectionsRepository{dbm: manager}
}

func (cr *dbCollectionsRepository) GetCollection(id uint64) (domain.Collection, error) {
	result, err := cr.dbm.Query(queryCountCollections)
	if err != nil {
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	dbsize := binary.BigEndian.Uint64(result[0][0]) // may be unnecessary idk
	if id > dbsize {
		return domain.Collection{}, domain.Err.ErrObj.SmallBd
	}

	respColl, err := cr.dbm.Query(queryGetCollections, id)
	if err != nil {
		return domain.Collection{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.MovieRow, 0)
	for i := range respColl {
		movies = append(movies, domain.MovieRow{
			Id:          fmt.Sprint((binary.BigEndian.Uint64(respColl[i][2]))),
			Poster:      string(respColl[i][3]),
			Title:       string(respColl[i][4]),
			Rating:      fmt.Sprint(math.Float64frombits(binary.BigEndian.Uint64(respColl[i][5]))),
			Info:        string(respColl[i][6]),
			Description: string(respColl[i][7]),
		})
	}

	out := domain.Collection{
		Title:       string(respColl[0][0]),
		Description: string(respColl[0][1]),
		MovieList:   movies,
	}

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.Feed, error) {
	resp, err := cr.dbm.Query(queryGetFeed)
	if err != nil {
		return domain.Feed{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.FeedRow, 0)
	for i := range resp {
		movies = append(movies, domain.FeedRow{
			Description: string(resp[i][0]),
			ImgSrc:      string(resp[i][1]),
			Page:        string(resp[i][2]),
			Num:         fmt.Sprint((binary.BigEndian.Uint64(resp[i][3]))),
		})
	}

	out := domain.Feed{
		Coll: movies,
	}

	return out, nil
}
