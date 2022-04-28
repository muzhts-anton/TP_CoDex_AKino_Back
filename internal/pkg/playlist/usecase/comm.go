package playlistusecase

import (
	"codex/internal/pkg/domain"
	"strings"
)

type playlistUsecase struct {
	playlistRepo domain.PlaylistRepository
}

func trimTitle(title *string) {
	*title = strings.Trim(*title, " ")
}

func InitPlaylistUsc(pr domain.PlaylistRepository) domain.PlaylistUsecase {
	return &playlistUsecase{
		playlistRepo: pr,
	}
}

func (pu playlistUsecase) CreatePlaylist(playlistData domain.PlaylistRequest) (domain.PlaylistResponse, error) {
	trimTitle(&playlistData.Title)

	alreadyExist, err := pu.playlistRepo.PlaylistAlreadyExist(playlistData)
	if err != nil {
		return domain.PlaylistResponse{}, err
	}	

	if (alreadyExist){
		return domain.PlaylistResponse{}, domain.Err.ErrObj.PlaylistExist
	}

	if !playlistData.TitleIsValid(){
		return domain.PlaylistResponse{}, domain.Err.ErrObj.InvalidTitle
	}

	playlistResponse, err := pu.playlistRepo.CreatePlaylist(playlistData)
	if err != nil {
		return domain.PlaylistResponse{}, err
	}	

	return playlistResponse, nil
}

func (pu playlistUsecase) AddMovie(addMovieInfo domain.AddToPlaylist) (error) {
	err := pu.playlistRepo.AddMovie(addMovieInfo)
	if err != nil {
		return err
	}	
	return  nil
}

func (pu playlistUsecase) DeleteMovie(deleteMovieInfo domain.DeleteMovieInfo) (error) {
	err := pu.playlistRepo.DeleteMovie(deleteMovieInfo)
	if err != nil {
		return err
	}	
	return  nil
}

func (pu playlistUsecase) DeletePlaylist(deletePlaylistInfo domain.DeletePlaylistInfo) (error) {
	err := pu.playlistRepo.DeletePlaylist(deletePlaylistInfo)
	if err != nil {
		return err
	}	
	return  nil
}

