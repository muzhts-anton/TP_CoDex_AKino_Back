package comdelivery

import (
	"codex/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentUsecase domain.CommentUsecase
}

func SetComHandlers(router *mux.Router, cu domain.CommentUsecase) {
	handler := &CommentHandler{
		CommentUsecase: cu,
	}

	router.HandleFunc(postCommentUrl, handler.PostComment).Methods("POST", "OPTIONS")
}