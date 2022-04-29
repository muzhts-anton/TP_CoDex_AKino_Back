package domain

type Genre struct{
	Href   string `json:"href"`
	Imgsrc string `json:"imgsrc"`
}

type GenresUsecase interface {
	GetMovies(genre string) ([]MovieBasic, error)
	GetGenres() ([]Genre, error)
}

type GenresRepository interface {
	GetMovies(genre string) ([]MovieBasic, error)
	GetGenres() ([]Genre, error)
}
