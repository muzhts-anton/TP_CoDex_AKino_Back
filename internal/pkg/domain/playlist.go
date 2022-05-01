package domain

const (
	BasePlaylistPicture = "/profile.svg"
	minTitleLength      = 4
	maxTitleLength      = 50
)

type PlaylistRequest struct {
	Title  string `json:"title"`
	UserId uint64 `json:"userId"`
	Public bool   `json:"public"`
}

func (pr PlaylistRequest) TitleIsValid() (isValid bool) {
	if len(pr.Title) < minTitleLength && len(pr.Title) > maxTitleLength {
		return false
	}
	return true
}

type PlaylistResponse struct {
	ID    string `json:"ID"`
	Title string `json:"title"`
}

type MovieInPlaylist struct {
	// UserId     uint64 `json:"userId"`
	MovieId    uint64 `json:"movieId"`
	PlaylistId uint64 `json:"bookmarkId"`
}

type Playlist struct {
	Id          uint64 `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Poster      string `json:"poster"`
	Public      bool   `json:"public"`
}

// type DeleteMovieInfo struct {
// 	MovieId    uint64 `json:"movieId"`
// 	PlaylistId uint64 `json:"bookmarkId"`
// }

type DeletePlaylistInfo struct {
	PlaylistId uint64 `json:"bookmarkId"`
}

type Plarepository interface {
	CreatePlaylist(playlist PlaylistRequest) (PlaylistResponse, error)
	AddMovie(addMovieInfo MovieInPlaylist) error
	DeleteMovie(movieDeleteInfo MovieInPlaylist) error
	DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
	PlaylistAlreadyExist(playlist PlaylistRequest) (bool, error)
}

type PlaylistUsecase interface {
	CreatePlaylist(playlist PlaylistRequest) (PlaylistResponse, error)
	AddMovie(addMovieInfo MovieInPlaylist) error
	DeleteMovie(MovieInPlaylist MovieInPlaylist) error
	DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
}
