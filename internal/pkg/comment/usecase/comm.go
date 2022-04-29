package comusecase

import (
	"codex/internal/pkg/comment/delivery/grpc"
	"codex/internal/pkg/domain"

	"context"
)

type commentUsecase struct {
	grpc.UnimplementedPosterServer
	commentRepo domain.CommentRepository
}

func InitComUsc(cr domain.CommentRepository) grpc.PosterServer {
	return &commentUsecase{
		commentRepo: cr,
	}
}

func (cu commentUsecase) PostComment(ctx context.Context, in *grpc.Data) (*grpc.Comment, error) {
	var comtype string
	if in.CommentType == 1 {
		comtype = "good"
	} else if in.CommentType == 2 {
		comtype = "neutral"
	} else if in.CommentType == 3 {
		comtype = "bad"
	} else {
		return nil, domain.Err.ErrObj.InvalidCommentType
	}

	out, err := cu.commentRepo.PostComment(in.MovieId, in.UserId, in.Content, comtype)

	return &grpc.Comment{
		Imgsrc:   out.Imgsrc,
		Username: out.Username,
		UserId:   out.UserId,
		Rating:   out.Rating,
		Date:     out.Date,
		Content:  out.Content,
		Type:     out.Type,
	}, err
}
