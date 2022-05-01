package annrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

const(
	title = "ЭТО ЛУЧШТЕ ПРЕМЬЕРЫ!!!11!Ы!!"
	description = "Понятия не имею, зачем описание в премьерах, но фронтовики хотят, и ментор по интерфейсам одобрил"
)

type dbAnnouncedRepository struct {
	dbm *database.DBManager
}

func InitAnnRep(manager *database.DBManager) domain.AnnouncedRepository {
	return &dbAnnouncedRepository{
		dbm: manager,
	}
}

func (ar *dbAnnouncedRepository) GetMovies() (domain.AnnouncedBasicResponse, error) {
	resp, err := ar.dbm.Query(queryGetMovies)
	if err != nil {
		log.Warn("{GetMovies} in query: " + queryGetMovies)
		log.Error(err)
		return domain.AnnouncedBasicResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.AnnouncedBasicResponse{}, domain.Err.ErrObj.SmallDb
	}

	movies := make([]domain.AnnouncedBasic, 0)
	for i := range resp {
		movies = append(movies, domain.AnnouncedBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Info:        "Дата премьеры: " + cast.TimeToStr(cast.ToTime(resp[i][3]), false) + ". Осталось " +  cast.ToString(resp[i][4]) + " дня.",
			Description: cast.ToString(resp[i][5]),
		})
	}
	var movieResponse domain.AnnouncedBasicResponse
	movieResponse.MovieList = movies
	movieResponse.Title = title
	movieResponse.Description = description
	return movieResponse, nil
}

func (ar *dbAnnouncedRepository) GetMovie(id uint64) (domain.Announced, error) {
	resp, err := ar.dbm.Query(queryGetAnnounced, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetAnnounced)
		log.Error(err)
		return domain.Announced{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetMovie}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Announced{}, domain.Err.ErrObj.SmallDb
	}

	row := resp[0]
	out := domain.Announced{
		Id:            cast.IntToStr(cast.ToUint64(row[0])),
		Poster:        cast.ToString(row[1]),
		Title:         cast.ToString(row[2]),
		TitleOriginal: cast.ToString(row[3]),
		Info:          cast.ToString(row[4]),
		Description:   cast.ToString(row[5]),
		Trailer:       cast.ToString(row[6]),
		Releasedate:   cast.ToString(row[7]),
		Country:       cast.ToString(row[8]),
		Director:      cast.ToString(row[9]),
	}

	resp, err = ar.dbm.Query(queryGetAnnouncedCast, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetAnnouncedCast)
		log.Error(err)
		return domain.Announced{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovie} no cast o_0")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Announced{}, domain.Err.ErrObj.SmallDb
	}

	actors := make([]domain.Cast, 0)
	for i := range resp {
		actors = append(actors, domain.Cast{
			Name: cast.ToString(resp[i][0]),
			Href: "/actors/" + cast.IntToStr(cast.ToUint64(resp[i][1])),
		})
	}

	out.Actors = actors



	resp, err = ar.dbm.Query(queryGetAnnouncedGenres, id)
	if err != nil {
		log.Warn("{GetMovie} in query: " + queryGetAnnouncedGenres)
		log.Error(err)
		return domain.Announced{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn("{GetMovie} no genres")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Announced{}, domain.Err.ErrObj.SmallDb
	}

	genres := make([]domain.GenreInMovie, 0)
	for i := range resp {
		genres = append(genres, domain.GenreInMovie{
			Href: "/geners/" + cast.ToString(resp[i][0]),
			Title: cast.ToString(resp[i][1]),
		})
	}

	out.Genres = genres

	return out, nil
}

func (ar *dbAnnouncedRepository) GetRelated(id uint64) ([]domain.AnnouncedSummary, error) {
	resp, err := ar.dbm.Query(queryGetRelated, id)
	if err != nil {
		log.Warn("{GetRelated} in query: " + queryGetRelated)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return []domain.AnnouncedSummary{}, nil
	}

	out := make([]domain.AnnouncedSummary, 0)
	for i := range resp {
		out = append(out, domain.AnnouncedSummary{
			Href:   "/announced/" + cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster: cast.ToString(resp[i][1]),
			Title:  cast.ToString(resp[i][2]),
		})
	}

	return out, nil
}
