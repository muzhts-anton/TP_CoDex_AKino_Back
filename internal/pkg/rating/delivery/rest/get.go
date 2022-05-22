package ratdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/rating/delivery/grpc"
	"codex/internal/pkg/utils/cast"

	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type RatingResp struct {
	NewMovieRating string `json:"newrating"`
}

type RatingReq struct {
	MovieId string `json:"movieId"`
	UserId  string `json:"userId"`
	Rating  string `json:"rating"`
}

func (handler *RatingHandler) PostRating(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ratingreq := new(RatingReq)
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

	movieRating, err := handler.RatingUsecase.PostRating(context.Background(), &grpc.Data{
		MovieId: movieId,
		UserId:  userId,
		Rating:  int32(rating),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(RatingResp{
		NewMovieRating: cast.FlToStr(float64(movieRating.GetRating())),
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
