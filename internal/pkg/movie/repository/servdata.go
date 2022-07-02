package movrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
	"strings"
)

type dbMovieRepository struct {
	dbm *database.DBManager
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func addSpaces(budget string) string {
	var sb strings.Builder
	budgetRune := []rune(budget)
	budgetLength := len(budgetRune)
	if(budgetLength > 1){
		counter := 0
		sb.WriteRune(budgetRune[budgetLength - 1])
		sb.WriteRune(budgetRune[budgetLength - 2])
		for i := budgetLength - 3; i >= 0; i-- {
			sb.WriteRune(budgetRune[i])
			counter++
			if counter%3 == 0 {
				sb.WriteString(" ")
			}
		}
		return reverse(sb.String())
	}else{
		return budget
	}
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
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Movie{}, domain.Err.ErrObj.SmallDb
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
		Motto:         cast.ToString(row[10]),
		Director:      cast.ToString(row[11]),
		Budget:        addSpaces(cast.ToString(row[12])),
		Gross:         addSpaces(cast.ToString(row[13])),
		Duration:      cast.ToString(row[14]),
		Actors:        []domain.Cast{},
		Genres:        []domain.GenreInMovie{},
	}
	log.Info("out.Budget" + out.Budget)
	log.Info("out.Gross" + out.Gross)

	resp, err = mr.dbm.Query(queryGetMovieCast, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetMovieCast)
		log.Error(err)
		return domain.Movie{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovie} no cast o_0")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Movie{}, domain.Err.ErrObj.SmallDb
	}

	actors := make([]domain.Cast, 0)
	for i := range resp {
		actors = append(actors, domain.Cast{
			Name: cast.ToString(resp[i][0]),
			Href: "/actors/" + cast.IntToStr(cast.ToUint64(resp[i][1])),
		})
	}

	out.Actors = actors

	resp, err = mr.dbm.Query(queryGetMovieGenres, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetMovieGenres)
		log.Error(err)
		return domain.Movie{}, domain.Err.ErrObj.InternalServer
	}
	// if len(resp) == 0 {
	// 	log.Warn("{GetMovie} no genres")
	// 	log.Error(domain.Err.ErrObj.SmallDb)
	// 	return domain.Movie{}, domain.Err.ErrObj.SmallDb
	// }
	genres := make([]domain.GenreInMovie, 0)
	if len(resp) != 0 {
		for i := range resp {
			genres = append(genres, domain.GenreInMovie{
				Href:  "/genres/" + cast.ToString(resp[i][0]),
				Title: cast.ToString(resp[i][1]),
			})
		}

	}
	out.Genres = genres

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
			Date:     cast.TimeToStr(cast.ToTime(resp[i][3]), false),
			Content:  cast.ToString(resp[i][4]),
			Type:     cast.ToString(resp[i][5]),
			Rating:   "",
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

func (mr *dbMovieRepository) GetCollectionsInfo(userId, movieId uint64) ([]domain.CollectionInfo, error) {
	resp, err := mr.dbm.Query(queryGetPlaylists, userId)
	if err != nil {
		log.Warn("{GetCollectionsInfo} in query: " + queryGetPlaylists)
		log.Error(err)
		return []domain.CollectionInfo{}, domain.Err.ErrObj.InternalServer
	}

	var CollectionsInfo []domain.CollectionInfo

	for i := range resp {
		CollectionInfo := domain.CollectionInfo{
			Collection: cast.ToString(resp[i][0]),
			BookmarkId: cast.ToUint64(resp[i][1]),
		}

		tmp, err := mr.dbm.Query(queryGetFilmAvailability, CollectionInfo.BookmarkId, movieId)
		if err != nil {
			log.Warn("{GetCollectionsInfo} in query: " + queryGetFilmAvailability)
			log.Error(err)
			return nil, domain.Err.ErrObj.InternalServer
		}

		if cast.ToUint64(tmp[0][0]) == 1 {
			CollectionInfo.HasMovie = true
		}

		CollectionsInfo = append(CollectionsInfo, CollectionInfo)
	}

	return CollectionsInfo, nil
}
