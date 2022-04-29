package serrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbSearchRepository struct {
	dbm *database.DBManager
}

func InitSerRep(manager *database.DBManager) domain.SearchRepository {
	return &dbSearchRepository{
		dbm: manager,
	}
}

func (cr *dbSearchRepository) SearchMovies(tag string) (domain.SearchMoviesResp, error) {
	resp, err := cr.dbm.Query(queryGetMovies, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchMovies} in query: " + queryGetMovies)
		log.Error(err)
		return domain.SearchMoviesResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchMoviesResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	movs := make([]domain.MovieBasic, 0)
	for i := range resp {
		movs = append(movs, domain.MovieBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Rating:      cast.FlToStr(cast.ToFloat64(resp[i][3])),
			Info:        cast.ToString(resp[i][4]),
			Description: cast.ToString(resp[i][5]),
		})
	}

	return domain.SearchMoviesResp{
		Empty: false,
		Data:  movs,
	}, nil
}

func (cr *dbSearchRepository) SearchActors(tag string) (domain.SearchActorsResp, error) {
	resp, err := cr.dbm.Query(queryGetActors, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchActors} in query: " + queryGetActors)
		log.Error(err)
		return domain.SearchActorsResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchActorsResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	acts := make([]domain.ActorBasic, 0)
	for i := range resp {
		acts = append(acts, domain.ActorBasic{
			Href:   "/actors/" + cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster: cast.ToString(resp[i][1]),
			Name:   cast.ToString(resp[i][2]),
		})
	}

	return domain.SearchActorsResp{
		Empty: false,
		Data:  acts,
	}, nil
}

func (cr *dbSearchRepository) SearchGenres(tag string) (domain.SearchGenresResp, error) {
	resp, err := cr.dbm.Query(queryGetGenres, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchGenres} in query: " + queryGetGenres)
		log.Error(err)
		return domain.SearchGenresResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchGenresResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	gens := make([]string, 0)
	for i := range resp {
		gens = append(gens, cast.ToString(resp[i][0]))
	}

	return domain.SearchGenresResp{
		Empty: false,
		Data:  gens,
	}, nil
}

func (cr *dbSearchRepository) SearchAnnounced(tag string) (domain.SearchAnnouncedResp, error) {
	resp, err := cr.dbm.Query(queryGetAnnounced, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchAnnounced} in query: " + queryGetAnnounced)
		log.Error(err)
		return domain.SearchAnnouncedResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchAnnouncedResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	anns := make([]domain.AnnouncedBasic, 0)
	for i := range resp {
		anns = append(anns, domain.AnnouncedBasic{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Poster:      cast.ToString(resp[i][1]),
			Title:       cast.ToString(resp[i][2]),
			Releasedate: cast.TimeToStr(cast.ToTime(resp[i][3]), false),
			Info:        cast.ToString(resp[i][4]),
			Description: cast.ToString(resp[i][5]),
		})
	}

	return domain.SearchAnnouncedResp{
		Empty: false,
		Data:  anns,
	}, nil
}

func (cr *dbSearchRepository) SearchBookmarks(tag string) (domain.SearchBookmarksResp, error) {
	resp, err := cr.dbm.Query(queryGetBookmarks, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchBookmarks} in query: " + queryGetBookmarks)
		log.Error(err)
		return domain.SearchBookmarksResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchBookmarksResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	bkms := make([]domain.Bookmark, 0)
	for i := range resp {
		bkms = append(bkms, domain.Bookmark{
			Id:          cast.IntToStr(cast.ToUint64(resp[i][0])),
			Description: cast.ToString(resp[i][1]),
			Imgsrc:      cast.ToString(resp[i][2]),
		})
	}

	return domain.SearchBookmarksResp{
		Empty: false,
		Data:  bkms,
	}, nil
}

func (cr *dbSearchRepository) SearchUsers(tag string) (domain.SearchUsersResp, error) {
	resp, err := cr.dbm.Query(queryGetUsers, "%"+tag+"%")
	if err != nil {
		log.Warn("{SearchUsers} in query: " + queryGetUsers)
		log.Error(err)
		return domain.SearchUsersResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchUsersResp{
			Empty: true,
			Data:  nil,
		}, nil
	}

	usrs := make([]domain.UserPublicInfo, 0)
	for i := range resp {
		usrs = append(usrs, domain.UserPublicInfo{
			Id:       cast.ToUint64(resp[i][0]),
			Username: cast.ToString(resp[i][1]),
			Imgsrc:   cast.ToString(resp[i][2]),
		})
	}

	return domain.SearchUsersResp{
		Empty: false,
		Data:  usrs,
	}, nil
}
