package movrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

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
		log.Warn("{GetMovie} in query: " + queryGetMovie)
		log.Error(err)
		return domain.Movie{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovie}")
		log.Error(domain.Err.ErrObj.SmallBd)
		return domain.Movie{}, domain.Err.ErrObj.SmallBd
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
		Actors:        []domain.Cast{},
	}

	resp, err = mr.dbm.Query(queryGetMovieCast, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetMovieCast)
		log.Error(err)
		return domain.Movie{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovie} no cast o_0")
		log.Error(domain.Err.ErrObj.SmallBd)
		return domain.Movie{}, domain.Err.ErrObj.SmallBd
	}

	actors := make([]domain.Cast, 0)
	for i := range resp {
		actors = append(actors, domain.Cast{
			Name: cast.ToString(resp[i][0]),
			Href: "/actors/" + cast.IntToStr(cast.ToUint64(resp[i][1])),
		})
	}

	out.Actors = actors

	return out, nil
}

func (mr *dbMovieRepository) GetRelated(id uint64) ([]domain.MovieSummary, error) {
	resp, err := mr.dbm.Query(queryGetRelated, id)
	if err != nil {
		log.Warn("{GetRelated} in query: " + queryGetRelated)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.MovieSummary{}, nil
	}

	out := make([]domain.MovieSummary, 0)
	for i := range resp {
		out = append(out, domain.MovieSummary{
			Href:   "/movies/" + cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster: cast.ToString(resp[i][1]),
			Title:  cast.ToString(resp[i][2]),
		})
	}

	return out, nil
}

func (mr *dbMovieRepository) GetComments(id uint64) ([]domain.Comment, error) {
	resp, err := mr.dbm.Query(queryGetComment, id)
	if err != nil {
		log.Warn("{GetComments} in query: " + queryGetComment)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.Comment{}, nil
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
			log.Warn("{GetComments} in query: " + queryGetRatingCount)
			log.Error(err)
			return nil, domain.Err.ErrObj.InternalServer
		}

		if cast.ToUint64(tmp[0][0]) == 1 {
			comm.Rating = cast.IntToStr(cast.ToUint64(resp[i][6]))
		}

		out = append(out, comm)
	}

	return out, nil
}

func (mr *dbMovieRepository) GetReviewRating(movieId, userId uint64) (string, string, error) {
	resp, err := mr.dbm.Query(queryGetCommentsCount, movieId, userId)
	if err != nil {
		log.Warn("{PostComment} in query: " + queryGetCommentsCount)
		log.Error(err)
		return "", "", domain.Err.ErrObj.InternalServer
	}

	var reviewExist string
	if cast.ToUint64(resp[0][0]) == 0 {
		reviewExist = ""
	} else {
		reviewExist = cast.IntToStr(cast.ToUint64(resp[0][0]))
	}

	resp, err = mr.dbm.Query(queryGetUserRating, userId, movieId)
	if err != nil {
		log.Warn("{GetComment} in query: " + queryGetUserRating)
		log.Error(err)
		return "", "", domain.Err.ErrObj.InternalServer
	}

	var userRating string
	if len(resp) == 1 {
		userRating = cast.IntToStr(cast.ToUint64(resp[0][0]))
	} else {
		userRating = ""
	}

	return reviewExist, userRating, nil
}

