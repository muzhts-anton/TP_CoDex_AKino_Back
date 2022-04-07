package coldelivery

import (
	"codex/internal/pkg/domain"

	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (handler *CollectionsHandler) GetCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	colId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.ParseId.Error(), http.StatusBadRequest)
		return
	}

	coll, err := handler.CollectionsUsecase.GetCollection(colId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(coll)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *CollectionsHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	feed, err := handler.CollectionsUsecase.GetFeed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(feed)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
