package domain

// TODO: define json names
type Announced struct {
	Id            string         `json:"ID"`
	Poster        string         `json:"poster"`
	Title         string         `json:"title"`
	TitleOriginal string         `json:"originalTitle"`
	Info          string         `json:"info"`
	Description   string         `json:"description"`
	Trailer       string         `json:"trailerHref"`
	Releasedate   string         `json:"releasedate"`
	Country       string         `json:"country"`
	Director      string         `json:"director"`
	Actors        []Cast         `json:"cast"`
	Genres        []GenreInMovie `json:"genres"`
}

type AnnouncedBasic struct {
	Id            string `json:"ID"`
	Poster        string `json:"poster"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	PremierDay    string `json:"premierDay"`
	PremierMonth  string `json:"premierMonth"`
}

type AnnouncedSearch struct {
	Id          string `json:"ID"`
	Poster      string `json:"poster"`
	Title       string `json:"title"`
	Releasedate string `json:"releasedate"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type AnnouncedBasicResponse struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	MovieList   []AnnouncedBasic `json:"movielist"`
}

type AnnouncedSummary struct {
	Href   string `json:"href"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
}

type AnnouncedResponse struct {
	Announced Announced          `json:"movie"`
	Related   []AnnouncedSummary `json:"related"`
}
type AnnouncedList struct {
	AnnouncedList      []Announced `json:"announced_list"`
	AnnouncedTotal          int    `json:"announced_total"`
}

type AnnouncedUsecase interface {
	GetMovies() (AnnouncedBasicResponse, error)
	GetMovie(id uint64) (Announced, error)
	GetRelated(id uint64) ([]AnnouncedSummary, error)
}

type AnnouncedRepository interface {
	GetMovies() (AnnouncedBasicResponse, error)
	GetMovie(id uint64) (Announced, error)
	GetRelated(id uint64) ([]AnnouncedSummary, error)	
	GetAnnouncedByMonthYear(month int, year int) (AnnouncedList, error)

}
