package domain

type Profile struct {
	Id            string `json:"ID"`
	avatarSrc     string `json:"avatarSrc"`
	name	      string `json:"name"`
	email	      string `json:"email"`
}

type Bookmark struct {
	description   string `json:"description"`
	imgSrc   string `json:"imgSrc"`
	number        string `json:"number"`
	// movies
}

type BookmarksSummary struct {
	Id            string `json:"ID"`
	bookmarksList []Bookmark
}

type Review struct {
	film_name     string `json:"film_name"`
	typee         string `json:"type"`
	rating        string `json:"rating"`
}

type ReviewsSummary struct {
	Id            string `json:"ID"`
	reviewsList   []Review
}



type ProfileRepository interface {
	GetProfile(id uint64) (Profile, error)
	GetBookmarks(id uint64) ([]BookmarksSummary, error)
	GetReviews(id uint64) ([]Reviews, error)
}

type ProfileUsecase interface {
	GetProfile(id uint64) (Profile, error)
	GetBookmarks(id uint64) ([]Bookmarks, error)
	GetReviews(id uint64) ([]Reviews, error)
}
