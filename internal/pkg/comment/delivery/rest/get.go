package comdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/sanitizer"

	"encoding/json"
	"net/http"
	"strconv"
)

func (handler *CommentHandler) PostComment(w http.ResponseWriter, r *http.Request) {
	type commentReq struct {
		MovieId string `json:"movieId"`
		UserId  string `json:"userId"`
		Content string `json:"reviewText"`
		Type    string `json:"reviewType"` //int {1 2 3} {defaul: 2}
	}

	defer r.Body.Close()
	commentreq := new(commentReq)
	err := json.NewDecoder(r.Body).Decode(&commentreq)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
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

	comm, err := handler.CommentUsecase.PostComment(movieId, userId, commentreq.Content, commenttype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type commentsResp struct {
		Comment domain.Comment `json:"review"`
	}

	out, err := json.Marshal(commentsResp{
		Comment: comm,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