func (mr *dbMovieRepository) PostRating(movieId uint64, userId uint64, rating int) (float64, error) {
	// checking ids
	resp, err := mr.dbm.Query(queryUserExist, userId)
	if err != nil {
		log.Warn("{PostRating} in query: " + queryUserExist)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		log.Warn("{PostRating}")
		log.Error(domain.Err.ErrObj.InvalidId)
		return 0.0, domain.Err.ErrObj.InvalidId
	}

	resp, err = mr.dbm.Query(queryMovieExist, movieId)
	if err != nil {
		log.Warn("{PostRating} in query: " + queryMovieExist)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		log.Warn("{PostRating}")
		log.Error(domain.Err.ErrObj.InvalidId)
		return 0.0, domain.Err.ErrObj.InvalidId
	}

	// get info from db
	resp, err = mr.dbm.Query(queryGetMovieRating, movieId)
	if err != nil {
		log.Warn("{PostRating} in query: " + queryGetMovieRating)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	var oldRating float64
	if len(resp) == 0 {
		oldRating = 0.0
	} else {
		oldRating = cast.ToFloat64(resp[0][0])
	}

	resp, err = mr.dbm.Query(queryGetMovieVotesnum, movieId)
	if err != nil {
		log.Warn("{PostRating} in query: " + queryGetMovieVotesnum)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	var oldVotesnum uint64
	if len(resp) == 0 {
		oldVotesnum = 0
	} else {
		oldVotesnum = cast.ToUint64(resp[0][0])
	}

	// check if rating is new for user
	resp, err = mr.dbm.Query(queryGetRatingUserCount, userId, movieId)
	if err != nil {
		log.Warn("{PostRating} in query: " + queryGetRatingUserCount)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	// 0 means the rating is new for the movie from this user. 1 means user changes his rating
	isOldRating := cast.ToUint64(resp[0][0])
	var newRating float64

	// compute new rating and push it to db movie table
	if isOldRating == 1 {
		resp, err = mr.dbm.Query(queryGetOldRatingUser, userId, movieId)
		if err != nil {
			log.Warn("{PostRating} in query: " + queryGetOldRatingUser)
			log.Error(err)
			return 0.0, domain.Err.ErrObj.InternalServer
		}
		userOldRating := cast.ToUint64(resp[0][0])
		if userOldRating > uint64(rating) {
			newRating = oldRating - ((float64(userOldRating - uint64(rating))) / float64(oldVotesnum))
		} else {
			newRating = oldRating + ((float64(uint64(rating) - userOldRating)) / float64(oldVotesnum))
		}
	} else {
		newRating = (oldRating*float64(oldVotesnum) + float64(rating)) / float64(oldVotesnum+1)
		//fmt.Println("oldR: %v\noldV: %v\n, rating: %v", oldRating, oldVotesnum, rating)
	}

	newRating = math.Round(newRating*100) / 100

	if isOldRating == 0 {
		_, err = mr.dbm.Query(queryIncrementVotesnum, movieId)
		if err != nil {
			log.Warn("{PostRating} in query: " + queryIncrementVotesnum)
			log.Error(err)
			return 0.0, domain.Err.ErrObj.InternalServer
		}
	}

	_, err = mr.dbm.Query(querySetMovieRating, newRating, movieId)
	if err != nil {
		log.Warn("{PostRating} in query: " + querySetMovieRating)
		log.Error(err)
		return 0.0, domain.Err.ErrObj.InternalServer
	}

	// append info to ratings table
	if isOldRating == 0 {
		_, err = mr.dbm.Query(queryPostRating, userId, movieId, rating)
		if err != nil {
			log.Warn("{PostRating} in query: " + queryPostRating)
			log.Error(err)
			return 0.0, domain.Err.ErrObj.InternalServer
		}
	} else {
		_, err = mr.dbm.Query(queryChangeRating, rating, userId)
		if err != nil {
			log.Warn("{PostRating} in query: " + queryChangeRating)
			log.Error(err)
			return 0.0, domain.Err.ErrObj.InternalServer
		}
	}

	return newRating, nil
}

func (mr *dbMovieRepository) PostComment(movieId uint64, userId uint64, content string, comtype string) (domain.Comment, error) {
	// checking ids
	resp, err := mr.dbm.Query(queryUserExist, userId)
	if err != nil {
		log.Warn("{PostComment} in query: " + queryUserExist)
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		log.Warn("{PostComment}")
		log.Error(domain.Err.ErrObj.InvalidId)
		return domain.Comment{}, domain.Err.ErrObj.InvalidId
	}

	resp, err = mr.dbm.Query(queryMovieExist, movieId)
	if err != nil {
		log.Warn("{PostComment} in query: " + queryMovieExist)
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}
	if cast.ToUint64(resp[0][0]) == 0 {
		log.Warn("{PostComment}")
		log.Error(domain.Err.ErrObj.InvalidId)
		return domain.Comment{}, domain.Err.ErrObj.InvalidId
	}

	// post comment
	_, err = mr.dbm.Query(queryPostComment, userId, movieId, time.Now().Format("2006-01-02 15:04:05"), comtype, content)
	if err != nil {
		log.Warn("{PostComment} in query: " + queryPostComment)
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}

	// get comment posted recently
	comm, err := mr.GetComment(userId, movieId)
	if err != nil {
		log.Warn("{PostComment} unable to get comment which I posted 2sec ago o_0")
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}

	return comm, nil
}

func (mr *dbMovieRepository) GetComment(userId, movieId uint64) (domain.Comment, error) {
	resp, err := mr.dbm.Query(queryGetUserComment, movieId, userId)
	if err != nil {
		log.Warn("{GetComment} in query: " + queryGetUserComment)
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) != 1 {
		log.Warn("{GetComment} in query: " + queryGetUserComment)
		log.Error(domain.Err.ErrObj.BadInput)
		return domain.Comment{}, domain.Err.ErrObj.BadInput
	}

	out := domain.Comment{
		Imgsrc:   cast.ToString(resp[0][0]),
		Username: cast.ToString(resp[0][1]),
		UserId:   cast.IntToStr(cast.ToUint64(resp[0][2])),
		Date:     cast.ToString(resp[0][3]),
		Content:  cast.ToString(resp[0][4]),
		Type:     cast.ToString(resp[0][5]),
		Rating:   "",
	}

	// check if this user has `rating`
	tmp, err := mr.dbm.Query(queryGetRatingCount, out.UserId)
	if err != nil {
		log.Warn("{GetComment} in query: " + queryGetRatingCount)
		log.Error(err)
		return domain.Comment{}, domain.Err.ErrObj.InternalServer
	}

	if cast.ToUint64(tmp[0][0]) == 1 {
		out.Rating = cast.IntToStr(cast.ToUint64(resp[0][6]))
	}

	return out, nil
}
