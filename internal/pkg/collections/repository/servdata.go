package colrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"

	"encoding/binary"
)



const (
	queryCountCollections = `
	SELECT COUNT(*) FROM Collection;
	`

	queryGetCollections = `
	SELECT * FROM collection
	JOIN movies ON collection.id = movies.collid
	WHERE collection.id = $1;
	`

	queryGetFeed = `
	SELECT * FROM feed
	`
)

type dbCollectionsRepository struct {
	dbm *database.DBManager
}

func InitColRep(manager *database.DBManager) domain.CollectionsRepository {
	return &dbCollectionsRepository{dbm: manager}
}

func (cr *dbCollectionsRepository) GetCollection(id uint64) (domain.CollType, error) {
	result, err := cr.dbm.Query(queryCountCollections)
	if err != nil {
		return domain.CollType{}, domain.Err.ErrObj.InternalServer
	}

	dbsize := binary.BigEndian.Uint64(result[0][0])
	if id >= dbsize {
		return domain.CollType{}, domain.Err.ErrObj.SmallBd
	}

	respColl, err := cr.dbm.Query(queryGetCollections, id)
	if err != nil {
		return domain.CollType{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.MovieType, 0)
	for i := range respColl {
		movies = append(movies, domain.MovieType{
			MovieHref:   string(respColl[i][2]),
			ImgHref:     string(respColl[i][3]),
			Title:       string(respColl[i][4]),
			Info:        string(respColl[i][5]),
			Rating:      string(respColl[i][6]),
			Description: string(respColl[i][7]),
		})
	}

	out := domain.CollType{
		Title:       string(respColl[0][0]),
		Description: string(respColl[0][1]),
		MovieList:   movies,
	}

	return out, nil
}

func (cr *dbCollectionsRepository) GetFeed() (domain.FilmSelection, error) {
	respFeed, err := cr.dbm.Query(queryGetFeed)
	if err != nil {
		return domain.FilmSelection{}, domain.Err.ErrObj.InternalServer
	}

	movies := make([]domain.FilmType, 0)
	for i := range respFeed {
		movies = append(movies, domain.FilmType{
			Description: string(respFeed[i][1]),
			ImgSrc:      string(respFeed[i][2]),
			Page:        string(respFeed[i][3]),
			Number:      string(respFeed[i][4]),
		})
	}

	out := domain.FilmSelection{
		Coll: movies,
	}

	return out, nil
}
