package serrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	_ "codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbSearchRepository struct {
	dbm *database.DBManager
}

func InitAnnRep(manager *database.DBManager) domain.SearchRepository {
	return &dbSearchRepository{
		dbm: manager,
	}
}

func (cr *dbSearchRepository) SearchMovies(tag string) (domain.SearchMoviesResp, error) {
	resp, err := cr.dbm.Query(queryGetMovies, tag)
	if err != nil {
		log.Warn("{SearchMovies} in query: " + queryGetMovies)
		log.Error(err)
		return domain.SearchMoviesResp{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.SearchMoviesResp{
			Empty: true,
			Data: nil,
		}, nil
	}

	movs := make([]domain.MovieBasic, 0)
	for _ = range resp {
		movs = append(movs, domain.MovieBasic{})
	}

	return domain.SearchMoviesResp{
		Empty: false,
		Data: movs,
	}, nil
}

func (cr *dbSearchRepository) SearchActors(tag string) (domain.SearchActorsResp, error) {
	return domain.SearchActorsResp{}, nil
}

func (cr *dbSearchRepository) SearchGenres(tag string) (domain.SearchGenresResp, error) {
	return domain.SearchGenresResp{}, nil
}

func (cr *dbSearchRepository) SearchAnnounced(tag string) (domain.SearchAnnouncedResp, error) {
	return domain.SearchAnnouncedResp{}, nil
}

func (cr *dbSearchRepository) SearchBookmarks(tag string) (domain.SearchBookmarksResp, error) {
	return domain.SearchBookmarksResp{}, nil
}

func (cr *dbSearchRepository) SearchUsers(tag string) (domain.SearchUsersResp, error) {
	return domain.SearchUsersResp{}, nil
}
