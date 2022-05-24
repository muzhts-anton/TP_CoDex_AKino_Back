package ratdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/rating/delivery/grpc"
	"codex/internal/pkg/utils/cast"
	
	"io/ioutil"
	"context"
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
)

func (handler *RatingHandler) PostRating(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ratingreq := new(RatingReq)
	err = easyjson.Unmarshal(b, ratingreq)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
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

	out, err := easyjson.Marshal(RatingResp{
		NewMovieRating: cast.FlToStr(float64(movieRating.GetRating())),
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
