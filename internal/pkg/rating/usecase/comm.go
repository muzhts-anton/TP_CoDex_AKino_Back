package ratusecase

import (
	"codex/internal/pkg/rating/delivery/grpc"
	"codex/internal/pkg/domain"

	"context"
)

type ratingUsecase struct {
	grpc.UnimplementedPosterServer
	ratingRepo domain.RatingRepository
}

func InitRatUsc(rr domain.RatingRepository) grpc.PosterServer {
	return &ratingUsecase{
		ratingRepo: rr,
	}
}

func (ru ratingUsecase) PostRating(ctx context.Context, data *grpc.Data) (*grpc.NewRating, error) {
	if data.GetRating() < 1 || data.GetRating() > 10 {
		return nil, domain.Err.ErrObj.InvalidRating
	}

	newRating, err := ru.ratingRepo.PostRating(data.GetMovieId(), data.GetUserId(), int(data.GetRating()))
	if err != nil {
		return nil, err
	}

	return &grpc.NewRating{Rating: newRating}, nil
}
