package domain

type Bookmark struct {
	Id          string `json:"ID"`
	Description string `json:"description"`
	Imgsrc      string `json:"imgSrc"`
	Public      bool   `json:"public"`
}

type BookmarkWithMovies struct {
	Id          string `json:"ID"`
	Description string `json:"description"`
	Imgsrc      string `json:"imgSrc"`
}

type BookmarkResp struct {
	UserId    uint64     `json:"ID"`
	Bookmarks []Bookmark `json:"bookmarksList"`
}
