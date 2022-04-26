package comusecase

import (
	"codex/internal/pkg/domain"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func InitComUsc(cr domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: cr,
	}
}

func (cu commentUsecase) PostComment(movieId uint64, userId uint64, content string, commenttype int) (domain.Comment, error) {
	var comtype string
	if commenttype == 1 {
		comtype = "good"
	} else if commenttype == 2 {
		comtype = "neutral"
	} else if commenttype == 3 {
		comtype = "bad"
	} else {
		return domain.Comment{}, domain.Err.ErrObj.InvalidCommentType
	}

	return cu.commentRepo.PostComment(movieId, userId, content, comtype)
}
