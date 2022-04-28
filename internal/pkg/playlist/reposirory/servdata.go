package playlistrepository

import (
	"codex/internal/pkg/database"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"
	"codex/internal/pkg/utils/log"

)

type dbPlaylistRepository struct {
	dbm *database.DBManager
}

func InitPlaylistRep(manager *database.DBManager) domain.PlaylistRepository {
	return &dbPlaylistRepository{
		dbm: manager,
	}
}

func (pr *dbPlaylistRepository) CreatePlaylist(playlist domain.PlaylistRequest) (domain.PlaylistResponse, error) {
	resp, err := pr.dbm.Query(queryCreatePlaylist, playlist.Title, playlist.Public)
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylist)
		log.Error(err)
		return domain.PlaylistResponse{}, err
	}
	var playlistResponse domain.PlaylistResponse
	playlistResponse.ID = cast.ToString(resp[0][0])
	
	resp, err = pr.dbm.Query(queryCreatePlaylistUser, playlist.UserId,  cast.ToUint64(resp[0][0]))
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylistUser)
		log.Error(err)
		return domain.PlaylistResponse{}, err
	}

	return domain.PlaylistResponse{}, nil
}

func (pr *dbPlaylistRepository) PlaylistAlreadyExist(playlist domain.PlaylistRequest) (bool, error) {
	resp, err := pr.dbm.Query(queryPlaylistExist, playlist.UserId ,playlist.Title)
	if err != nil {
		log.Warn("{CreatePlaylist} in query: " + queryCreatePlaylist)
		log.Error(err)
		return false, err
	}

	if(cast.ToUint64(resp[0][0]) == 0){
		return true, nil	
	}
	return false, nil
}

func (pr *dbPlaylistRepository) AddMovie(addMovieInfo domain.AddToPlaylist) (error) {
	_, err := pr.dbm.Query(queryAddMovie, addMovieInfo.PlaylistId, addMovieInfo.MovieId)
	if err != nil {
		log.Warn("{AddMovie} in query: " + queryAddMovie)
		log.Error(err)
		return err
	}

	return nil
}

func (pr *dbPlaylistRepository) DeleteMovie(deleteMovieInfo domain.DeleteMovieInfo) (error) {
	_, err := pr.dbm.Query(queryDeleteMovie, deleteMovieInfo.PlaylistId, deleteMovieInfo.MovieId)
	if err != nil {
		log.Warn("{DeleteMovie} in query: " + queryDeleteMovie)
		log.Error(err)
		return err
	}
	
	return nil
}

func (pr *dbPlaylistRepository) DeletePlaylist(deletePlaylistInfo domain.DeletePlaylistInfo) (error) {
	_, err := pr.dbm.Query(queryDeletePlaylist, deletePlaylistInfo.PlaylistId)
	if err != nil {
		log.Warn("{DeletePlaylist} in query: " + queryDeletePlaylist)
		log.Error(err)
		return err
	}
	
	return nil
}

// router.HandleFunc(createPlaylistUrl, handler.CreatePlaylist).Methods("POST", "OPTIONS")
// router.HandleFunc(addMovieUrl, handler.AddMovie).Methods("POST", "OPTIONS")
// router.HandleFunc(deleteMovieUrl, handler.DeleteMovie).Methods("POST", "OPTIONS")
// router.HandleFunc(deletePlaylistUrl, handler.DeletePlaylist).Methods("POST", "OPTIONS")
