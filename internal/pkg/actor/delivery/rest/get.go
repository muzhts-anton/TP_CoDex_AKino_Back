package actdelivery

import (
	"codex/internal/pkg/domain"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func (handler *ActorHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	actId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	movies, err := handler.ActorUsecase.GetMovies(actId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	related, err := handler.ActorUsecase.GetRelated(actId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actor, err := handler.ActorUsecase.GetActor(actId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(domain.ActorResponse{
		Person:  actor,
		Related: related,
		Movies:  movies,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
