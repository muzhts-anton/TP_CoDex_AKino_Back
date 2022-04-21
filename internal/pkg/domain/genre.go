package domain

type GenresUsecase interface {
	GetMovies(genre string) ([]MovieBasic, error)
}

type GenresRepository interface {
	GetMovies(genre string) ([]MovieBasic, error)
}
