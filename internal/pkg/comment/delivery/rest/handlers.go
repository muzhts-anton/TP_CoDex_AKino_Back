package comdelivery

import (
	"codex/internal/pkg/comment/delivery/grpc"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentUsecase grpc.PosterClient
}

func SetComHandlers(router *mux.Router, pc grpc.PosterClient) {
	handler := &CommentHandler{
		CommentUsecase: pc,
	}

	router.HandleFunc(postCommentUrl, handler.PostComment).Methods("POST", "OPTIONS")
}
