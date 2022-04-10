package movrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	
	"math"
	"time"
)

type dbMovieRepository struct {
	dbm *database.DBManager
}

func InitMovRep(manager *database.DBManager) domain.MovieRepository {
	return &dbMovieRepository{
		dbm: manager,
	}
}

func (mr *dbMovieRepository) GetMovie(id uint64) (domain.Movie, error) {
	resp, err := mr.dbm.Query(queryGetMovie, id)
	if err != nil {
		return domain.Movie{}, domain.Err.ErrObj.InternalServer
	}

	row := resp[0]
	out := domain.Movie{
		Id:            cast.IntToStr(cast.ToUint64(row[0])),
		Poster:        cast.ToString(row[1]),
		Title:         cast.ToString(row[2]),
		TitleOriginal: cast.ToString(row[3]),
		Rating:        cast.FlToStr(cast.ToFloat64(row[4])),
		Info:          cast.ToString(row[5]),
		Description:   cast.ToString(row[6]),
		Trailer:       cast.ToString(row[7]),
		ReleaseYear:   cast.ToString(row[8]),
		Country:       cast.ToString(row[9]),
		Genre:         cast.ToString(row[10]),
		Motto:         cast.ToString(row[11]),
		Director:      cast.ToString(row[12]),
		Budget:        cast.ToString(row[13]),
		Gross:         cast.ToString(row[14]),
		Duration:      cast.ToString(row[15]),
	}

	return out, nil
}

func (mr *dbMovieRepository) GetRelated(id uint64) ([]domain.MovieSummary, error) {
	resp, err := mr.dbm.Query(queryGetRelated, id)
	if err != nil {
		return nil, domain.Err.ErrObj.InternalServer
	}

	out := make([]domain.MovieSummary, 0)
	for i := range resp {
		out = append(out, domain.MovieSummary{
			Id:     cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster: cast.ToString(resp[i][1]),
			Title:  cast.ToString(resp[i][2]),
		})
	}

	return out, nil
}

func (mr *dbMovieRepository) GetComments(id uint64) ([]domain.Comment, error) {
	resp, err := mr.dbm.Query(queryGetComment, id)
	if err != nil {
		return nil, domain.Err.ErrObj.InternalServer
	}

	out := make([]domain.Comment, 0)
	for i := range resp {
		comm := domain.Comment{
			Imgsrc:   cast.ToString(resp[i][0]),
			Username: cast.ToString(resp[i][1]),
			UserId:   cast.IntToStr(cast.ToUint64(resp[i][2])),
			Date:     cast.ToString(resp[i][3]),
			Content:  cast.ToString(resp[i][4]),
			Type:     cast.ToString(resp[i][5]),
			Rating:   "",
		}

		tmp, err := mr.dbm.Query(queryGetRatingCount, comm.UserId)
		if err != nil {
			return nil, domain.Err.ErrObj.InternalServer
		}

		if cast.ToUint64(tmp[0][0]) == 1 {
			comm.Rating = cast.IntToStr(cast.ToUint64(resp[i][6]))
		}

		out = append(out, comm)
	}

	return out, nil
}

func (mr *dbMovieRepository) PostRating(movieId uint64, userId uint64, rating int) (float64, error) {
	// checking ids
	resp, err := mr.dbm.Query(queryUserExist, userId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		return 0.0, domain.Err.ErrObj.InvalidId
	}

	resp, err = mr.dbm.Query(queryMovieExist, movieId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		return 0.0, domain.Err.ErrObj.InvalidId
	}

	// get info from db
	resp, err = mr.dbm.Query(queryGetMovieRating, movieId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	oldRating := cast.ToFloat64(resp[0][0])

	resp, err = mr.dbm.Query(queryGetMovieVotesnum, movieId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	oldVotesnum := cast.ToUint64(resp[0][0])

	// compute new rating and push it to db movie table
	newRating := (oldRating*float64(oldVotesnum) + float64(rating)) / float64(oldVotesnum+1)
	newRating = math.Round(newRating*100) / 100

	_, err = mr.dbm.Query(queryIncrementVotesnum, movieId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	_, err = mr.dbm.Query(querySetMovieRating, newRating, movieId)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	// append info to ratings table
	_, err = mr.dbm.Query(queryPostRating, userId, movieId, rating)
	if err != nil {
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	return newRating, nil
}

func (mr *dbMovieRepository) PostComment(movieId uint64, userId uint64, content string, comtype string) (error) {
	// time.Now().Format("2006-01-02 15:04:05")
	// checking ids
	resp, err := mr.dbm.Query(queryUserExist, userId)
	if err != nil {
		return domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		return domain.Err.ErrObj.InvalidId
	}

	resp, err = mr.dbm.Query(queryMovieExist, movieId)
	if err != nil {
		return domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		return domain.Err.ErrObj.InvalidId
	}

	// post comment
	_, err = mr.dbm.Query(queryPostComment, userId, movieId, time.Now().Format("2006-01-02 15:04:05"), comtype, content)
	if err != nil {
		return domain.Err.ErrObj.InternalServer
	}

	return nil
}
