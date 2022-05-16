package domain

const (
	BasePlaylistPicture = "/profile.svg"
	minTitleLength      = 4
	maxTitleLength      = 50
)

type PlaylistRequest struct {
	Title  string `json:"title"`
	UserId string `json:"userId"`
	Public bool   `json:"public"`
}

func (pr PlaylistRequest) TitleIsValid() (isValid bool) {
	if len(pr.Title) < minTitleLength && len(pr.Title) > maxTitleLength {
		return false
	}
	return true
}

type PlaylistResponse struct {
	ID     string `json:"ID"`
	Title  string `json:"title"`
	ImgSrc string `json:"imgSrc"`
	Public bool   `json:"public"`

}

type MovieInPlaylist struct {
	MovieId    string `json:"movieId"`
	PlaylistId string `json:"bookmarkId"`
}

type Playlist struct {
	Id          string `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImgSrc      string `json:"imgSrc"`
	Public      bool   `json:"public"`
}

type DeletePlaylistInfo struct {
	PlaylistId string `json:"bookmarkId"`
}

type AlterPlaylistPublicInfo struct {
	PlaylistId string `json:"bookmarkId"`
	Public     bool    `json:"public"`
}

type PlaylistWithMovies struct {
	Title     string       `json:"title"`
	UserId    string       `json:"userId"`
	Public    bool         `json:"public"`
	MovieList []MovieBasic `json:"movieList"`
}

type Plarepository interface {
	CreatePlaylist(playlist PlaylistRequest) (PlaylistResponse, error)
	AddMovie(addMovieInfo MovieInPlaylist) error
	DeleteMovie(movieDeleteInfo MovieInPlaylist) error
	DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
	PlaylistAlreadyExist(playlist PlaylistRequest) (bool, error)
	AlterPlaylistPublic(alterPlaylistPublicInfo AlterPlaylistPublicInfo) error

}

type PlaylistUsecase interface {
	CreatePlaylist(playlist PlaylistRequest) (PlaylistResponse, error)
	AddMovie(addMovieInfo MovieInPlaylist) error
	DeleteMovie(MovieInPlaylist MovieInPlaylist) error
	DeletePlaylist(deletePlaylistInfo DeletePlaylistInfo) error
	AlterPlaylistPublic(alterPlaylistPublicInfo AlterPlaylistPublicInfo) error
}
