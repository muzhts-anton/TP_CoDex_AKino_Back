package domain

type Movie struct {
	Id            string `json:"ID"`
	Poster        string `json:"poster"`
	Title         string `json:"title"`
	TitleOriginal string `json:"originalTitle"`
	Rating        string `json:"rating"`
	VotesNum      uint64 `json:"-"`
	Info          string `json:"info"`
	Description   string `json:"description"`
	Trailer       string `json:"trailerHref"`
	ReleaseYear   string `json:"year"`
	Country       string `json:"country"`
	Genre         string `json:"genre"`
	Motto         string `json:"motto"`
	Director      string `json:"director"`
	Budget        string `json:"budget"`
	Gross         string `json:"gross"`
	Duration      string `json:"duration"`
	Actors        []Cast `json:"cast"`
}

type Cast struct {
	Name string `json:"name"`
	Href string `json:"href"`
}

type MovieBasic struct {
	Id          string `json:"ID"`
	Poster      string `json:"poster"`
	Title       string `json:"title"`
	Rating      string `json:"rating"`
	Info        string `json:"info"`
	Description string `json:"description"`
}

type MovieSummary struct {
	Href   string `json:"href"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
}

type MovieResponse struct {
	Movie       Movie          `json:"movie"`
	Related     []MovieSummary `json:"related"`
	Comments    []Comment      `json:"reviews"`
	ReviewExist string         `json:"reviewex"`
	UserRating  string         `json:"userrating"`
}

type MovieRepository interface {
	GetMovie(id uint64) (Movie, error)
	GetRelated(id uint64) ([]MovieSummary, error)
	GetComments(id uint64) ([]Comment, error)
	GetReviewRating(movieId, userId uint64) (string, string, error)
	PostRating(movieId uint64, userId uint64, rating int) (float64, error)
	PostComment(movieId uint64, userId uint64, content string, comtype string) (Comment, error)
}

//mockgen -destination=../movie/usecase/mock/usecase_mock.go  -package=mock codex/internal/pkg/domain MovieUsecase
type MovieUsecase interface {
	GetMovie(id uint64) (Movie, error)
	GetRelated(id uint64) ([]MovieSummary, error)
	GetComments(id uint64) ([]Comment, error)
	GetReviewRating(movieId, userId uint64) (string, string, error)
	PostRating(movieId uint64, userId uint64, rating int) (float64, error)
	PostComment(movieId uint64, userId uint64, content string, commenttype int) (Comment, error)
}
