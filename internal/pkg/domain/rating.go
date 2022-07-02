package domain

type RatingRepository interface {
	PostRating(movieId uint64, userId uint64, rating int) (float64, error)
}

type RatingUsecase interface {
	PostRating(movieId uint64, userId uint64, rating int) (float64, error)
}
