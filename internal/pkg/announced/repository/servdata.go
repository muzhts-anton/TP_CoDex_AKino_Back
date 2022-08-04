package annrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
	// "time"

	"fmt"
)

const (
	title       = "Самые долгожданные премьеры"
	description = "Премьеры были собраны лучшими кинокритиками, чтобы вы не упустили самое интересное!"
)

type dbAnnouncedRepository struct {
	dbm *database.DBManager
}

func InitAnnRep(manager *database.DBManager) domain.AnnouncedRepository {
	return &dbAnnouncedRepository{
		dbm: manager,
	}
}

func intToStringMonth(number string) (string, error) {
	switch number {
	case "1":
		return "января", nil
	case "2":
		return "февраля", nil
	case "3":
		return "марта", nil
	case "4":
		return "апреля", nil
	case "5":
		return "мая", nil
	case "6":
		return "июня", nil
	case "7":
		return "июля", nil
	case "8":
		return "августа", nil
	case "9":
		return "сентября", nil
	case "10":
		return "октября", nil
	case "11":
		return "ноября", nil
	case "12":
		return "декабря", nil
	}
	log.Warn("{intToStringMonth}")
	return "", domain.Err.ErrObj.InternalServer
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
		PremierMonth, err := intToStringMonth(cast.ToString(resp[i][4]))
		if err != nil {
			log.Warn("{GetMovies}")
			log.Error(err)
			return domain.AnnouncedBasicResponse{}, domain.Err.ErrObj.InternalServer
		}
		// if cast.ToDate((resp[i][6])).After(time.Now()){
			movies = append(movies, domain.AnnouncedBasic{
				Id:            cast.IntToStr(cast.ToUint64(resp[i][0])),
				Poster:        cast.ToString(resp[i][1]),
				Title:         cast.ToString(resp[i][2]),
				OriginalTitle: cast.ToString(resp[i][3]),
				PremierMonth:  PremierMonth,
				PremierDay:    cast.ToString(resp[i][5]),
			})
		// }
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
			Href:  "/genres/" + cast.ToString(resp[i][0]),
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

func (fr *dbAnnouncedRepository) GetAnnouncedByMonthYear(month int, year int) (domain.AnnouncedList, error) {
	log.Info(fmt.Sprintf("Month = %d Year = %d", month, year))

	resp, err := fr.dbm.Query(queryCountAnnouncedByMonthYear, month, year)
	if err != nil {
		return domain.AnnouncedList{}, domain.Err.ErrObj.InternalServer
	}
	announcedQuantity := int(cast.ToUint64(resp[0][0]))
	log.Info(fmt.Sprintf("announcedQuantity = %d", announcedQuantity))
	// if skip >= dbSize && skip != 0 {
	// 	return domain.FilmList{}, customErrors.ErrorSkip
	// }
	// moreAvailable := skip+limit < dbSize

	resp, err = fr.dbm.Query(queryGetAnnouncedsByMonthYear, month, year)
	if err != nil {
		return domain.AnnouncedList{}, err
	}

	bufferAnnounced := make([]domain.Announced, 0)
	for i := range resp {
		announced := domain.Announced{
			Id:             cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:         cast.ToString(resp[i][1]),
			Title:          cast.ToString(resp[i][2]),
			TitleOriginal:  cast.ToString(resp[i][3]),
			Info:           cast.ToString(resp[i][4]),
			Description:    cast.ToString(resp[i][5]),
			Trailer:        cast.ToString(resp[i][6]),
			// Releasedate:    cast.ToString(resp[i][7]),
			Country:        cast.ToString(resp[i][8]),
			Director:       cast.ToString(resp[i][9]),
		}
		dateString, err := cast.DateToStringUnderscore(resp[i][7])
		if err != nil {
			return domain.AnnouncedList{}, err
		}
		announced.Releasedate = dateString

		// if cast.ToDate((resp[i][7])).After(time.Now()){
			bufferAnnounced = append(bufferAnnounced, announced)
		// }
	}
	announcedList := domain.AnnouncedList{
		AnnouncedList:  bufferAnnounced,
		AnnouncedTotal: announcedQuantity,
	}
	return announcedList, nil
}

