package domain

type SearchMoviesResp struct {
	Empty bool         `json:"isEmpty"`
	Data  []MovieBasic `json:"data"`
}

type SearchGenresResp struct {
	Empty bool     `json:"isEmpty"`
	Data  []string `json:"data"`
}

type SearchActorsResp struct {
	Empty bool         `json:"isEmpty"`
	Data  []ActorBasic `json:"data"`
}

type SearchAnnouncedResp struct {
	Empty bool              `json:"isEmpty"`
	Data  []AnnouncedSearch `json:"data"`
}

type SearchBookmarksResp struct {
	Empty bool       `json:"isEmpty"`
	Data  []Bookmark `json:"data"`
}

type SearchUsersResp struct {
	Empty bool             `json:"isEmpty"`
	Data  []UserPublicInfo `json:"data"`
}

type SearchResponse struct {
	Movies    SearchMoviesResp    `json:"movies"`
	Genres    SearchGenresResp    `json:"genres"`
	Actors    SearchActorsResp    `json:"actors"`
	Announced SearchAnnouncedResp `json:"announced"`
	Bookmarks SearchBookmarksResp `json:"bookmarks"`
	Users     SearchUsersResp     `json:"users"`
}

type SearchRepository interface {
	SearchMovies(tag string) (SearchMoviesResp, error)
	SearchGenres(tag string) (SearchGenresResp, error)
	SearchActors(tag string) (SearchActorsResp, error)
	SearchAnnounced(tag string) (SearchAnnouncedResp, error)
	SearchBookmarks(tag string) (SearchBookmarksResp, error)
	SearchUsers(tag string) (SearchUsersResp, error)
}

type SearchUsecase interface {
	Search(tag string) (SearchResponse, error)
}
