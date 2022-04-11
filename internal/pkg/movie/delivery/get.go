package movdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/utils/cast"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (handler *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	movie, err := handler.MovieUsecase.GetMovie(movId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	related, err := handler.MovieUsecase.GetRelated(movId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := handler.MovieUsecase.GetComments(movId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(domain.MovieResponse{
		Movie:    movie,
		Related:  related,
		Comments: comments,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *MovieHandler) PostRating(w http.ResponseWriter, r *http.Request) {
	type ratingReq struct {
		MovieId string `json:"movieId"`
		UserId  string `json:"userId"`
		Rating  string `json:"rating"`
	}

	defer r.Body.Close()
	ratingreq := new(ratingReq)
	err := json.NewDecoder(r.Body).Decode(&ratingreq)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	movieId, err := strconv.ParseUint(ratingreq.MovieId, 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(ratingreq.UserId, 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	rating, err := strconv.Atoi(ratingreq.Rating)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	movieRating, err := handler.MovieUsecase.PostRating(movieId, userId, rating)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type ratingResp struct {
		NewMovieRating string `json:"newrating"`
	}

	out, err := json.Marshal(ratingResp{
		NewMovieRating: cast.FlToStr(movieRating),
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *MovieHandler) PostComment(w http.ResponseWriter, r *http.Request) {
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

	err = handler.MovieUsecase.PostComment(movieId, userId, commentreq.Content, commenttype)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comms, err := handler.MovieUsecase.GetComments(movieId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type commentsResp struct {
		Comments []domain.Comment `json:"reviews"`
	}

	out, err := json.Marshal(commentsResp{
		Comments: comms,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
