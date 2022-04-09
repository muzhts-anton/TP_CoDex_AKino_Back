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
	SELECT COUNT(*) FROM Collections;
	`

	queryGetCollections = `
	SELECT * FROM Collections
	JOIN Movies on Collections.id = Movies.incollection
	WHERE Collections.id = $1;
	`

	queryGetFeed = `
	SELECT * FROM Feeds;
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

	dbsize := binary.BigEndian.Uint64(result[0][0])
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
			Id:          fmt.Sprint((binary.BigEndian.Uint64(respColl[i][3]))),
			ImgHref:     string(respColl[i][4]),
			Title:       string(respColl[i][5]),
			Rating:      fmt.Sprint(math.Float64frombits(binary.BigEndian.Uint64(respColl[i][6]))),
			Info:        string(respColl[i][7]),
			Description: string(respColl[i][8]),
		})
	}

	out := domain.Collection{
		Title:       string(respColl[0][1]),
		Description: string(respColl[0][2]),
		MovieList:   movies,
	}

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.Feed, error) {
	respFeed, err := cr.dbm.Query(queryGetFeed)
	if err != nil {
		return domain.Feed{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.FeedRow, 0)
	for i := range respFeed {
		movies = append(movies, domain.FeedRow{
			Description: string(respFeed[i][1]),
			ImgSrc:      string(respFeed[i][2]),
			Page:        string(respFeed[i][3]),
			Num:         fmt.Sprint((binary.BigEndian.Uint64(respFeed[i][4]))),
		})
	}

	out := domain.Feed{
		Coll: movies,
	}

	return out, nil
}
