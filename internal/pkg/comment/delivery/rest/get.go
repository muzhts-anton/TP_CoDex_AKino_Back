package comdelivery

import (
	"codex/internal/pkg/comment/delivery/grpc"
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/sanitizer"

	"context"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
)

type commentsResp struct {
	Comment domain.Comment `json:"review"`
}

type commentReq struct {
	MovieId string `json:"movieId"`
	UserId  string `json:"userId"`
	Content string `json:"reviewText"`
	Type    string `json:"reviewType"` //int {1 2 3} {default: 2}
}

func (handler *CommentHandler) PostComment(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentreq := new(commentReq)
	err = easyjson.Unmarshal(b, commentreq)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		return
	}

	sanitizer.SanitizeComment(&commentreq.Content)

	movieId, err := strconv.ParseUint(commentreq.MovieId, 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(commentreq.UserId, 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	commenttype, err := strconv.Atoi(commentreq.Type)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	comm, err := handler.CommentUsecase.PostComment(context.Background(), &grpc.Data{
		MovieId:     movieId,
		UserId:      userId,
		Content:     commentreq.Content,
		CommentType: int32(commenttype),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(commentsResp{
		Comment: domain.Comment{
			Imgsrc:   comm.Imgsrc,
			Username: comm.Username,
			UserId:   comm.UserId,
			Rating:   comm.Rating,
			Date:     comm.Date,
			Content:  comm.Content,
			Type:     comm.Type,
		},
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
