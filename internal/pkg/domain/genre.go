package domain

type Genre struct{
	Href        string `json:"href"`
	Imgsrc      string `json:"imgsrc"`
}

type GenreWithMovies struct{
	Href        string       `json:"href"`
	Imgsrc      string       `json:"imgsrc"`
	Description string       `json:"description"`
	Title       string       `json:"title"`
	MovieList   []MovieBasic `json:"movielist"`
}

type GenresUsecase interface {
	GetGenre(genre string) (GenreWithMovies, error)
	GetGenres() ([]Genre, error)
}

type GenresRepository interface {
	GetGenre(genre string) (GenreWithMovies, error)
	GetGenres() ([]Genre, error)
}
