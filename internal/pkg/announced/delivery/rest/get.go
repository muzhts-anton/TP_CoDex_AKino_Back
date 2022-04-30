package anndelivery

import (
	"codex/internal/pkg/domain"

	"net/http"
	"encoding/json"

	// "github.com/gorilla/mux"
	// "strconv"
)

func (handler *AnnouncedHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movs, err := handler.AnnouncedUsecase.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(movs)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// func (handler *AnnouncedHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	movId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	movie, err := handler.MovieUsecase.GetMovie(movId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	related, err := handler.MovieUsecase.GetRelated(movId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	comments, err := handler.MovieUsecase.GetComments(movId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var rexiewExist, userRating string

// 	userId, err := sessions.CheckSession(r)
// 	if err == domain.Err.ErrObj.UserNotLoggedIn {
// 		rexiewExist = ""
// 		userRating = ""
// 	} else if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	} else {
// 		rexiewExist, userRating, err = handler.MovieUsecase.GetReviewRating(movId, userId)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 	}

// 	out, err := json.Marshal(domain.MovieResponse{
// 		Movie:       movie,
// 		Related:     related,
// 		Comments:    comments,
// 		ReviewExist: rexiewExist,
// 		UserRating:  userRating,
// 	})
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }
