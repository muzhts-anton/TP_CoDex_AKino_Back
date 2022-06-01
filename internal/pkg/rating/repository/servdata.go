package ratrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

	"fmt"
	"math"
)

type dbRatingRepository struct {
	dbm *database.DBManager
}

func InitRatRep(manager *database.DBManager) domain.RatingRepository {
	return &dbRatingRepository{
		dbm: manager,
	}
}

func (mr *dbRatingRepository) PostRating(movieId uint64, userId uint64, rating int) (float64, error) {
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
		// log.Info("oldVotesnum = " + string(oldVotesnum))
		newRating = (oldRating*float64(oldVotesnum) + float64(rating)) / float64(oldVotesnum+1)
		log.Info(fmt.Sprintf("oldR: %v\n oldV: %v\n, rating: %v", oldRating, oldVotesnum, rating))
	}

	newRating = math.Round(newRating*100) / 100

	if isOldRating == 0 {
		_, err = mr.dbm.Query(queryIncrementVotesnum, movieId)
		if err != nil {
			log.Info("INCREMENT VOTESNUM")
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
		_, err = mr.dbm.Query(queryChangeRating, rating, userId, movieId)
		if err != nil {
			log.Warn("{PostRating} in query: " + queryChangeRating)
			log.Error(err)
			return 0.0, domain.Err.ErrObj.InternalServer
		}
	}

	return newRating, nil
}
