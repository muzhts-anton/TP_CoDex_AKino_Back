package anndelivery

import (
	"codex/internal/pkg/domain"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func (handler *AnnouncedHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movs, err := handler.AnnouncedUsecase.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(movs)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *AnnouncedHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	movie, err := handler.AnnouncedUsecase.GetMovie(movId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	related, err := handler.AnnouncedUsecase.GetRelated(movId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(domain.AnnouncedResponse{
		Announced: movie,
		Related:   related,
	})
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
