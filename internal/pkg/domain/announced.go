package domain

// TODO: define json names
type Announced struct {
	Id            string `json:"ID"`
	Poster        string `json:"poster"`
	Title         string `json:"title"`
	TitleOriginal string `json:"originalTitle"`
	Info          string `json:"info"`
	Description   string `json:"description"`
	Trailer       string `json:"trailerHref"`
	Releasedate   string `json:"releasedate"`
	Country       string `json:"country"`
	Director      string `json:"director"`
	Actors        []Cast  `json:"cast"`
	Genres        []GenreInMovie `json:"genres"`
}

type AnnouncedBasic struct {
	Id          string `json:"ID"`
	Poster      string `json:"poster"`
	Title       string `json:"title"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type AnnouncedSummary struct {
	Href   string `json:"href"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
}

type AnnouncedResponse struct {
	Announced   Announced          `json:"Announced"`
	Related     []AnnouncedSummary `json:"related"`
}

type AnnouncedUsecase interface {
	GetMovies() ([]AnnouncedBasic, error)
	GetMovie(id uint64) (Announced, error)
	GetRelated(id uint64) ([]AnnouncedSummary, error)
}

type AnnouncedRepository interface {
	GetMovies() ([]AnnouncedBasic, error)
	GetMovie(id uint64) (Announced, error)
	GetRelated(id uint64) ([]AnnouncedSummary, error)

}
