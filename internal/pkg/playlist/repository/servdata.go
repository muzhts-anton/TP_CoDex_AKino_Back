package plarepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"
)

type dbplarepository struct {
	dbm *database.DBManager
}

func InitPlaRep(manager *database.DBManager) domain.Plarepository {
	return &dbplarepository{
		dbm: manager,
	}
}

func (pr *dbplarepository) CreatePlaylist(playlist domain.PlaylistRequest) (domain.PlaylistResponse, error) {
	resp, err := pr.dbm.Query(queryCreatePlaylist, playlist.Title, playlist.Public)
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylist)
		log.Error(err)
		return domain.PlaylistResponse{}, err
	}

	_, err = pr.dbm.Query(queryCreatePlaylistUser, playlist.UserId, cast.ToUint64(resp[0][0]))
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylistUser)
		log.Error(err)
		return domain.PlaylistResponse{}, err
	}

	return domain.PlaylistResponse{
		ID:     cast.IntToStr(cast.ToUint64(resp[0][0])),
		Title:  cast.ToString(resp[0][1]),
		ImgSrc: cast.ToString(resp[0][2]),
		Public: cast.ToBool(resp[0][3]),
	}, nil
}

func (pr *dbplarepository) PlaylistAlreadyExist(playlist domain.PlaylistRequest) (bool, error) {
	resp, err := pr.dbm.Query(queryPlaylistExist, playlist.UserId, playlist.Title)
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylist)
		log.Error(err)
		return false, err
	}

	if cast.ToUint64(resp[0][0]) != 0 {
		return true, nil
	}
	return false, nil
}

func (pr *dbplarepository) AddMovie(addMovieInfo domain.MovieInPlaylist) error {
	_, err := pr.dbm.Query(queryAddMovie, addMovieInfo.PlaylistId, addMovieInfo.MovieId)
	if err != nil {
		log.Warn("{AddMovie} in query: " + queryAddMovie)
		log.Error(err)
		return err
	}

	return nil
}

func (pr *dbplarepository) DeleteMovie(MovieInPlaylist domain.MovieInPlaylist) error {
	_, err := pr.dbm.Query(queryDeleteMovie, MovieInPlaylist.PlaylistId, MovieInPlaylist.MovieId)
	if err != nil {
		log.Warn("{DeleteMovie} in query: " + queryDeleteMovie)
		log.Error(err)
		return err
	}

	return nil
}

func (pr *dbplarepository) DeletePlaylist(deletePlaylistInfo domain.DeletePlaylistInfo) error {
	_, err := pr.dbm.Query(queryDeletePlaylist, deletePlaylistInfo.PlaylistId)
	if err != nil {
		log.Warn("{DeletePlaylist} in query: " + queryDeletePlaylist)
		log.Error(err)
		return err
	}

	return nil
}

func (pr *dbplarepository) AlterPlaylistPublic(alterPlaylistPublicInfo domain.AlterPlaylistPublicInfo) error {
	_, err := pr.dbm.Query(queryAlterPlaylistPublic, alterPlaylistPublicInfo.PlaylistId, alterPlaylistPublicInfo.Public)
	if err != nil {
		log.Warn("{AlterPlaylistPublic} in query: " + queryAlterPlaylistPublic)
		log.Error(err)
		return err
	}

	return nil
}

func (pr *dbplarepository) AlterPlaylistTitle(alterPlaylistTitleInfo domain.AlterPlaylistTitleInfo) error {
	_, err := pr.dbm.Query(queryAlterPlaylistTitle, alterPlaylistTitleInfo.PlaylistId, alterPlaylistTitleInfo.NewTitle)
	if err != nil {
		log.Warn("{AlterPlaylistTitle} in query: " + queryAlterPlaylistTitle)
		log.Error(err)
		return err
	}

	return nil
}
