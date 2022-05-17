package movdelivery

import (
	"codex/internal/pkg/domain"
	"codex/internal/pkg/sessions"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	var rexiewExist, userRating string
	var collectionsInfo []domain.CollectionInfo
	userId, err := sessions.CheckSession(r)
	if err == domain.Err.ErrObj.UserNotLoggedIn {
		rexiewExist = ""
		userRating = ""
		collectionsInfo = []domain.CollectionInfo{}
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		rexiewExist, userRating, err = handler.MovieUsecase.GetReviewRating(movId, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		collectionsInfo, err = handler.MovieUsecase.GetCollectionsInfo(userId, movId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	out, err := json.Marshal(domain.MovieResponse{
		Movie:           movie,
		Related:         related,
		Comments:        comments,
		ReviewExist:     rexiewExist,
		UserRating:      userRating,
		CollectionsInfo: collectionsInfo,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
