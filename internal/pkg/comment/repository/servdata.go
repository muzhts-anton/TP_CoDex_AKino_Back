package comrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

	"time"
)

type dbCommentRepository struct {
	dbm *database.DBManager
}

func InitComRep(manager *database.DBManager) domain.CommentRepository {
	return &dbCommentRepository{
		dbm: manager,
	}
}

func (mr *dbCommentRepository) PostComment(movieId uint64, userId uint64, content string, comtype string) (domain.Comment, error) {
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
	log.Info("content = " + content)
	_, err = mr.dbm.Query(queryPostComment, userId, movieId, time.Now().Format("2006.01.02 15:04:05"), comtype, content)
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

func (mr *dbCommentRepository) GetComment(userId, movieId uint64) (domain.Comment, error) {
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
		Date:     cast.TimeToStr(cast.ToTime(resp[0][3]), true),
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
