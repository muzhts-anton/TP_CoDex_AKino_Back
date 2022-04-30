package domain

// TODO: define json names
type Announced struct {
	Id            string `json:""`
	Poster        string `json:""`
	Title         string `json:""`
	TitleOriginal string `json:""`
	Info          string `json:""`
	Description   string `json:""`
	Trailer       string `json:""`
	Releasedate   string `json:""`
	Country       string `json:""`
	Director      string `json:""`
}

type AnnouncedBasic struct {
	Id          string `json:"ID"`
	Poster      string `json:"poster"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type AnnouncedUsecase interface {
	GetMovies() ([]AnnouncedBasic, error)
	// GetMovie(id uint64) (Announced, error)
}

type AnnouncedRepository interface {
	GetMovies() ([]AnnouncedBasic, error)
	// GetMovie(id uint64) (Announced, error)
}
